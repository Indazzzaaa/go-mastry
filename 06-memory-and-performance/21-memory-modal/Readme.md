# Chapter 21 – Memory Model

## Stack

- Fast, Automatic cleanup, No Fragmentation
- OS creates stack for each thread.
- OS gives virtual space for theads (ie. 8MB for linux, 1MB for windows) but allocates small (1-2 pages ) of RAM initially.
- Stack grows from higher memory address to lower.
- **Segmentation Fault** -> Guard Pages(At the end of Stack range i.e 8MB for LINUX) -> stack pointer enters -> OS interrups -> Segementation Fault.

## Heap

- Managed by GC, Slower, Shared
- OS intially gives Virtual space to program and when actuall address request and not present( Page Fault + Page allocation)
- OS allocates pages in quata of (4kB, 8KB, 32KB) -> Then language runtime uses that pages to allocate the request space for heap variable.
- Leads to internal fragmenation(required 13bytes, runtimes allocates 16 just to fullfill the alignment thing)
- Leads to external Fragmentation (require 64 bytes, but there are many holes of 32 but not any contiguous hole of 64).
- Heap grows from lower memory address to higher.
- Go Heap Structure : ( Pages -> 8KB(go uses large pages, it go internal OS will have 4Kb and go will ask 2\*4KB .), Span-> group of pages, Arena-> 64MB)
- Flow:
  - We do : make([] int, 10)
  - Go allocators looks in the spans -> if not available -> Request mores pages from OS using mmap -> OS maps pages to program -> Go subdivides and retures the asked(taking care of alignment)

## Escape Analysis

Compiler Checks : can this variable safely stay on stack? => if no -> moved to heap

Common escape cases:

- Returning pointer
- Stored in interface
- Used/Captured in goroutine
- Global reference

Inspect escape analysis => go build -gcflags="-m"

## Allocation patterns

Bad:
append to nil slice repeatedly

Good:
make with capacity

## Tools

go build -gcflags=\"-m\"

## Rules

- Avoid unnecessary pointers
- Preallocate slices
- Reduce heap allocations
- Measure with pprof
