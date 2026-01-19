package main

import (
	"fmt"
	"os"
	"time"
)

func demoStack() {
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
}

func demoArguments() {
	x := 10
	defer fmt.Println("deferred x =", x)
	x = 20
}

func demoFile() {
	file, err := os.Create("temp.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("hello")
}

func demoTimer() {
	start := time.Now()
	defer func() {
		fmt.Println("elapsed:", time.Since(start))
	}()
	time.Sleep(200 * time.Millisecond)
}

func demoPanicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	panic("boom")
}

func main() {

	fmt.Println("Stack order:")
	demoStack()

	fmt.Println("\nArgument evaluation:")
	demoArguments()

	fmt.Println("\nFile example:")
	demoFile()

	fmt.Println("\nTimer example:")
	demoTimer()

	fmt.Println("\nPanic recover:")
	demoPanicRecover()

	fmt.Println("\nProgram continues...")
}
