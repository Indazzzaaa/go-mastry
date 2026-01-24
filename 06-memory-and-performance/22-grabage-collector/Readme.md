# Chapter 22 – Garbage Collector

## Type Of GC

Concurrent , tri-color, mark-and-sweep grabage collector

## Tricolor algorithm

- White: unreachable( grabage candidate)
- Gray: reachable, not scanned
- Black: reachable, scanned

Process: Start with root object(gray) -> scan gray(mark children black) -> Repeat until no gray objects-> all remaining white objects -> garbage.

Roots -> Gray -> Black -> Unreached -> White -> Collected.

## STW(Stop-The-World) phases

- Mark start
- Mark termination
- Sweep mostly concurrent

## Tuning

GOGC:

- 50 → low memory, more CPU
- 100 → default
- 200 → high memory, less CPU

## Tools

GODEBUG=gctrace=1
runtime.ReadMemStats()

## Reduce GC pressure

- Preallocate slices
- Reuse objects
- **sync.Pool : Object reuse to reduce GC pressue.**
- Avoid temp objects

# Rule : we don't fight the GC - we feed it fewer allocations.
