package main

import "fmt"

type User struct {
	Name string
}

// escapes: returning pointer
func createUser() *User {
	u := User{Name: "Alice"}
	return &u
}

// no escape
func createUserValue() User {
	u := User{Name: "Bob"}
	return u
}

// slice allocation
func buildSliceBad() []int {
	s := []int{}
	for i := 0; i < 1000; i++ {
		s = append(s, i)
	}
	return s
}

func buildSliceGood() []int {
	s := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		s = append(s, i)
	}
	return s
}

func main() {

	u1 := createUser()
	fmt.Println("user from heap:", u1.Name)

	u2 := createUserValue()
	fmt.Println("user from stack:", u2.Name)

	a := buildSliceBad()
	b := buildSliceGood()

	fmt.Println("slice sizes:", len(a), len(b))
}
