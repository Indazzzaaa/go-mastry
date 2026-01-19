package main

import "fmt"

// interface
type Speaker interface {
	Speak()
}

// struct implementing interface
type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Dog says: Woof")
}

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Cat says: Meow")
}

func describe(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}

func main() {

	// implicit implementation
	var s Speaker

	s = Dog{}
	s.Speak()

	s = Cat{}
	s.Speak()

	// interface value internals
	var s2 Speaker
	fmt.Println("s2 is nil?", s2 == nil)

	var d *Dog = nil
	s2 = d
	fmt.Println("s2 is nil after assigning nil Dog?", s2 == nil)

	// empty interface
	var x interface{}
	x = 10
	describe(x)

	x = "hello"
	describe(x)

	// type assertion
	value, ok := x.(string)
	fmt.Println("assert string:", value, ok)

	_, ok2 := x.(int)
	fmt.Println("assert int:", ok2)

	// type switch
	testType(100)
	testType("golang")
	testType(true)
}

func testType(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Println("int:", t)
	case string:
		fmt.Println("string:", t)
	default:
		fmt.Println("unknown type")
	}
}
