# Module 3: Threads in Go

## Topic 3.1: Goroutines

- Go provides **built-in concurrency support** via Goroutines.
- A **Goroutine** is a lightweight thread managed by the Go runtime.

### Key Concepts

- Every Go program starts with a **main Goroutine** (executing `main()`).
- To create additional Goroutines, use the `go` keyword.

### Goroutine Example

```go
func foo() {
    // some task
}

func main() {
    a := 1          // main Goroutine
    go foo()        // creates a new Goroutine to run foo()
    a = 2           // continues main Goroutine
}
```

- `go foo()` creates a **new Goroutine** running `foo()` concurrently.
- Execution of `a = 2` continues **immediately**, regardless of whether `foo()` has completed.

### Blocking vs Non-Blocking

| Style        | Description                                                    |
| ------------ | -------------------------------------------------------------- |
| Without `go` | Main blocks on `foo()`; next line waits until `foo()` returns. |
| With `go`    | `foo()` runs concurrently; next line runs immediately.         |

### Goroutine Lifecycle

- A Goroutine ends when its function completes.
- **If main Goroutine ends**, all other Goroutines are **forced to terminate**, even if they haven't finished.

---

## Topic 3.2: Exiting Goroutines

### Problem

- If the `main()` ends before other Goroutines finish, they are killed.

### Example

```go
func main() {
    go fmt.Println("new routine")
    fmt.Println("main routine")
}
```

- Expected: prints both messages.
- Often: only prints `main routine` because `main()` exits too quickly.

### Hacky Fix (Not Recommended)

```go
import (
    "fmt"
    "time"
)

func main() {
    go fmt.Print("new routine")
    time.Sleep(100 * time.Millisecond)  // sleep in main Goroutine
    fmt.Print("main routine")
}
```

- `Sleep()` lets scheduler run other Goroutines.
- **Works by chance**, not a reliable method.

### Why This is Bad

- Relies on **timing assumptions**:
  - Scheduler may behave differently across systems or Go versions.
  - OS may pause the thread, leading to unpredictable behavior.
- Leads to **intermittent bugs** that are hard to reproduce/debug.

### Best Practice

- Use **synchronization primitives** (e.g., `sync.WaitGroup`) to coordinate Goroutines properly.

> Note: Formal synchronization constructs will be covered in the next topic.

---

## Summary

- Goroutines are easy to create with `go` keyword.
- Main Goroutine exiting early kills all others.
- Avoid relying on `time.Sleep()` for coordination.
- Use synchronization constructs for safe concurrent programming.

# Module 3: Threads in Go

## Topic 2.1: Basic Synchronization

- **Synchronization** means multiple threads agree on the timing of global events.
- Goroutines are **largely independent** and unaware of each other’s state or progress.
- Synchronization introduces **global events** visible to all threads, enabling **conditional execution**.

### Motivation

- Without synchronization, Goroutines may interleave unpredictably.
- To **enforce a specific order**, such as ensuring updates occur before dependent operations, synchronization is required.

### Race Condition Example

**Thread 1:**

```
x = 1
x = x + 1
```

**Thread 2:**

```
print(x)
```

Two interleavings:

- `print(x)` after line 1: outputs `1`
- `print(x)` after line 2: outputs `2`

### Fix via Global Event

- Introduce a **global event** after the update.
- **Thread 1:**

```
x = 1
x = x + 1
global_event()
```

- **Thread 2:**

```
if global_event_occurred:
    print(x)
```

This ensures `print(x)` happens **only after** `x` is fully updated.

### Key Insights

- **Concurrency**: Maximizes hardware by allowing any interleaving.
- **Synchronization**: Restricts execution to ensure correctness.
- Every use of synchronization **reduces performance** by constraining scheduler flexibility.

> Synchronization is "bad but necessary" — it limits concurrency but ensures correctness where ordering matters.

---

## Topic 2.2: WaitGroups

- `sync.WaitGroup` is Go’s **primitive for Goroutine synchronization**.
- Useful for blocking a Goroutine until other Goroutines finish.

### Methods

| Method   | Description                        |
| -------- | ---------------------------------- |
| `Add(n)` | Increments internal counter by `n` |
| `Done()` | Decrements the counter             |
| `Wait()` | Blocks until counter reaches zero  |

### Use Case

Fix the problem where `main()` exits before spawned Goroutines complete.

### Example

