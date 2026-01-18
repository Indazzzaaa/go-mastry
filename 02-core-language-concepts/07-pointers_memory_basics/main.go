package main

import "fmt"

func changeValue(x int) {
	x = 100
}

func changePointer(x *int) {
	*x = 100
}

func createPointer() *int {
	x := 42
	return &x // escapes to heap
}

func main() {

	// address & dereference
	a := 10
	p := &a

	fmt.Println("a:", a)
	fmt.Println("address of a:", p)
	fmt.Println("value via pointer:", *p)

	*p = 20
	fmt.Println("a after change:", a)

	// nil pointer
	var np *int
	fmt.Println("nil pointer:", np)

	if np == nil {
		fmt.Println("np is nil")
	}

	// pass by value
	b := 50
	changeValue(b)
	fmt.Println("after changeValue:", b)

	// pointer parameter
	changePointer(&b)
	fmt.Println("after changePointer:", b)

	// escape analysis
	ptr := createPointer()
	fmt.Println("escaped value:", *ptr)
}
