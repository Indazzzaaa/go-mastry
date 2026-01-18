package main

import "fmt"

// basic struct
type User struct {
	ID   int
	Name string
	Age  int
}

// struct with tags
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// composition
type Address struct {
	City  string
	State string
}

type Employee struct {
	Name string
	Address
}

func main() {

	// initialization (named)
	u1 := User{
		ID:   1,
		Name: "Alice",
		Age:  25,
	}

	fmt.Println("User1:", u1)

	// positional initialization
	u2 := User{2, "Bob", 30}
	fmt.Println("User2:", u2)

	// zero value struct
	var u3 User
	fmt.Println("User3 (zero value):", u3)

	// anonymous struct
	temp := struct {
		Key   string
		Value int
	}{
		Key:   "score",
		Value: 100,
	}

	fmt.Println("Anonymous struct:", temp)

	// struct tags example
	p := Product{ID: 101, Name: "Laptop", Price: 999.99}
	fmt.Println("Product:", p)

	// composition
	emp := Employee{
		Name: "Charlie",
		Address: Address{
			City:  "Mumbai",
			State: "MH",
		},
	}

	fmt.Println("Employee:", emp)
	fmt.Println("Employee city (promoted field):", emp.City)
}
