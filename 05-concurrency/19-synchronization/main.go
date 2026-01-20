package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func mutexExample() {
	var mu sync.Mutex
	counter := 0

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Mutex counter:", counter)
}

func rwMutexExample() {
	var mu sync.RWMutex
	data := 0

	mu.Lock()
	data = 42
	mu.Unlock()

	mu.RLock()
	fmt.Println("RWMutex read:", data)
	mu.RUnlock()
}

func onceExample() {
	var once sync.Once

	init := func() {
		fmt.Println("Initialized once")
	}

	for i := 0; i < 3; i++ {
		go once.Do(init)
	}

	time.Sleep(200 * time.Millisecond)
}

func condExample() {
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)
	ready := false

	go func() {
		mu.Lock()
		for !ready {
			cond.Wait()
		}
		fmt.Println("Worker: proceeding")
		mu.Unlock()
	}()

	time.Sleep(300 * time.Millisecond)

	mu.Lock()
	ready = true
	cond.Signal()
	mu.Unlock()
}

func atomicExample() {
	var counter int64

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println("Atomic counter:", counter)
}

func main() {

	fmt.Println("Mutex example:")
	mutexExample()

	fmt.Println("\nRWMutex example:")
	rwMutexExample()

	fmt.Println("\nWaitGroup & Once example:")
	onceExample()

	fmt.Println("\nCond example:")
	condExample()

	fmt.Println("\nAtomic example:")
	atomicExample()
}
