package main

import "sync"

func main() {

	c := AtomicCounter{
		n: 0,
	}

	wg := &sync.WaitGroup{}

	PerformCounter(&c, wg, 16)

	wg.Wait()
	
	c.GetValue()
}
