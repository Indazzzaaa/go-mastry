## 1. What is a cache-line?

- Modern CPUs don't load memory byte-by-byte, they load memory in fixed-size blocks called cache lines.(Typical size: 64bytes).

  ```
  Registers (fastest)
  L1 cache
  L2 cache
  L3 cache
  RAM (slowest)

  >> Data moves in cache line chunks between these layers

  // so if we read
  x := arr[0]

  // CPU actually loads into cache
  arr[0] ... arr[63]   (64 bytes)


  ```

## 2. Cache coherence problem

- Each CPU core has it own cache. When multiple cores modify memory: CPU must keep caches coherent using protocols like MESI(Modifed, Exclusive, Shared, Invalid).

## 3. What happen when goroutines modify the same cache

- This causes false sharing.( Memory logically independent but physically shared at cache-line level)

  ```
  type Data struct {
      a int64  // goroutine 1 writes
      b int64  // goroutine 2 writes
  } // both a, and b are next to each other in same cache line

  // What CPU does
  1. Goroutine 1 (Core 1) => Modifies a (cache line becomes Modified in core 1)
  2. Goroutine 2 (Core 2) => Modifies b (CPU invalidates cache line in core 1, fetches line into core 2)
  3. Then core 1 write again => Invalidates core 2 , ping-pong of cache line between cores.

  ```

  - Event though goroutines mare modifying different variables they are causing invalidate of another and side effect is poor performance.
    - Massive cache invalidations
    - memory bus traffic
    - pipeline stalls
    - performance collapse

## 4. Performance impact

- Separate cache lines : 500million (ops/sec)
- Same cache lines : 20million(ops/sec)

## 5. How to fix it?

- Padding

  ```go
  type Data struct {
      a int64
      _ [56]byte  // pad to 64 bytes
      b int64
  }

  ```

  - Use cache line alignment

    ```go
    type Data struct {
        a int64
        _ [cacheLineSize]byte
        b int64
    }

    ```

  - Use per-goroutine data : Instead of shared counters then sum later.

    ```go
    counters := make([]int64, runtime.GOMAXPROCS())

    ```
