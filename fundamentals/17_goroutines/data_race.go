/*
Write a program that:
Increments a shared counter from 10 goroutines
Run with -race

Then:
Fix it using a mutex
Fix it again using a channel

Compare:
Code size
Readability
Mental load
*/

// package main

// import "fmt"
package main

import (
	"fmt"
	"sync"
	"time"
)

func GoRoutineRace(r int) int {
	var counter int
	for range r {
		go func() {
			counter++
		}()
	}
	return counter
}

// Protect shared state
// Only one goroutine can access counter at a time
func GoRoutineFixedMutex() {
	// Declaring count and Mutex
	var (
		counter int
		mu      sync.Mutex
	)

	// Locking and unlocking shared resource/state to prevent race condition
	for range 100 {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	// Will show 100 very unlikely because nothing control goroutines
	fmt.Println(counter)
}

func GoRoutineFixedChannel() {
	// Declaring channels
	incr := make(chan int)
	done := make(chan struct{})

	go func() {
		counter := 0
		for v := range incr {
			counter += v
		}
		done <- struct{}{}
		fmt.Println(counter)
	}()

	for range 1000 {
		go func() {
			incr <- 1
		}()
	}


	time.Sleep(time.Second)
	close(incr)
	// <-done

}

func GoroutineStartSignalled(){
	signal := make(chan any)
	wg := &sync.WaitGroup{}

	for i := range 5 {

		// Adding goroutine to WaitGroup
		wg.Add(1)

		go func(){
			// Making sure counter is decremented
			defer wg.Done()

			// Ensuring simultaneous start for all 5 goroutines. Nothing is execute below
			// until the channel is closed
			<-signal
			fmt.Println("index", i)
		}()

		// wg.Go(func() {
		// 	<-signal
		// 	fmt.Println("index", i)
		// })
	}
	
	// Closing channel
	close(signal)
	wg.Wait()

}