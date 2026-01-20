## 1. What a channel really is?

- A channel is: A thread-safe queue + synchronization primitive managed by the Go Scheduler.
- It does two jobs:
  - Transfers data
  - Coordinates goroutines (blocking / waking) we'll see how?

## 2. Unbuffered Channel (capacity =0)

- ch := make(chan int) // or make(chan int, 0)

### Behavior

| Operation              | Result               |
| ---------------------- | -------------------- |
| send without receiver  | blocks               |
| receive without sender | blocks               |
| both present           | data copied directly |

```go
ch := make(chan int)

go func() {
    ch <- 10       // blocks until someone receives
    fmt.Println("sent")
}()

fmt.Println(<-ch)  // unblocks sender

// Output Order
10
sent

```

Key Properties

- Synchronous
- Strong ordering
- Perfect for handshakes
- Slower under heavy throughput
- can deadlcok easily if misued.

## 3. Buffered channel (capacity > 0)

- ch := make(chan int, 3)

### Behavior

| Operation                     | Result       |
| ----------------------------- | ------------ |
| send when buffer not full     | non-blocking |
| send when buffer full         | blocks       |
| receive when buffer not empty | non-blocking |
| receive when empty            | blocks       |

```go
ch := make(chan int, 2)

ch <- 1
ch <- 2        // both succeed immediately

go func() {
    ch <- 3    // blocks (buffer full)
}()

fmt.Println(<-ch) // frees space

```

Key Properties:

- Asynchronous
- Higher Throughput
- Reduces goroutine blocking
- More memory
- Can hide bugs
- Ordering issues if misused

# 4. Scheduler-level difference (important)

### Unbuffered

Scheduler must:

- Park sender goroutine
- Wake receiver
- Perform direct handoff
- Resume sender

More context switching.

---

### Buffered

Scheduler:

- Push to queue
- Continue running sender
- Receiver handles later

Less scheduling overhead.

## 5. When to use What (real-world rules)

- Use UNBUFFERD when
  - You want synchronization : since it blocks other.
  - Ordering matters strictly: request -> process -> repond
  - Backpressure is important: Force producers to slow down.
  - Low volume / correctness critical
- Use BUFFERED When :
  - Produer faster then consumer
  - Worker pools : jobs: = make(chan Job, 100)
  - I/O pipelines
  - Avoid groroutine explosion

## 6. Common Mistakes

- Thinking buffered channels never blocks : They block when full.
- Using huge buffers to "fix" deadlocks : This hides bugs.
- Using unbuffered channels in high throughput systems : causes unnecessary context switching overheads.
- Forgetting to close channels : Receiver block forever.
