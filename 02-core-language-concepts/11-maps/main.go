package main

import "fmt"

func main() {

	// create map
	scores := make(map[string]int)

	// insert
	scores["alice"] = 90
	scores["bob"] = 80

	// read
	fmt.Println("alice:", scores["alice"])

	// update
	scores["alice"] = 95

	// delete
	delete(scores, "bob")

	fmt.Println("map:", scores)

	// zero value map
	var m map[string]int
	fmt.Println("nil map read:", m["x"])

	// existence check
	val, ok := scores["alice"]
	fmt.Println("alice exists?", ok, "value:", val)

	val2, ok2 := scores["bob"]
	fmt.Println("bob exists?", ok2, "value:", val2)

	// iteration order
	scores["charlie"] = 85
	scores["david"] = 88

	fmt.Println("Iterating map:")
	for k, v := range scores {
		fmt.Println(k, v)
	}
}
