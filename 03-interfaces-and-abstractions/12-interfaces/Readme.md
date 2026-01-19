# Chapter 12 – Interfaces (Basics)

## Implicit Implementation

- No "implements" keyword
- If methods match → interface implemented
  - A type implements an interface automatically if it has the methods.

  ```go
  type Speaker interface {
      Speak()
  }

  type Dog struct{}

  func (d Dog) Speak() {
      fmt.Println("Woof")
  }

  ```

  - Dog Implement Speaker Implicitly

## Interface Values

- Stored as (type, value)
- Interface can be non-nil even if underlying value is nil

  ```go
  var s Speaker
  s = Dog{}

  // Internally
  (type = Dog, value = Dog{})

  ```

  If both are nil → interface is nil

  If type exists but value nil → interface is NOT nil (important pitfall)

## Empty Interface

- interface{} accepts any type
- Used for generic behavior

  ```go
  var x interface{}

  x = 10
  x = "hello"
  x = true

  ```

## Type Assertions

- v := i.(T)
- v, ok := i.(T) // safe

## Type Switch

- switch v := i.(type)

## Important Pitfall

```go
var d *Dog = nil
var s Speaker = d

s == nil  // false ❗
(type = *Dog, value = nil)

```

## Rules

- Prefer small interfaces
- Accept interfaces, return structs
