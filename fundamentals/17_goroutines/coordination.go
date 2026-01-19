package main

/*
Goal: Learn that sleeping is wrong
New tool: sync.WaitGroup

Rewrite many_goroutines:

Remove all Sleep
Use WaitGroup
Ensure program exits only after all goroutines finish

Key takeways:

wg sync.WaitGroup manages goroutines
wg.Add() and wg.Wait() apply outside the goroutine.

*/

import (
	"fmt"
	"sync"
)



func InvokeRoutineWaitGroup(){
	fmt.Println("Before calling goroutine")
	
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1) // Adding a goroutine as we go
		go func(){
			defer wg.Done()
			fmt.Printf("Starting goroutine %d\n", i)
		}()
	}
	// Waiting for goroutines to finish
	wg.Wait()

	fmt.Println("After calling goroutine")
}

