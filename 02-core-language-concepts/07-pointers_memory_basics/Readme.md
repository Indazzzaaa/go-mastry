# Chapter 7 – Pointers & Memory Basics

## Address & Dereference

- &x → address
- \*p → value
- We can use pointers in go but go does not allows the pointer arithmetics.

  ```
  variable → value
  pointer → address → value

  ```

## Nil pointers

- Default value is nil : `var p *int // default value of p is ` nil.
- Dereferencing nil causes panic

## Passing by value

- Go copies values by default

## Pointer parameters

- Allows modifying original variable
- Avoids copying large structs

## Escape analysis

- Compiler decides stack vs heap
- Returning local address causes heap allocation

```go
func create() *int {
    x := 10
    return &x
}
// x escapes → allocated on heap.
```

## Commands

go build -gcflags="-m" => To see the escaped variables

## Rules

- Use pointers when:
  - You want to modify original value
  - Avoid copying large structs
  - Represent optional values
