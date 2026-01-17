package main

import "fmt"

/*
Spawn 5 goroutines, each printing its index.

Takeways:

Goroutines run concurrently, not sequentially.
If order matters, you must synchronize (channels, WaitGroup, etc.).

*/


func routine_id(id int){
	fmt.Printf("Starting goroutine %d\n", id)
}
