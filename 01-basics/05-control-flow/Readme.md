# Chapter 5 – Control Flow

## if, else if, else

```go
if x > 10 {
}

// with short statement
if n := len(s); n > 5 {
}

```

## switch

- No **break** needed
- Supports expression-less form
- Supports type switch

  ```go
  switch day {
  case 1:
      fmt.Println("Monday")
  case 2:
      fmt.Println("Tuesday")
  default:
      fmt.Println("Unknown")
  }

  // Expression Less
  switch {
  case x > 0:
  case x < 0:
  }

  // Type Switch
  switch v := i.(type) {
  case int:
  case string:
  }



  ```

## for

- Only loop in Go
- Supports classic, while, infinite, range forms

  ```go
  for i := 0; i < 5; i++ {} // classic for loop
  for x < 10 {} // while-style
  for {} // infinite loop
  for i, v := range arr {} // Range

  ```

## break & continue

- Can be labeled

  ```go
  // GO Supports Labeled break
  outer:
  for {
      for {
          break outer
      }
  }

  ```

## goto

- Rarely used
- Jumps to label within function

## defer

- Executes after function returns
- LIFO order
- Common for cleanup (files, locks)

  ```go
  defer fmt.Println(1)
  defer fmt.Println(2)

  ```

## Common uses of defer

- Close files
- Unlock mutex
- Recover from panic
