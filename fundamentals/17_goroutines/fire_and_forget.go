package main
/*
Task:

Start a goroutine that prints "worker started"
Print "main finished" in main
Observe output order
Then make sure the worker always prints


Concepts:

go keyword
main lifecycle
Why goroutines die when main exits
*/

import (
	"fmt"
)

func StartWorker(){
	fmt.Println("Worker started")
}

