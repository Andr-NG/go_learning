package main

/*
Scenario:
Metrics collector running in background.

Problem:
A goroutine collecting  ticks every second
On each tick, it “collects metrics”
One more goroutine flushes metrics every 5 seconds.
Every 3rd flush fails. Memory preserved.
Main goroutine shuts it down after N seconds.

Requirements:
Implement channel ownership model
Use time.Ticker
Use context propagation
Ensure tickers are stopped
Worker must exit immediately on shutdown

Expected Result:
Clean shutdown
No leaked ticker
No blocked goroutines
*/



import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	var genwg sync.WaitGroup
	cmdCh := make(chan command)

	fmt.Println("Running storage. Collecting metrics...")
	wg.Add(1)
	go runStorage(&wg, ctx, cmdCh)

	genwg.Go(func() {
		generateCmd(ctx, cmdCh)
	})

	time.Sleep(time.Second * 30)
	cancel()
	// making sure generateCmd finishes first
	genwg.Wait()
	// making sure runStorage finishes after generateCmd
	wg.Wait()
	fmt.Println("Finishing running storage...")

}


/*
Stores and mutates shared slice through commands sent via cmdCh by other goroutines.
Eliminates mutex to provide clear shared state ownership
*/
func runStorage(wg *sync.WaitGroup, ctx context.Context, cmdCh chan command) {
	defer wg.Done()
	items := make([]int, 0)
	flushCounter := 1

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("final flush of remaining %d metrics\n", len(items))
			items = items[:0]
			fmt.Printf("%d metrics after final flush\n", len(items))
			return

		case cmd := <-cmdCh:

			// c := cmd.(type) means assigning a possible type to c
			// so we can reach embedded fields
			switch c := cmd.(type) {
			case storeCmd:
				items = append(items, c.value...)

			case flushCmd:
				fmt.Printf("%d metrics to flush...\n", len(items))
				if flushCounter%3 == 0 {
					fmt.Println("flush failed (buffer preserved)...", len(items))
				} else {
					fmt.Printf("flush success (%d metrics)...\n", len(items))
					// reset to len(0)
					items = items[:0]
				}
				flushCounter++
			}

		}

	}
}

/*
Generates commands sent to cmdCh. These commands are received and used by runStorage
*/
func generateCmd(ctx context.Context, cmdCh chan command) {
	ticker := time.NewTicker(1 * time.Second)
	flushTicker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	defer flushTicker.Stop()

	for {
		randInt := rand.IntN(20)+1
		select {
		case <-ctx.Done():
			fmt.Println("done sending commands")
			return

		case <-flushTicker.C:
			fmt.Println("sending flush cmd")
			select {
			case cmdCh <- flushCmd{}:
			case <-ctx.Done():
				return
			}

		case <-ticker.C:
			fmt.Println("sending store cmd")
			select {
			case cmdCh <- storeCmd{value: make([]int, randInt)}:
			case <-ctx.Done():
				return
			}

		}
	}

}

type command any

type storeCmd struct {
	value []int
}

type flushCmd struct {
}
