package main

import (
	"context"
	"fmt"
	"time"
)

// ---------------- Worker Pool ----------------

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %d processing job %d\n", id, j)
		time.Sleep(200 * time.Millisecond)
		results <- j * 2
	}
}

// ---------------- Pipeline ----------------

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// ---------------- Context cancellation ----------------

func longTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("task cancelled")
			return
		default:
			fmt.Println("working...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

// ---------------- Rate limiting ----------------

func rateLimitedRequests() {
	limiter := time.Tick(500 * time.Millisecond)

	for i := 1; i <= 5; i++ {
		<-limiter
		fmt.Println("request", i)
	}
}

// ---------------- Main ----------------

func main() {

	// Worker pool
	fmt.Println("== Worker Pool ==")
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= 5; i++ {
		fmt.Println("result:", <-results)
	}

	// Pipeline
	fmt.Println("\n== Pipeline ==")
	c1 := gen(1, 2, 3)
	c2 := square(c1)

	for v := range c2 {
		fmt.Println(v)
	}

	// Context cancellation
	fmt.Println("\n== Context cancellation ==")
	ctx, cancel := context.WithCancel(context.Background())
	go longTask(ctx)

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)

	// Rate limiting
	fmt.Println("\n== Rate limiting ==")
	rateLimitedRequests()
}
