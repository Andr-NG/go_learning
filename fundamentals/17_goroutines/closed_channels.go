package main

import "fmt"

func ReadFromClosedChannel() {
	ch := make(chan int)

	go func() {
		ch <- 5
	}()

	val, ok := <-ch
	fmt.Println("Value", val)
	fmt.Println("Exists", ok)

	close(ch)
	val1, ok1 := <-ch
	fmt.Println("Value", val1)
	fmt.Println("Exists", ok1)
}

func PanicSendingOnClosedChannel() {

	/*
	This function will cause panic, because the sender function closes the channel prematurely.
	Sender sends ch <-
	Recevier gets val, ok := <-ch
	Execution continues and sender hits close(ch)
	*/


	ch := make(chan int)

	go func() {
		for i := range 5 {
			ch <- i
			/*
				close() runs after the for loop
				The loop can only advance when each send succeeds
				Each send succeeds only when a receiver receives
				So close() cannot execute until all sends are finished.
			*/
			close(ch)
		}
		// close(ch) THE FIX! Only close once the loop is totally done

	}()
	val, ok := <-ch
	fmt.Println(val, ok)

	// The receiver stays here, meeting the sender 5 times
    // for val := range ch {
    //     fmt.Println("Received:", val)
    // }   
}
