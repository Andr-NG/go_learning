package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func generateData(ctx context.Context, out chan int) {
	defer close(out)
	for {
		select {
		case <-ctx.Done():
			return
		case out <- rand.IntN(50):
			fmt.Println("sending data to out...")
		}
	}
}
func consumeData(ctx context.Context, out chan int) {
	for {
		select {
        // select blocks the execution until one of these happens:
        // case <-ctx.Done() => return OR case num, ok := <-out => processing data
		case <-ctx.Done():
			return

		case num, ok := <-out:
        // processing data one by one inside select. No range
			if !ok {
                fmt.Println("channel closed")
                return
            } else {
                time.Sleep(time.Millisecond * 800)
                fmt.Printf("receving data %d...\n", num)
            }
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	out := make(chan int, 50)

	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()

	go generateData(ctx, out)

	consumeData(ctx, out)

}
