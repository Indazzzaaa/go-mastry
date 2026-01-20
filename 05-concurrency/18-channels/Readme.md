# Chapter 18 – Channels

## Channel basics

- Used to communicate between goroutines

  ```go
  ch := make(chan int) // create channel
  ch <- 10 // send data into channel
  x := <-ch // Receive data from channel
  ```

## Unbuffered channels

- Send and receive block(data) until both ready
- Used for synchronization

  ```go
  ch := make(chan int)
  ```

## Buffered channels

- Can hold N values
- Send blocks only when buffer full
- Receive blocks when empty

  ```go
  ch := make(chan int, 3)
  ```

## Closing rules

- Only sender closes, Never close from receiver
- Never close twice
- Never send after close otherwise Panic .
- Receiving from closed channed -> Zero value + ok = false => v, ok := <-ch

## Range over channels

- Automatically stops when channel closed and drained.

  ```go
  for v := range ch {
      fmt.Println(v)
  }

  ```

## Select statement

- Waits on multiple channels
- First ready case executes

  ```go
  select {
  case v := <-ch1:
  case v := <-ch2:
  default:
  }

  ```

## Rules

- Don't communicate by sharing memory
- Share memory by communicating

# Practice tasks

1. Demonstrate deadlock with unbuffered channel
2. Create worker using buffered channel
3. Detect closed channel using `ok`
4. Use select with timeout
5. Implement producer–consumer
