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
		fmt.Println(counter)
		done <- struct{}{}
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
