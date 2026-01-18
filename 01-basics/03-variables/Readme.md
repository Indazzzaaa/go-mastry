# Chapter 3 – Variables & Constants

## var vs :=

- var can be used globally and locally => var x int = 10
- := only inside functions => z := 30
- := requires at least one new variable

## Multiple assignment

- a, b := 1, 2
- a, b = b, a

## Shadowing

- Inner scope redeclares variable
- Can cause bugs if not careful

## Constants

- Compile-time values
- Untyped by default, eg. const A = 10

## iota

- Auto-incrementing constant generator
- Useful for enums and flags

  ```go
  const (
      A = iota // 0
      B        // 1
      C        // 2
  )


  // Advanced Patterns
  const (
  		KB = 1 << (10 * iota) // iota: 0,
  		MB                    // 1
  		GB                    // 2
  	)

  ```

## Key rule

Prefer := for short-lived local variables.
Use var for globals and clarity.
