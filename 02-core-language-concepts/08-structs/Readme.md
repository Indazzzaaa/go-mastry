# Chapter 8 – Structs

## Definition

Struct groups related fields.

```go
type User struct {
    ID   int
    Name string
    Age  int
}

```

## Initialization

- Named fields (recommended)
- Positional
- Zero value

```go
// Named Fileds
u := User{
    ID:   1,
    Name: "Alice",
    Age:  25,
}

// Positional
u := User{1, "Alice", 25}

// Zero(Default) value
var u User // in this case all its fields will get the default value.(eg. int : 0, string: "" etc.)

```

## Anonymous Structs

- No type name
- Useful for temporary data

  ```go
  person := struct {
      Name string
      Age  int
  }{
      Name: "Bob",
      Age:  30,
  }

  ```

## Struct Tags

- Metadata for fields
- Used by JSON, DB, validation, etc.

  ```go
  type User struct {
      ID   int    `json:"id"`
      Name string `json:"name"`
  }
  // Access using reflection (later topic)
  ```

## Composition

- Embed structs
- Fields are promoted (Like they are part of same struct)
- Preferred over inheritance

  ```go
  type Address struct {
      City  string
      State string
  }

  type User struct {
      Name string
      Address  // Embedding => u.City
      AnotherAddress Address  // this is not emedding, u.AnotherAddress.City
  }

  ```

## Rules

- Structs are value types
- Copying struct copies all fields
- Use pointers for large structs
