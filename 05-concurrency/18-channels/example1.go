package main

import (
	"fmt"
)

type Job struct {
	ID   int
	Data string
}

type Result struct {
	JobID  int
	Output string
}

func worker(jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		// simulate processing
		output := "processed: " + job.Data

		results <- Result{
			JobID:  job.ID,
			Output: output,
		}
	}
}

func main_run() {
	jobs := make(chan Job, 5) // buffered
	results := make(chan Result, 5)

	// start workers
	for i := 0; i < 3; i++ {
		go worker(jobs, results)
	}

	// send jobs
	for i := 1; i <= 10; i++ {
		jobs <- Job{ID: i, Data: fmt.Sprintf("task-%d", i)}
	}
	close(jobs)

	// collect results
	for i := 1; i <= 10; i++ {
		res := <-results
		fmt.Println(res)
	}
}