```go
import (
    "fmt"
    "sync"
)

func foo(wg *sync.WaitGroup) {
    defer wg.Done()  // Ensure Done is called at end
    fmt.Println("new routine")
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)        // Expect 1 Goroutine
    go foo(&wg)      // Start Goroutine
    wg.Wait()        // Block until Done is called
    fmt.Println("main routine")
}
```

### Flow

1. Main initializes WaitGroup `wg`
2. Calls `wg.Add(1)` before launching Goroutine
3. Goroutine does its work, then calls `wg.Done()`
4. Main blocks at `wg.Wait()` until counter = 0

### Generalization

- Can wait for **any number of Goroutines** by adjusting the initial `Add(n)`
- Each Goroutine must invoke `Done()`
- Use `defer` to ensure Done is executed even on panic or early return

### Summary

- **WaitGroups** solve premature exit problem cleanly
- Don't rely on `sleep()` or timing assumptions
- Always coordinate Add/Done/Wait correctly for predictable synchronization

# Module 3: Threads in Go

## Topic 3.1: Communication

- Goroutines often collaborate to solve **larger tasks**, not operate independently.
- Communication is needed to **exchange data** between Goroutines.

### Example Use Case: Web Server

- Web servers handle many simultaneous clients.
- For each connection, a new Goroutine/thread handles the client.
- They all **share server data**, e.g., serving UCI's website content.

### Example: Product of Four Integers

- Use 3 Goroutines: main + 2 workers
- Each worker multiplies a pair of integers.
- Workers send result to main, which multiplies both results.

### Data Communication

- Use **channels** to send/receive data between Goroutines.
- Channels are **typed** (e.g., `chan int`, `chan string`).
- Use `make` to create: `c := make(chan int)`

### Sending and Receiving

- **Send:** `c <- 3` sends 3 into channel `c`
- **Receive:** `x := <-c` reads from channel `c`

### Key Point

- You can pass arguments **into** Goroutines when launching (not via channel)
- Use channels for **any ongoing communication** after launch

### Complete Code Example

```go
func prod(v1, v2 int, c chan int) {
    c <- v1 * v2
}

func main() {
    c := make(chan int)
    go prod(1, 2, c)
    go prod(3, 4, c)
    a := <-c
    b := <-c
    fmt.Println(a * b)
}
```

---

## Topic 3.2: Blocking on Channels

### Default Channel: Unbuffered

- Unbuffered channels **cannot hold data in transit**
- Both sender and receiver **must be ready at the same time**

### Blocking Behavior

- **Send blocks** until receive occurs
- **Receive blocks** until send occurs

```go
// Sending (blocks until someone receives)
c <- 3

// Receiving (blocks until someone sends)
x := <-c
```

### Synchronization via Channels

- Channels provide **both data transfer and synchronization**
- You can use them just to **synchronize** (throw away the value)

```go
c <- 3      // Send
<-c         // Receive and ignore
```

- This effectively synchronizes without using the value

### Key Idea

- Channels act like **WaitGroups** when used for synchronization
- They block one Goroutine until another reaches the matching send/receive

---

## Topic 3.3: Buffered Channels

### Concept

- Buffered channels can **store a finite number of values** in transit
- Created using `make(chan Type, capacity)`

```go
c := make(chan int, 3)  // Buffer size = 3
```

### Blocking Behavior

- **Send blocks** only if buffer is **full**
- **Receive blocks** only if buffer is **empty**

### Example

```go
c := make(chan int, 1)
c <- 3         // OK, buffer has room
x := <-c       // OK, reads from buffer
```

### Producer-Consumer Scenario

- **Producer** sends data (e.g., from microphone, sensor)
- **Consumer** processes data

### Buffer Benefits

- Helps when **producer/consumer speeds vary**
- Reduces unnecessary blocking
- Temporary mismatch is tolerated if buffer isn’t full/empty

### Limits

- Buffer is finite: If producer is too fast or consumer too slow, blocking still happens eventually
- Buffer helps **smooth bursty communication**, but does not eliminate coordination needs

### Summary

| Channel Type   | Send Blocks When  | Receive Blocks When |
| -------------- | ----------------- | ------------------- |
| Unbuffered     | No receiver ready | No sender ready     |
| Buffered (n=3) | Buffer is full    | Buffer is empty     |

### Final Note

- Choose buffer size based on:
  - Communication burstiness
  - Timing mismatches
  - Memory availability
