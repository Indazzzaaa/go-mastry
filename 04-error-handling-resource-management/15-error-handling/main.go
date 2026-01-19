package main

//! Errors are values. Handle them explicitly.
import (
	"errors"
	"fmt"
)

// sentinel error
var ErrNotFound = errors.New("not found")

// custom error
type MyError struct {
	Code int
	Msg  string
}

func (e MyError) Error() string {
	return fmt.Sprintf("code=%d msg=%s", e.Code, e.Msg)
}

func findUser(id int) error {
	if id == 0 {
		return ErrNotFound
	}
	if id < 0 {
		return MyError{Code: 400, Msg: "invalid id"}
	}
	return nil
}

func wrapper() error {
	err := findUser(0)
	if err != nil {
		return fmt.Errorf("wrapper failed: %w", err)
	}
	return nil
}

func causePanic() {
	panic("something went wrong")
}

func safeCall() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	causePanic()
}

func main() {

	err := wrapper()

	// errors.Is
	if errors.Is(err, ErrNotFound) {
		fmt.Println("user not found")
	}

	// errors.As
	var myErr MyError
	if errors.As(err, &myErr) {
		fmt.Println("custom error:", myErr)
	}

	fmt.Println("error:", err)

	// panic & recover
	safeCall()

	fmt.Println("program continues...")
}
