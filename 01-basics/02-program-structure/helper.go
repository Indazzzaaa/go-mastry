package main

import "fmt"

func init() {
	fmt.Println("helper.go init() called")
}

//unexported
func sayHello() {
	fmt.Println("Hello from private function")
}

//exported
func SayPublicHello() {
	fmt.Println("Hello from public function")
}
