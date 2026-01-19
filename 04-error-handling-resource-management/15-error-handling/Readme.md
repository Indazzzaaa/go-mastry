# Chapter 15 – Error Handling

## error interface

- Any type implementing Error() string is an error.

```go
type error interface {
    Error() string
}

```

## Sentinel errors

Predefined errors we compare againts

```go
var ErrNotFound = errors.New("not found")
// usage
if err == ErrNotFound {}

```

## Custom errors

- Deffine our own error types by implementing Error() string method

  ```go
  type MyError struct {
      Code int
      Msg  string
  }

  func (e MyError) Error() string {
      return e.Msg
  }

  ```

## Wrapping

- Add context while preseving original error

  `return fmt.Errorf("open file: %w", err) `

## errors.Is / errors.As

- Is → check specific error
- As → extract error type

  ```go
  errors.Is(err, ErrNotFound)
  var myErr MyError
  errors.As(err, &myErr)

  ```

## Panic & recover

- Panic for programmer bugs
- Recover in top-level only

  ```go
  defer func() {
      if r := recover(); r != nil {
          fmt.Println("recovered:", r)
      }
  }()

  ```

## Best practices

- Return errors, don't panic
- Add context when wrapping
- Avoid string comparisons
