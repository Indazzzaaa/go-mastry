package main

import "fmt"

func main() {

	// if
	x := 10
	if x > 5 {
		fmt.Println("x > 5")
	}

	// if with short statement
	if n := len("golang"); n > 5 {
		fmt.Println("length > 5:", n)
	}

	// switch
	day := 2
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Unknown day")
	}

	// expression-less switch
	num := -3
	switch {
	case num > 0:
		fmt.Println("Positive")
	case num < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}

	// for classic
	for i := 0; i < 3; i++ {
		fmt.Println("i =", i)
	}

	// while-style
	count := 0
	for count < 3 {
		fmt.Println("count =", count)
		count++
	}

	// range loop
	arr := []int{10, 20, 30}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// break & continue
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		if i == 4 {
			break
		}
		fmt.Println("loop:", i)
	}

	// goto
	goto END

	fmt.Println("This will be skipped")

END:
	fmt.Println("Reached END label")

	// defer
	defer fmt.Println("World")
	defer fmt.Println("Hello")

	fmt.Println("Function ending")
}
