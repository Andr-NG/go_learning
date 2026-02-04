package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	done := make(chan struct{})
	storage := newStorage()
	osCh := make(chan os.Signal, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(1 * time.Second)
		flushTicker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		defer flushTicker.Stop()

		for {
			select {
			case <-flushTicker.C:
				fmt.Println("flushing with current counter...", storage.flushCounter)
				storage.flush()
			case <-ticker.C:
				fmt.Println("collecting metrics...")
				storage.store(rand.IntN(50))
			case <-done:
				fmt.Println("FINAL flushing with current counter...", storage.flushCounter)
				storage.flush()
				return
			}
		}

	}()

	fmt.Println("main working until cancelled")
	signal.Notify(osCh, syscall.SIGINT, syscall.SIGQUIT)
	<-osCh
	close(done)
	wg.Wait()

}

type Storage struct {
	items        []int
	flushCounter int
}

func newStorage() *Storage {
	return &Storage{
		items:        make([]int, 0),
		flushCounter: 1,
	}
}

func (strg *Storage) store(item int) {

	strg.items = append(strg.items, item)
}

func (strg *Storage) flush() {

	if strg.flushCounter%3 == 0 {
		fmt.Println("flush failed (buffer preserved)")
		strg.flushCounter++

	} else {
		fmt.Printf("flush success (%d metrics)\n", len(strg.items))
		strg.items = make([]int, 0)
		strg.flushCounter++

	}

}
