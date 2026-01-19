package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
    n int64
}

func (c *AtomicCounter) Inc() {
    atomic.AddInt64(&c.n, 1)
}

func (c *AtomicCounter) GetValue() {
    fmt.Println(atomic.LoadInt64(&c.n))
}

// *AtomCounter because we mutate a shared resource, which is instantiated in a different file
func PerformCounter(c *AtomicCounter, wg *sync.WaitGroup, r int){
    
    for i := 1; i < r; i++ {
        wg.Add(1)
        go func(){
            defer wg.Done()
            c.Inc()        
        }()
    }
}