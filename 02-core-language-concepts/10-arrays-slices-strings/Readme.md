# Chapter 10 – Arrays, Slices & Strings

## Arrays

- Fixed size => var a [3]int
- Rarely used directly

## Slices

- Dynamic => s := []int{1,2,3}
- Backed by array
- Has pointer, length, capacity

```go
// Slice Internals
type slice struct {
    ptr *array
    len int
    cap int
}

s := make([]int, 2, 5) // create slice of 2 lenght, 5 capacity
len(s)  // 2
cap(s) // 5

// Creating slice with existing array
arr := [5]int{10, 20, 30, 40, 50}
s := arr[1:4] => s = [20, 30, 40]


```

## len vs cap

- len = usable elements
- cap = allocated space

## append

- Extends slice
- May allocate new array

  ```go
  s = append(s, 10)
  // if capacity allows -> same array
  // if not -> new array allocated
  // Consequences : Append may break sharing so we might think we are mutating orginal array but it may be different all together.
  ```

## Copy semantics

- Slices are reference-like
- Slice assignment shares same array
- Use copy() for deep copy

  ```go
  a := []int{1, 2}
  b := a
  b[0] = 99 // both changes.

  // To Copy safely
  c := make([]int, len(a))
  copy(c, a)


  ```

## Strings

- Immutable
- UTF-8 bytes
- Convert to []byte to modify

  ```go
  s := "hello"
  s[0] = 'H' // ❌ not allowed

  // Convert before modifying
  b := []byte(s)
  b[0] = 'H'
  s = string(b)

  ```

## Make vs New

- New(T)
  - Allocates zeroed memory for type T => means all variables will get default values.
  - And Returns \*T (pointer)

    ```go
    p : new(int)
    ```

  - Can be used with any type which we want to allocate on heap.(Do not use with map, channels and slice because it will assigne internal pointers to nil.)

- make(T, ...)
  - Used only for slice, map, channels
  - It allocates and initialize internal strucutre(unlike new which only initialize with zero values)

    ```go
    s := make([]int, 3)
    m := make(map[string]int)
    c := make(chan int)

    ```

  - Returns **value** , not pointer.

## Rules

- Pass slices carefully
- Never assume append keeps same memory
- Copy when needed
