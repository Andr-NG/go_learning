package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)



func main(){
    fmt.Println(Status(2))
    sl := []string{"sadda", "dsadsa", "dsadsd"}
    fmt.Println(sl)
}


type Store struct {
	mu   sync.Mutex
	jobs map[int]Job
}

func NewStore() *Store {
	return &Store{
		jobs: make(map[int]Job),
	}
}

func (st *Store) CreateJob(id int, name string) error {
	st.mu.Lock()
	defer st.mu.Unlock()

	_, ok := st.jobs[id]
	if ok {
		return ErrDuplicateJob
	}

	for _, j := range st.jobs {
		if j.Name == name {
			return ErrSameJobName
		}
	}

	nj := NewJob(id, name)
	st.jobs[id] = *nj

	return nil
}

func (st *Store) GetJob(id int) (Job, error) {
	st.mu.Lock()
	defer st.mu.Unlock()

	job, ok := st.jobs[id]
	if ok {
		return job, nil
	}

	return Job{}, ErrJobNotFound
}

func (st *Store) UpdateJob(id int, updater func(Job) (Job, error)) error {
	st.mu.Lock()
	defer st.mu.Unlock()

	cj, ok := st.jobs[id]
	if !ok {
		return ErrJobNotFound
	}

	nj, err := updater(cj)
	if err != nil {
		return err
	}
	if nj.Name == cj.Name {
		return ErrSameJobName
	}

	st.jobs[id] = nj

	return nil
}

func (st *Store) ListAllJobs() []Job {
	st.mu.Lock()
	defer st.mu.Unlock()
	res := make([]Job, 0, len(st.jobs))
	total := len(st.jobs)
	if total == 0 {
		return res
	}

	for _, job := range st.jobs {
		res = append(res, job)
	}
	return res

}

type Job struct {
	ID           int
	Name         string
	Status       Status
	ErrorMessage string
	UpdatedAt    time.Time
}

func NewJob(id int, name string) *Job {
	return &Job{ID: id, Name: name, Status: 0}
}

type Status int

const (
	Pending Status = iota
	Running
	Failed
	Completed
)


func (s Status) String() string {
	return []string{
		"pending",
		"running",
		"failed",
		"completed",
	}[s]
}


var (
	ErrDuplicateJob = errors.New("No duplicate jobs allowed")
	ErrSameJobName  = errors.New("No same job names allowed")
	ErrJobNotFound  = errors.New("Job ID not found")
)
