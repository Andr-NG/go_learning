package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	jobs := []Job{{ID: 1}, {ID: 15}, {ID: 19}, {ID: 20}, {ID: 12}, {ID: 212}, {ID: 453}, {ID: 43}, {ID: 64}, {ID: 30}}
	jobsc := make(chan Job)
	res := make(chan Result)

	workerCount := 3
	wg.Add(workerCount)

	for i := 1; i <= workerCount; i++ {
		go worker(i, &wg, jobsc, res)
	}

	go produceJobs(jobs, jobsc)

	// coordinator
	go func(){
		wg.Wait()
		close(res)
	}()
	
	fin := collectResults(res)

	fmt.Println(fin)

}

