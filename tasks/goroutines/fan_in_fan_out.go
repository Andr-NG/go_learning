package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

func StreamLogs(logs []string, logsc chan string) {
	for _, log := range logs {
		logsc <- log
	}
	close(logsc)
}

func ParseLogs(logsc chan string, parsedLogc chan ParsedLog) {
	for log := range logsc {
		s := strings.Split(log, "|")

		t, err := FormatDate(s[0])
		if err != nil {
			fmt.Printf("error when formating date str-to-int: %v\n", err)
			return
		}

		num, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Printf("error when formating user id str-to-int: %v\n", err)
		}

		sts, err := strconv.Atoi(s[3])
		if err != nil {
			fmt.Printf("error when formating status str-to-int: %v\n", err)
		}

		ltn, err := strconv.Atoi(s[4])
		if err != nil {
			fmt.Printf("error when formating latency str-to-int: %v\n", err)
		}

		parsed := ParseLog(t, num, sts, ltn)

		parsedLogc <- parsed
	}
	close(parsedLogc)
}

func EnrichLogs(log ParsedLog, wg *sync.WaitGroup, output chan EnrichedLog) {
	defer wg.Done()
	var userTier, severity string

	if log.UserID%2 == 0 {
		userTier = "premium"
	} else {
		userTier = "free"
	}

	if log.Status >= 500 {
		severity = "error"
	} else if log.Status >= 400 {
		severity = "warning"

	} else {
		severity = "info"
	}

	enriched := EnrichParsedLog(log, severity, userTier)
	output <- enriched
}

func CollectEnrichedLogs(output chan EnrichedLog) []EnrichedLog {
	result := make([]EnrichedLog, 0, len(output))
	for enrichedLog := range output {
		result = append(result, enrichedLog)
	}
	return result
}

const layout = "2006-01-02T15:04:05Z"

func FormatDate(s string) (time.Time, error) {
	t, err := time.Parse(layout, s)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func ParseLog(t time.Time, num, sts, ltn int) ParsedLog {
	return ParsedLog{
		Timestamp: t,
		UserID:    num,
		Status:    sts,
		LatencyMs: ltn,
	}
}

func EnrichParsedLog(l ParsedLog, s, ut string) EnrichedLog {
	return EnrichedLog{
		ParsedLog: l,
		UserTier:  ut,
		Severity:  s,
	}
}

type ParsedLog struct {
	Timestamp time.Time
	UserID    int
	Action    string
	Status    int
	LatencyMs int
}

type EnrichedLog struct {
	ParsedLog
	UserTier string // "free" or "premium"
	Severity string // "info" | "warn" | "error"
}
