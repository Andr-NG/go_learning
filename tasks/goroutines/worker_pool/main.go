package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	jobCh := make(chan Job)
	resCh := make(chan Result)
	generateJobs := func(num int) []Job {
		s := make([]Job, 0, num)
		for i := 1; i <= num; i++ {
			s = append(s, Job{ID: i})
		}
		return s
	}
	jobsArray := generateJobs(rand.IntN(50))
	workerPool := rand.IntN(9) + 1

	go produceJobs(jobsArray, jobCh)

	// kickstarts workPool number of goroutines and exists.
	// goroutines live independently.
    // the loop starts concurrency.
    // channels and WaitGroups control it.
	wg.Add(workerPool)
	for i := 1; i <= workerPool; i++ {
		go worker(i, &wg, jobCh, resCh)
	}

	// coordinator
	go func() {
		wg.Wait()
		close(resCh)
	}()

	res := collectResults(resCh)
	fmt.Println(res)
}

type Result struct {
	JobID    int
	WorkerID int
}

type Job struct {
	ID int
}

func worker(id int, wg *sync.WaitGroup, jobs <-chan Job, res chan<- Result) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Processing job %d with Job ID %d\n", id+1, job.ID)
		res <- Result{WorkerID: id, JobID: job.ID}
	}
}

func produceJobs(s []Job, jobsc chan Job) {
	for id, job := range s {
		fmt.Printf("Printing job %d with Job ID %d\n", id, job.ID)
		jobsc <- job
	}
	close(jobsc)
}

func collectResults(res chan Result) []Result {
	resArray := make([]Result, 0)

	for result := range res {
		resArray = append(resArray, result)
	}

	return resArray

}
