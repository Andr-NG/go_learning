package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println(">>> CANCELING CONTEXT")
		cancel()
	}()
	generateJobs := func(num int) []Job {
		s := make([]Job, 0, num)
		for i := 1; i <= num; i++ {
			s = append(s, Job{ID: i, Payload: fmt.Sprintf("Payload %d", i)})
		}
		return s
	}
	jobsList := generateJobs(rand.IntN(100) + 1)

	res, err := ProcessJobs(ctx, jobsList, 5)
	fmt.Println(res)
	if err != nil {
		fmt.Print(err)
	}

}

func ProcessJobs(ctx context.Context, jobs []Job, workers int) ([]Result, error) {

	jobsCh := make(chan Job)
	resCh := make(chan Result)
	resList := make([]Result, 0, len(jobs))
	var wg sync.WaitGroup

	wg.Add(workers)

	go func() {
		defer close(jobsCh)
		for _, job := range jobs {
			select {
			case <-ctx.Done():
				fmt.Println("Producer exiting due to context cancellation")
				return
			case jobsCh <- job:
				fmt.Println("Adding job to chan")
			}
		}
	}()

	for range workers {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Worker exiting due to context cancellation")
					return
				case j, ok := <-jobsCh:
					if !ok {
						return
					}
					fmt.Println("Appending job to final res list")
					time.Sleep(690 * time.Millisecond)
					res := Result{
						JobId: j.ID, Value: j.Payload,
					}
					select {
					case resCh <- res:
					case <-ctx.Done():
						return
					}
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resCh)
	}()

	for {
		select {
		case res, ok := <-resCh:
			if !ok {
				return resList, nil
			}
			fmt.Println("Appending res to resList")
			resList = append(resList, res)
		case <-ctx.Done():
			return resList, ctx.Err()
		}
	}

}

type Job struct {
	ID      int
	Payload string
}

type Result struct {
	JobId int
	Value string
}
