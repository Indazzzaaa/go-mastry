package main

import "fmt"

// multiple return values
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// named return
func sum(a, b int) (result int) {
	result = a + b
	return
}

// variadic function
func add(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// closure example
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {

	// multiple return
	res, err := divide(10, 2)
	fmt.Println("divide:", res, err)

	_, err2 := divide(10, 0)
	fmt.Println("divide error:", err2)

	// named return
	fmt.Println("sum:", sum(3, 4))

	// variadic
	fmt.Println("add:", add(1, 2, 3))
	nums := []int{4, 5, 6}
	fmt.Println("add slice:", add(nums...))

	// anonymous function
	func(msg string) {
		fmt.Println("anonymous:", msg)
	}("hello")

	double := func(x int) int {
		return x * 2
	}
	fmt.Println("double:", double(5))

	// closure
	c := counter()
	fmt.Println("counter:", c())
	fmt.Println("counter:", c())
	fmt.Println("counter:", c())
}
