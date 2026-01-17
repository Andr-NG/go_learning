package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Exercise 7: Simple pipeline

Build:
Producer goroutine → channel → consumer goroutine
Producer sends numbers 1–5
Consumer prints them

Rules:
Producer closes the channel
Consumer uses range

Key questions:
Who closes the channel? Sender. ALWAYS.
What happens if you don’t? Deadlock, because range will hang forever.
*/

func Logger(wg *sync.WaitGroup,chnl chan string) {
	defer wg.Done()

	fmt.Println("Inside Pipeline: start iterating over channel")
	for msg := range chnl {
		fmt.Printf("Message from pipeline: %s\n", msg)
	}

	// Will NOT be printed unless wg.Waitgroup is present 
	fmt.Println("Inside Pipeline: channel closed, exiting")
}

func InvokeLogger(slc []string) {

	// Initialising wg to control goroutines
	var wg sync.WaitGroup

	// Initialising channel to hand off string values
	ch := make(chan string)

	// Adding goroutine to group
	wg.Add(1)
	go Logger(&wg, ch)

	for _, val := range slc {
		time.Sleep(500 * time.Millisecond)
		ch <- val
	}

	// Closing a channel
	close(ch)

	// Waiting for goroutines to finish
	wg.Wait()
}