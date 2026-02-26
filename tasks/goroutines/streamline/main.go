package main

import (
	"fmt"
	"sync"
)

/*<timestamp>|<user_id>|<action>|<status>|<latency_ms>

Rules:

UserTier
even user ID → "premium"
odd user ID → "free"

Severity
status >= 500 → "error"
status >= 400 → "warn"
else → "info"
*/

func main() {
	fmt.Println("Starting main....")
	wg := sync.WaitGroup{}

	logs := []string{
		"2026-01-27T10:15:03Z|42|login|200|120",
		"2026-01-27T10:15:04Z|17|purchase|500|430",
		"2026-01-27T10:15:05Z|42|view|200|30",
		"2026-01-27T10:15:06Z|99|login|401|80",
		"2026-01-27T10:15:07Z|17|logout|200|20",
		"2026-01-27T10:15:08Z|42|purchase|200|210",
	}
	logsCh := make(chan string)
	parsedLogCh := make(chan ParsedLog)
	output := make(chan EnrichedLog)

	// stream logs
	// close(logsCh), so range logsCh can further exit from the loop and doesn't block
	go StreamLogs(logs, logsCh)

	// parse streamed logs
	// close(parsedLogCh), so range parsedLogc can further exit from the loop and doesn't block
	go ParseLogs(logsCh, parsedLogCh)

	// enrich logs
	for log := range parsedLogCh {
		wg.Add(1)
		go EnrichLogs(log, &wg, output)

	}
	// coordinator
	go func() {
		// wg.Wait() blocks the invoking goroutine waiting for wg.Done() in go EnrighLogs
		// to bring the counter to, so the function can proceed to close(output)
		wg.Wait()
		close(output)
	}()

	result := CollectEnrichedLogs(output)
	fmt.Println(result)
}
