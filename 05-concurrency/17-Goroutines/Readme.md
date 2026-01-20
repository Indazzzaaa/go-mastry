# Chapter 17 – Goroutines

## Creation

- Use `go` keyword => go myFunction()
- Can run functions or anonymous functions => go func() {...}()

## Stack growth

- OS Threads : Fixed stack( usually 1-8MB)
- Goroutines : Very small (~2 KB), Grows and shrink dynamically
- Benefit : Create milliions of goroutines, very memory efficient.

## Scheduler basics

- G = goroutine
- M = OS thread
- P = processor (logical)

  G runs on M using P

- Schedular Handles: Preemption, Context Switching, Work Stealing

## Properties

- Lightweight
- Managed by runtime
- Cheap to create

## Warnings

- Program exits when main exits
- Goroutines are not threads

## More on Go Scheduler Queues

| Queue                                       | Scope                                 | Stores               |
| ------------------------------------------- | ------------------------------------- | -------------------- |
| Local Run Queue(LRP)                        | Each P (processor ) has its own queue | Runnable goroutines  |
| Global Run Queue                            | Global                                | Runnable goroutines  |
| Netpoll Queue                               | Global                                | I/O ready goroutines |
| Timer Heap (time.sleep, after, withTimeout) | Per P                                 | Sleeping goroutines  |
| Channel sendq                               | Per channel                           | Blocked senders      |
| Channel recvq                               | Per channel                           | Blocked receivers    |
| Mutex wait queues                           | Per lock                              | Blocked goroutines   |
| GC queues                                   | GC system                             | GC workers           |

# 9. Scheduling flow (simplified)

When a P needs work:

1. Check **local run queue**
2. Check **global run queue**
3. Poll **netpoll queue**
4. Steal from other P
5. Park thread
