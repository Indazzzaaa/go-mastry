# Chapter 16 – Defer & Resource Management

## Defer

- `defer` schedules a function to run **after the surrounding function returns** .

```go
defer fmt.Println("bye")
fmt.Println("hello")

// output
hello
bye
```

## Stack behavior

- Deferred calls run in **Last-In-First-Out** order.

## Argument evaluation

- Arguments evaluated immediately
- ```go
  x := 10
  defer fmt.Println(x)
  x = 20
  // output : 10

  ```

## Common resource management patterns

- defer file.Close()
- defer mutex.Unlock()
- recover() => recover from panic

  ```go
  defer func() {
      if r := recover(); r != nil {
          fmt.Println("recovered")
      }
  }()

  ```

- timing functions

  ```go
  start := time.Now()
  defer func() {
      fmt.Println(time.Since(start))
  }()

  ```

## Warnings

- Defer in loops can delay resource release, because until defer functions runst it will capture the environment.
- Can increase memory usage

## Best practices

- Use defer immediately after acquiring resource
- Avoid defer in hot loops

## Stack behavior

- LIFO order

## Argument evaluation

- Arguments evaluated immediately

## Common patterns

- file.Close()
- mutex.Unlock()
- recover()
- timing functions

## Warnings

- Defer in loops can delay resource release
- Can increase memory usage

## Best practices

- Use defer immediately after acquiring resource
- Avoid defer in hot loops
