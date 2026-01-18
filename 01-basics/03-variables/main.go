package main

import "fmt"

var globalCount int = 0

// anotherCount := 0; //! Error

func main() {

	// var vs :=
	var a int = 10
	b := 20

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// multiple assignment
	x, y := 1, 2
	fmt.Println("before swap:", x, y)

	x, y = y, x
	fmt.Println("after swap:", x, y)

	// shadowing
	value := 100
	fmt.Println("outer value:", value)

	{
		value := 200
		fmt.Println("inner value:", value)
	}

	fmt.Println("outer value again:", value)

	// constants
	const Pi = 3.14159
	fmt.Println("Pi =", Pi)

	// iota
	const (
		Red = iota
		Green
		Blue
	)

	fmt.Println("Colors:", Red, Green, Blue)

	const (
		KB = 1 << (10 * iota) // iota: 0,
		MB                    // 1
		GB                    // 2
	)

	fmt.Println("Sizes:", KB, MB, GB)

	// Bit masking flags
	const (
		Read    = 1 << iota // 0001 (1)
		Write               // 0010 (2)
		Execute             // 0100 (4)
	)
}
