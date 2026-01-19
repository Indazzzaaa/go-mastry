package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}

func (d *Dog) Speak() {
	fmt.Println("woof")
}

func checkNil(i interface{}) {
	if i == nil {
		fmt.Println("interface is nil")
	} else {
		fmt.Printf("interface NOT nil | type=%T | value=%v\n", i, i)
	}
}

func main() {

	// case 1: nil interface
	var s Speaker
	fmt.Println("s == nil ?", s == nil)

	// case 2: typed nil
	var d *Dog = nil
	var s2 Speaker = d

	fmt.Println("s2 == nil ?", s2 == nil)

	checkNil(s)
	checkNil(s2)

	// dynamic type/value demo
	var x interface{}

	x = 42
	fmt.Printf("type=%T value=%v\n", x, x)

	x = "golang"
	fmt.Printf("type=%T value=%v\n", x, x)
}
