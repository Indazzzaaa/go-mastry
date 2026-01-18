package main

import "fmt"

type User struct {
	Name string
}

// value receiver
func (u User) Greet() {
	fmt.Println("Hello,", u.Name)
}

// pointer receiver
func (u *User) SetName(name string) {
	u.Name = name
}

// embedded example
type Animal struct{}

func (a Animal) Speak() {
	fmt.Println("Animal speaks")
}

type Dog struct {
	Animal
}

func main() {

	u := User{Name: "Alice"}

	// value receiver
	u.Greet()

	// pointer receiver
	u.SetName("Bob")
	u.Greet()

	// method sets
	var u2 User
	var p *User = &u2

	u2.SetName("Charlie") // compiler auto-takes address
	p.Greet()             // compiler auto-dereferences

	fmt.Println("u2:", u2.Name)

	// embedded methods
	d := Dog{}
	d.Speak() // promoted from Animal
}
