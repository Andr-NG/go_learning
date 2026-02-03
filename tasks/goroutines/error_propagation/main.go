package main

import (
	"context"
	"fmt"
	"math/rand/v2"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	errCh := make(chan error, 1)

	someFunc := func(ctx context.Context, errCh chan error) {

		if ctx.Err() != nil {
			return
		}

		if rand.IntN(100) < 40 {
			fmt.Println("This goroutine has failed")

			// Report error ONLY if not cancelled
			select {
			case errCh <- fmt.Errorf("goroutine failed"):
			case <-ctx.Done():
			}
		} else {
			fmt.Println("This goroutine is working")
		}

	}

	for range 20 {
		go someFunc(ctx, errCh)
	}

	<-errCh
	cancel()
}
