package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 1000; i++ {
		fmt.Println("number:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	fmt.Println("Main started")

	// create goroutine
	go printNumbers()

	// anonymous goroutine
	go func() {
		fmt.Println("Hello from goroutine")
	}()

	// main thread sleep so goroutines can run
	time.Sleep(1 * time.Second)

	fmt.Println("Main finished")
}
