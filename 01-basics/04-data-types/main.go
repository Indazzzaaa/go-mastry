package main

import "fmt"

// alias
type MyAlias = int

// defined type
type MyInt int

func main() {

	// numeric types
	var a int = 10
	var b int64 = 20
	var c float64 = 3.14

	fmt.Println("Numbers:", a, b, c)

	// boolean
	var ok bool = true
	fmt.Println("Boolean:", ok)

	// string
	s := "Hello"
	fmt.Println("String:", s)

	// rune
	r := '世'
	fmt.Println("Rune:", r)
	fmt.Printf("Rune as char: %c\n", r)

	// string indexing
	word := "Go"
	fmt.Println("First byte:", word[0])

	// conversion
	var x int = 100
	var y float64 = float64(x)
	fmt.Println("Converted:", y)

	// alias type
	var m MyAlias = 50
	var n int = m
	fmt.Println("Alias works:", n)

	// defined type
	var p MyInt = 60
	//! var q int = p   // ❌ error
	var q int = int(p) // ✅ conversion
	fmt.Println("Defined type:", q)
}
