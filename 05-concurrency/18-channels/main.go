package main

import (
	"fmt"
	"time"
)

func unbufferedExample() {
	ch := make(chan int)

	go func() {
		fmt.Println("Sending 10")
		ch <- 10
		fmt.Println("Sent 10")
	}()

	time.Sleep(500 * time.Millisecond)
	v := <-ch
	fmt.Println("Received:", v)
}

func bufferedExample() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println("Buffered receive:", <-ch)
	fmt.Println("Buffered receive:", <-ch)
}

func closeAndRangeExample() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println("range received:", v)
	}
}

func selectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(600 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case v := <-ch1:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		}
	}
}

func main() {

	fmt.Println("Unbuffered channel:")
	unbufferedExample()

	fmt.Println("\nBuffered channel:")
	bufferedExample()

	fmt.Println("\nClose and range:")
	closeAndRangeExample()

	fmt.Println("\nSelect statement:")
	selectExample()
}
