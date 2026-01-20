# Chapter 19 – Synchronization

## Mutex

- Used to Protects shared data
- Exclusive lock => Only once goroutine can hold the lock.

  ```go
  var mu sync.Mutex

  mu.Lock()
  shared++
  mu.Unlock()

  ```

## RWMutex ( sync. RWMutex)

Two types of locks: 1. Read lock (RLock) -> many readers allowed, 2. Write lock(Lock) -> Exclusive

```go
mu.RLock()
mu.RUnlock()

mu.Lock()
mu.Unlock()

```

## WaitGroup

- Wait for goroutines to finish

  ```go
  wg.Add(1)
  go func() {
      defer wg.Done()
  }()
  wg.Wait()

  ```

## Once

- Execute code exactly once (thead safe)
- Used for lazy initialization.

  ```go
  once.Do(initFunc)
  ```

## Cond

- Goroutine coordination via signals
- Used when goroutines must wait for some condition

  ```go
  cond.Wait()
  cond.Signal()
  cond.Broadcast()

  ```

## Atomic

- Lock-free operations
- Used for counters and flags
- Faster than mutex for counters.

  ```
  atomic.AddInt64(&x, 1)
  atomic.LoadInt64(&x)

  ```

## Rules

- Prefer channels first
- Use mutex for shared state
- Use atomic for simple counters
