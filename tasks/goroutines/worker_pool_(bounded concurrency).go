package main

import (
	"fmt"
	"sync"
)

/*
Scenario
You’re building a service that processes incoming jobs (e.g., image resizing).
You must limit CPU usage.

Problem
Input: slice of Job{ID int}
Process jobs concurrently using exactly N workers
Each job returns Result{JobID int, WorkerID int}
Preserve all results (order does NOT matter)

Requirements

Use goroutines + channels
Use sync.WaitGroup
Workers must exit cleanly when work is done

Expected Result
Every job processed once
No goroutine leaks
Program terminates deterministically
*/

// Multiple goroutines write into one channel.
// A separate goroutine to wait and close is required.
func MultipleGoroutineProcessJobs(s []Job, wg *sync.WaitGroup) []Result {

	jobsc := make(chan Result)
	res := make([]Result, 0, len(s))

	for ind, job := range s {
		wg.Add(1)
		// capturing creates new variables scoped to the loop iteration, so each goroutine
		// now captures its own copy of ind and job.
		go func(ind int, job Job) {
			defer wg.Done()
			jobsc <- Result{JobID: job.ID, WorkerID: ind}
			fmt.Println("Processing job", job.ID)
		}(ind, job)
	}

	// A separate goroutine to wait and close is required.
	go func() {
		wg.Wait()
		close(jobsc)
	}()

	// range jobsc only exits when the channel is closed. Otherwise, deadlock, becasue range jobsc
	// will keep reading
	for val := range jobsc {
		res = append(res, val)
	}

	return res
}

// One goroutine writes into one channel.
// Channel is closed right after values are sent. No more goroutines required.
func OneGoroutineProceessJobs(s []Job) []Result {
	jobs := []Job{{ID: 1000}, {ID: 15212}, {ID: 19122}, {ID: 2320}, {ID: 13132}}

	jobsc := make(chan Result)
	res := make([]Result, 0, len(jobs))

	go func() {
		for ind, job := range jobs {
			jobsc <- Result{JobID: job.ID, WorkerID: ind}
			fmt.Println("Processing job", job.ID)
		}
		// Sender knows that no values to be sent. Closing channel.
		close(jobsc)
	}()

	// range jobsc only exits when the channel is closed. Otherwise, deadlock, becasue range jobsc
	// will keep reading
	for val := range jobsc {
		res = append(res, val)
	}
	return res
}

type Result struct {
	JobID    int
	WorkerID int
}

type Job struct {
	ID int
}

func WorkerPool(s []Job, wg *sync.WaitGroup, workerPool int) {

	// go func() {
	// 	close(res)
	// 	wg.Wait()
	// }()

	// fin := collectResults(res)
	// fmt.Println(fin)
}

func worker(id int, wg *sync.WaitGroup, jobs chan Job, res chan Result) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Processing job %d with Job ID %d\n", id+1, job.ID)
		res <- Result{WorkerID: id, JobID: job.ID}
	}
}

func produceJobs(s []Job, jobsc chan Job) {
	for id, job := range s {
		fmt.Printf("Printing job %d with Job ID %d\n", id+1, job.ID)
		jobsc <- job
	}
	close(jobsc)
}

func coordinator(res chan Result, wg *sync.WaitGroup) {
	wg.Wait()
	close(res)
}

func collectResults(res chan Result) []Result {
	resArray := make([]Result, 0, len(res))

	for result := range res {
		resArray = append(resArray, result)
	}

	return resArray

}
