package main

import "fmt"

func main() {

	// array
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("array:", arr)

	// slice
	s := []int{10, 20, 30}
	fmt.Println("slice:", s)

	// len vs cap
	a := make([]int, 2, 5)
	fmt.Println("len:", len(a), "cap:", cap(a))

	// append behavior
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	fmt.Println("after append:", a, "len:", len(a), "cap:", cap(a))

	// sharing
	x := []int{1, 2, 3}
	y := x
	y[0] = 99
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	// copy safely
	z := make([]int, len(x))
	copy(z, x)
	z[1] = 777
	fmt.Println("x after copy:", x)
	fmt.Println("z:", z)

	// string immutability
	str := "hello"
	fmt.Println("string:", str)

	bytes := []byte(str)
	bytes[0] = 'H'
	str = string(bytes)

	fmt.Println("modified string:", str)
}
