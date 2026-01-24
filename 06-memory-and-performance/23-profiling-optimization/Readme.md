# Chapter 23 – Profiling & Optimization

## Benchmarking

go test -bench=. -benchmem

## pprof

CPU: -cpuprofile
MEM: -memprofile

## Tracing

go test -trace trace.out

## CPU vs Memory profiling

CPU → slow code
MEM → allocations
Block → channel/lock waits
Mutex → lock contention

## sync.Pool

- Reuse objects
- Reduce GC pressure
- Not guaranteed to retain objects

## False sharing

- **Multiple goroutines modify same cache line : Read on it.**
- Fix using padding

## Rules

- Measure before optimizing
- Reduce allocations
- Preallocate slices
- Avoid interfaces in hot paths

Benchmark → find slow code
Profile → find why
Optimize → reduce allocations
Reuse memory → reduce GC
Avoid cache contention

### Truth : The fastest code is the code that allocates nothing.
