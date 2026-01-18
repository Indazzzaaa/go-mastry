# Chapter 11 – Maps

## Operations

- Create with make or literal => m := make(map[string]int)
- Insert, read, update, delete

  ```go
  m := map[string]int{
      "a": 1,
      "b": 2,
  }

  // Insert  or Update
  m["c"] = 3

  // Read
  v := m["a"]

  // Delete
  delete(m, "a")

  ```

## Zero value

- nil map
- Can read
- Cannot write

  ```go
  // Zero Values
  var m map[string]int // Value = nil, Read -> Ok, Write -> Panic ❌
  Fix : m = make(map[string]int)


  ```

## Existence check

- value, ok := m[key]

## Iteration order

- Random
- Never rely on order

  ```go
  for k, v := range m {}

  ```

## Internals (intro)

- Hash table
- Buckets
- Open Addressing
- Rehashing on Growth
- Not thread-safe

## Rules

- Always initialize map before writing
- Use comma-ok for safety
