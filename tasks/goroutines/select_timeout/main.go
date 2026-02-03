package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	output := make(chan Response, 1)
	d1 := []byte(`{"greeting": "Hello, Jack Mills! You have 1 unread message.", "favoriteFruit": "strawberry"}`)
	d2 := []byte(`{"greeting": "Hello, Robert Turner! You have 2 unread messages.", "favoriteFruit": "apple"}`)
	d3 := []byte(`{"greeting": "Hello, Vicki Turner! You have 3 unread messages.", "favoriteFruit": "pineapple"}`)

	go connectOne(ctx, output, d1)
	go connectTwo(ctx, output, d2)
	go connectThree(ctx, output, d3)

	select {
	case res := <-output:
		fmt.Println("Printing output:", string(res.Resource))
	case <-ctx.Done():
		fmt.Println("Connection timed out")
	}
}

func connectOne(ctx context.Context, output chan Response, byteSt []byte) {

	fmt.Println("Attempting to connect 1....")

	select {
	case <-ctx.Done():
		return
	case <-time.After(5 * time.Second):

        // every blocking operation needs its own select with cancellation. 
        // Sending is a blocking operation.
		select { 
		case output <- Response{Resource: byteSt}:
		case <-ctx.Done():
			return
		}
	}
}

func connectTwo(ctx context.Context, output chan Response, byteSt []byte) {

	fmt.Println("Attempting to connect 2....")

	select {
	case <-ctx.Done():
		return
	case <-time.After(4 * time.Second):
		select {
		case output <- Response{Resource: byteSt}:
		case <-ctx.Done():
			return
		}
	}
}

func connectThree(ctx context.Context, output chan Response, byteSt []byte) {

	fmt.Println("Attempting to connect 3....")

	select {
	case <-ctx.Done():
		fmt.Println("Goroutine timed out")
		return
	case <-time.After(4 * time.Second):
		select {
		case output <- Response{Resource: byteSt}:
		case <-ctx.Done():
			return
		}
	}
}

type Response struct {
	Resource []byte
}

/*
Main goroutine starts. It spawns 3 independant goroutines, and blocks at the main select while
these 3 goroutines are doing their work. Each has their own instructions with the select statement.
go connectOne(ctx, output, d1) prints and moves to its select statement. It blocks. The select
statement instructs if you get a call from the parent goroutine (main) <-ctx.Done(), then exist.
If not, keep on for <-time.After(5 * time.Second) and do your job.

In this case, this goroutine doesn't get to send to output, because the context cancels itself
when the deadline expires after 3 secs. ctx.Done() channel closes. After 3 secs, 
go connectOne(ctx, output, d1) receives a done signal ctx.Done() and exits/stops. Based on the 
code all the goroutines have the same fate.

Going back to the main goroutine, it is blocked at its select. 3 secs pass, cancellation fires,
ctx.Done() is sent. The select statement has a done signal. main() is released and defer cancel()
is invoked to remove ctx and clean up resources.
*/
