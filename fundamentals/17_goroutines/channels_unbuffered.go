/*
Goal: Use channels only as a handoff mechanism
Tools allowed: make(chan T), send, receive

Write a program where:

Goroutine computes a number
Sends it through a channel
main receives and prints it

Rules:
One send
One receive
No loops

Key takeaways:
There must be at least 2 goroutines (sender and receiver) for unbuffered channel communication 
Channel send = synchronization point
Channels = synchronization + value handoff
Channels != resource manager
*/


package main

import (
	"time"
	"fmt"
)


func Worker(i int, channeledValue chan int){
	fmt.Println("Starting worker...")
	time.Sleep(1 * time.Second)

	// Send occurs. Worker blocks here until receiver is ready
	channeledValue <- i
	fmt.Println("Completing send...")
}


func InvokeChannelling(){

	// Creating unbuffered channel
	chanValue := make(chan int)

	// Worker runs concurrently 
	go Worker(5, chanValue)

	fmt.Println("Inside main: waiting for worker")

	// Caller blocks until worker sends. Value is received and printed
	fmt.Println(<-chanValue)

	fmt.Println("Inside main: received value, existing")

}