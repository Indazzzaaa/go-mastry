package main

import "fmt"

// generic function
func Print[T any](v T) {
	fmt.Println(v)
}

// constraint
type Number interface {
	int | float64
}

func Add[T Number](a, b T) T {
	return a + b
}

// comparable constraint
func Contains[T comparable](arr []T, v T) bool {
	for _, x := range arr {
		if x == v {
			return true
		}
	}
	return false
}

// generic struct
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() T {
	n := len(s.items)
	val := s.items[n-1]
	s.items = s.items[:n-1]
	return val
}

func main() {

	// type parameter
	Print(10)
	Print("golang")

	// constraint
	fmt.Println("Add int:", Add(3, 4))
	fmt.Println("Add float:", Add(1.5, 2.5))

	// comparable
	nums := []int{1, 2, 3}
	fmt.Println("Contains 2:", Contains(nums, 2))

	words := []string{"go", "rust", "java"}
	fmt.Println("Contains rust:", Contains(words, "rust"))

	// generic struct
	var s Stack[int]
	s.Push(10)
	s.Push(20)
	fmt.Println("Pop:", s.Pop())
}
