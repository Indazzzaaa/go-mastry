package main

import "fmt"

func init() {
	fmt.Println("main.go init() called")
}

func main() {
	fmt.Println("Inside main()")
	sayHello()
	SayPublicHello()
}
