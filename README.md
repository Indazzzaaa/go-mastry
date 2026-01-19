# Go Language Mastery – Complete Study Index

A structured, textbook-style roadmap to master the Go programming language from syntax to internals, backend engineering, and distributed systems.

You can follow this incrementally like a course or book.

---

## PART I — Go Grammar & Syntax Foundations

### 1. Introduction to Go ✅

- Why Go
- Go philosophy
- Go toolchain overview
- Hello World anatomy
- `go run`, `go build`, `go install`

### 2. Program Structure ✅

- Packages
- `main` package
- Imports
- Visibility rules (exported/unexported) => Capital Letters (public), small Letters (Private) to package
- `init()` function : Runs before main

### 3. Variables & Constants ✅

- `var` vs `:=`
- Multiple assignments
- Shadowing
- Constants & `iota`

### 4. Data Types ✅

- Numeric types (int, uint size depends on OS)
- Boolean
- String & rune (String is utf-8, rune is int32)
- Type conversion vs casting (Conversion not allowed in go, we have to case float32(a))
- Aliases vs defined types (aliases can be used and assigned to infer type, but defined types needs conversion )

### 5. Control Flow ✅

- `if`, `switch`
- `for` (all forms)
- `break`, `continue`, `goto`
- `defer` statement

### 6. Functions (First Class Citizens)✅

- Multiple return values
- Named returns
- Variadic functions
- Anonymous functions
- Closures

---

## PART II — Core Language Concepts

### 7. Pointers & Memory Basics ✅

- Address & dereference ( &, \*, default value is nil)
- Nil pointers ( var p \*int // defalut value of p is nil)
- Passing by value
- Pointer parameters
- Escape analysis (intro) // variable that used outside the scope it declared candidate for escapte analysis(so those will be stored on **Heap**)

### 8. Structs ✅

- Definition
- Initialization
- Anonymous structs
- Struct tags (metadata about struct fields)
- Composition

### 9. Methods & Receivers ✅

- Value receivers
- Pointer receivers
- Method sets
- Embedded methods

### 10. Arrays, Slices & Strings ✅

- Slice internals
- Capacity vs length
- Append behavior
- Copy semantics
- String immutability
- Make vs New
- Slice from Existing Array => arr[1:4] => [1,4)
  - s := arr[low:high: capacity]

### 11. Maps ✅

- Map operations
- Zero values
- Existence check
- Iteration order
- Map internals (intro)

---

## PART III — Interfaces & Abstractions

### 12. Interfaces ✅

- Implicit implementation => no implement, just create method and you have the interface.
- Interface values => (type, value) and if both nil then only interface will be nil.
- Empty interface => varxinterface{}
- Type assertions => value, ok := x.(int)
- Type switches => x.(type)

### 13. Interface Internals ✅

- itab
- Dynamic type/value
- Nil interface pitfalls

### 14. Generics (Go 1.18+) ✅

- Type parameters
- Constraints
- Comparable
- Type inference
- Generic interfaces
- Performance considerations

---

## PART IV — Error Handling & Resource Management

### 15. Error Handling ✅

- error interface
- Sentinel errors
- Custom errors
- Wrapping
- `errors.Is` / `errors.As`
- Panic & recover

### 16. Defer & Resource Management ✅

- Stack behavior
- Deferred execution order
- Common patterns

---

## PART V — Concurrency

### 17. Goroutines

- Creation
- Stack growth
- Scheduler basics

### 18. Channels

- Buffered vs unbuffered
- Closing rules
- Range over channels
- Select statement

### 19. Synchronization

- Mutex
- RWMutex
- WaitGroup
- Once
- Cond
- Atomic

### 20. Concurrency Patterns

- Worker pool
- Fan-in / fan-out
- Pipeline
- Cancellation with context
- Rate limiting
- Backpressure

---

## PART VI — Memory & Performance

### 21. Memory Model

- Stack vs heap
- Escape analysis
- Allocation patterns

### 22. Garbage Collector

- Tricolor algorithm
- STW phases
- Tuning GC

### 23. Profiling & Optimization

- Benchmarking
- pprof
- tracing
- CPU vs memory profiling
- `sync.Pool`
- False sharing

---

## PART VII — Standard Library Mastery

### 24. IO & Files

- os
- io
- bufio
- fs
- mmap

### 25. Networking

- net
- net/http
- TLS
- HTTP/2
- HTTP clients

### 26. Serialization

- json
- gob
- protobuf

### 27. Time & Scheduling

- time package
- Timers
- Tickers

---

## PART VIII — Tooling & Testing

### 28. Modules

- go mod
- semantic versioning
- replace
- vendor

### 29. Testing

- Unit tests
- Table tests
- Mocks
- Fuzzing
- Benchmarking

### 30. Toolchain

- go vet
- race detector
- golangci-lint
- Delve debugger
- Build tags
- Cross compile

---

## PART IX — Backend Engineering

### 31. HTTP Servers

- net/http internals
- Middleware
- Routing
- Graceful shutdown

### 32. Databases

- database/sql
- Connection pools
- Transactions
- Migrations
- ORM design

### 33. APIs

- REST design
- Pagination
- Filtering
- Authentication
- Authorization

### 34. Realtime

- WebSockets
- SSE
- gRPC

---

## PART X — Go Internals

### 35. Runtime

- Scheduler (GMP)
- Goroutine lifecycle

### 36. Compiler

- Parsing
- SSA
- Inlining
- Escape analysis

### 37. Data Structures Internals

- Slice layout
- Map hashing
- Channel internals
- Interface representation

---

## PART XI — Security

### 38. Secure Coding

- Race conditions
- SQL injection
- TLS config
- Secrets
- crypto/rand

---

## PART XII — Distributed Systems in Go

### 39. Distributed Concepts

- CAP theorem
- Consistency models
- Idempotency
- Backoff

### 40. Messaging Systems

- Kafka architecture
- Queues
- Pub/Sub

### 41. P2P Systems

- NAT traversal
- WebRTC
- BitTorrent

---

## PART XIII — Capstone Projects

### 42. Projects

- High performance HTTP server
- Distributed job queue
- WebSocket chat server
- Torrent-like file sharing system
- Key-value store (Redis-like)
- Kafka-like message broker
- Rate limiter library
- Custom logger
- Mini ORM
- RPC framework

---

## Suggested Usage

- Treat each **part as 1–2 weeks** of focused study
- Each chapter ≈ 1 day
- Build a small project after every major part
- Revisit internals & concurrency regularly

---

## Goal

By completing this index, you will reach **senior-level proficiency in Go** for:

- Backend engineering
- High performance systems
- Concurrency-heavy services
- Distributed systems
- System design interviews

---

Happy hacking 🚀
