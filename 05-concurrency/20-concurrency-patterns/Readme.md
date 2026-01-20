# Chapter 20 – Concurrency Patterns

### These patterns are used in:

- Web Servers
- Job queues
- Kafka consumers
- Crawlers
- Data pipelines
- P2P systems

## Worker Pool

- Process many jobs using limited workers

  ```go
  jobs → [ workers ] → results

  ```

## Fan-out / Fan-in

- Distribute work to goroutines
- Collect results into one channel

## Pipeline

- Multi-stage processing
- Each stage runs concurrently

## Context cancellation

- Graceful stop => ctx, cancel := context.WithCancel(context.Background())
- Timeouts
- Deadlines
- Used In
  - HTTP Servers
  - DB calls
  - Workers
  - APIs
- Important Properties
  - Cancellation is Cooperative : Goroutine must check ctx.Done()
  - Cancellation is broadCast : Closing channel wakes everyone.
  - No Force Kill : Go cannot kill goroutines.

## Rate limiting

- Control request rate
- time.Ticker or token bucket

## Backpressure

- Slow consumers block producers
- Buffered channels help control flow

## Used in

- Servers
- Job systems
- Stream processing
- Distributed systems
