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
	"time"
)

func RoutineIdWaitGroup(wg *sync.WaitGroup, id int){
	defer wg.Done()
	fmt.Printf("Starting goroutine %d\n", id)
	time.Sleep(2 *time.Second)
}

func InvokeRoutineWaitGroup(){
	fmt.Println("Before calling goroutine")
	
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Adding a goroutine as we go
		go RoutineIdWaitGroup(&wg, i)
	}
	// Waiting for goroutines to finish
	wg.Wait()

	fmt.Println("After calling goroutine")
}

