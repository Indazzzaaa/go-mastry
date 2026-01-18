# Chapter 6 – Functions

**Note : Functions are first class citizens in go. That mens ( 1. Assign functions to variables, 2. Pass functions as arguments, 3. Return functions from functions, 4. Store functions in data structures, 5. Use closures(Capture variables))**

## Multiple return values

- Common for returning value + error

  ```go
  func divide(a, b int) (int, error) {
      if b == 0 {
          return 0, fmt.Errorf("division by zero")
      }
      return a / b, nil
  }

  ```

## Named returns

- Return variables declared in signature
- return without values uses them

  ```go
  func sum(a, b int) (result int) {
      result = a + b
      return
  }

  ```

## Variadic functions

- Accept ...type
- Treated as slice inside function

  ```go
  func add(nums ...int) int {
      total := 0
      for _, n := range nums {
          total += n
      }
      return total
  }

  // calling it
  add(1,2,3)
  add(arr...) // pass slice

  ```

## Anonymous functions

- Functions without names
- Can be assigned to variables

  ```go
  func() {
      fmt.Println("Hello")
  }()

  // Functions are first class citizens in go
  f := func(x int) int {
      return x * 2
  }


  ```

## Closures

- Capture outer variables
- Preserve state between calls
- ```go
  func counter() func() int {
      count := 0
      return func() int {
          count++
          return count
      }
  }
  ```

## Common uses

- Callbacks
- Middleware
- Goroutines
- Encapsulation
