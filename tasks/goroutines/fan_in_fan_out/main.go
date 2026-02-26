package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	done := make(chan struct{})
	fmt.Println("Start generating data...")
	stream := generateData(done, 100000)

	fmt.Println("Start verifying evenness...")
    // fanning N goroutines to improve performance
	chns := fanOut(done, stream, 10)

    // merging N streams
	output := fanIn(done, chns)

	fmt.Println("Start collecting even numbers...")
	res := collector(done, output, 15)

	for num := range res {
		fmt.Println(num)
	}
	fmt.Println(time.Since(start))
}

// Generating data for further use by putting random ints to a channel
func generateData(done chan struct{}, num int) <-chan int {
	outCh := make(chan int)

	go func() {
		defer close(outCh)
		for {
			select {
			case <-done:
				fmt.Println("Stopping generation")
				return
			case outCh <- rand.IntN(num):
			}
		}
	}()

	return outCh

}

// Spinning a certain number of goroutine to process data sent across the channel
// and adding each channel to a slice
func fanOut(done chan struct{}, stream <-chan int, workers int) []<-chan int {
	chSlice := make([]<-chan int, workers)

	for i := range workers {
		chSlice[i] = findEven(done, stream)
	}

	return chSlice
}

// Collecting data from the N launched goroutines concurrently into one channel,
// draining all channels simultaneously, because each input needs an independent forwarder
func fanIn(done chan struct{}, chSlice []<-chan int) <-chan int {
	output := make(chan int)

    // wg is used to ensure that we close output only after all the input channels are drained
	var wg sync.WaitGroup
    
    // helper to extract data from each input channel
	transfer := func(ch <-chan int) {
		defer wg.Done()
		for val := range ch {
			select {
			case <-done:
				return
			case output <- val:
			}
		}
	}

    // ranging over each channel from the slice of channel, apllying the helper
	for _, ch := range chSlice {
		wg.Add(1)
		go transfer(ch)
	}

	go func() {
        // coordinator to wait for all the goroutine finish before closing
		wg.Wait()
		close(output)
	}()
	return output
}

func findEven(done chan struct{}, dataCh <-chan int) <-chan int {
	outCh := make(chan int)
	isEven := func(randInt, randTime int) bool {
		// adding more unnecessary logic to slow down work
		time.Sleep(time.Duration(randTime) * time.Millisecond)
		if randInt%2 == 0 {
			return true
		} else {
			return false
		}
	}

	go func() {
		defer close(outCh)
		for {
			timeFrame := []int{500, 1200, 700, 900}
			randTime := timeFrame[rand.IntN(len(timeFrame)-1)]
			select {
			case <-done:
				return
			case num := <-dataCh:
				if isEven(num, randTime) {
					select {
					case <-done:
						fmt.Println("Abort checking evenness")
					case outCh <- num:
					}
				}
			}
		}
	}()
	return outCh
}

// Collecting N number of streamed values. It stops once it hits the cap
func collector(done chan struct{}, stream <-chan int, cap int) <-chan int {
	collected := make(chan int)

	go func() {
		defer close(collected)
		for range cap {
			select {
			case <-done:
				fmt.Println("Aborting collecting data...")
				return
			case num := <-stream:
				select {
				case <-done:
					fmt.Println("Aborting collecting data...")
					return
				case collected <- num:
				}
			}
		}
	}()
	return collected

}
