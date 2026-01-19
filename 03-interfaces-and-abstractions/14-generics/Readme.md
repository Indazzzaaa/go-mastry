# Chapter 14 – Generics

## Type Parameters

```go
func Print[T any](v T) {
    fmt.Println(v)
}

// Call
Print("Hello") // inferred
```

## Constraints

- Limit what types are allowed

  ```go
  type Number interface {
      int | float64
  }

  func Add[T Number](a, b T) T {
      return a + b
  }

  ```

## comparable

- Built-in constraint this allows ==, !=

  ```go
  func Contains[T comparable](arr []T, v T) bool

  ```

  - Used for map keys and equality checks types.

## Type Inference

Compiler deduces types automatically

```go
Add(3, 4)        // inferred as int
Add(1.2, 2.3)    // inferred as float64

```

## Generic Interfaces

- Interfaces can be generic too

  ```go
  type Stack[T any] interface {
      Push(T)
      Pop() T
  }

  // Implementation
  type IntStack struct { ... }
  // or
  type Box[T any] struct {
      Value T
  }

  ```

## Performance

- Go generics use **monomorphization** (mostly):

* Compiler generates specialized code per type
* Almost zero runtime cost
* Slightly larger binary

## When to use generics

- Algorithms
- Data structures
- Utilities
- Libraries

## When NOT to use

- Simple business logic
- Single type usage
