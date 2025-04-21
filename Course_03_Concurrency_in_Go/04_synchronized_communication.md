# Module 4: Synchronized Communication

## Topic 1.1: Blocking on Channels

### Iterating Through Channels

- Common pattern: **consumer reads from a channel in a loop**
- Use `for range` to process data until the channel is closed

```go
for i := range c {
    fmt.Println(i)
}
```

- Each value from channel `c` is assigned to `i`
- Loop terminates when the channel is **closed** by the sender

### Closing a Channel

- Use `close(c)` to indicate no more values will be sent
- Required only when using `range` on channels
- Signals the consumer to exit the loop cleanly

### Reading from Multiple Channels

#### Sequential Reads (Need All Data)

```go
a := <-c1
b := <-c2
fmt.Println(a * b)
```

- Useful when task requires inputs from multiple sources
- **Both reads are blocking**

#### Selective Read (Need Any One)

- Use `select` to wait for whichever channel is ready first

```go
select {
case a := <-c1:
    // handle data from c1
case b := <-c2:
    // handle data from c2
}
```

- Avoids blocking on a channel that may not receive data

---

## Topic 1.2: Select

### Select: Blocking on Send or Receive

- `select` allows you to **wait on multiple channel operations**
- Can include **both sends and receives**

```go
select {
case a := <-inchan:
    fmt.Println("received", a)
case outchan <- b:
    // send b to outchan
}
```

- Whichever operation is ready first **executes**, others are skipped

### Abort Pattern with Select

- Useful for **cancelling** long-running tasks

```go
for {
    select {
    case a := <-c:
        fmt.Println(a)
    case <-abort:
        return
    }
}
```

- Keeps processing from `c` until something arrives on `abort`
- No need to care about abort data, just the signal

### Select with Default

- A `default` case allows **non-blocking** select behavior

```go
select {
case x := <-c1:
    fmt.Println(x)
case y := <-c2:
    fmt.Println(y)
default:
    fmt.Println("no communication ready")
}
```

- If no channels are ready, `default` is executed immediately
- Prevents the Goroutine from **blocking**

### Summary

- `select` supports flexible communication:
  - Choose among multiple sends/receives
  - Handle cancellation with `abort` channels
  - Use `default` to avoid blocking
- Ideal for **dynamic coordination** between concurrent Goroutines





# Module 4: Threads in Go

## Topic 2.1: Mutual Exclusion

### The Problem

- Sharing variables between Goroutines can cause **race conditions**.
- Two Goroutines modifying the same variable concurrently can lead to **inconsistent state**.

### Example: Race Condition

```go
i := 0
var wg sync.WaitGroup
wg.Add(2)
go inc(&wg)
go inc(&wg)
wg.Wait()
fmt.Println(i)
```

```go
func inc(wg *sync.WaitGroup) {
    i = i + 1
    wg.Done()
}
```

- Expected output: `2`
- Actual output: Sometimes `1`, due to interleaving at **machine instruction level**:
  1. Read i
  2. Increment
  3. Write i
- If both Goroutines read `i = 0` before either writes back, both write `1`

### Key Takeaway

- Interleaving occurs at **machine code level**, not Go source level.
- Concurrency bugs are hard to see but critical to avoid.

---

## Topic 2.2: Mutex

### Goal

- Prevent simultaneous access to shared variables.
- Ensure **only one Goroutine** modifies shared state at a time.

### Concept

- Use **mutual exclusion (mutex)** to protect critical sections.
- Declare **code blocks** that cannot be interleaved with others.

### Analogy

- Think of a **mailbox flag**:
  - Up = someone is using shared data
  - Down = data is free to use
- Goroutines must **follow the protocol**: Check flag, raise it, use data, lower it.

### Solution: Go’s `sync.Mutex`

- Go provides `sync.Mutex` for this exact purpose

---

## Topic 2.3: Mutex Methods

### Methods

| Method     | Description                                   |
| ---------- | --------------------------------------------- |
| `Lock()`   | Waits until mutex is available, then locks it |
| `Unlock()` | Releases the lock                             |

### Usage Pattern

```go
mut.Lock()
i = i + 1
mut.Unlock()
```

- Lock before using shared data
- Unlock after done to allow others in

### What Happens?

- If a Goroutine holds the lock, others calling `Lock()` **block**
- When `Unlock()` is called, **next waiting Goroutine** gets in

### Example Fixing Race Condition

```go
var i int
var mut sync.Mutex
var wg sync.WaitGroup

func inc() {
    mut.Lock()
    i = i + 1
    mut.Unlock()
    wg.Done()
}

func main() {
    wg.Add(2)
    go inc()
    go inc()
    wg.Wait()
    fmt.Println(i) // Always prints 2
}
```

### Summary

- Mutex ensures **one Goroutine at a time** accesses critical data
- Always pair `Lock()` and `Unlock()`
- Avoid forgetting `Unlock()` to prevent deadlocks
- A reliable way to make functions **concurrency-safe**






# Module 4: Threads in Go

## Topic 3.1: Once Synchronization

### Problem
- In multi-goroutine programs, initializing data exactly once is difficult due to unpredictable execution order.

### Solution: `sync.Once`
- Ensures that a given function is executed only once across multiple Goroutines.
- Blocks all other `Once.Do()` calls until the function finishes.

### Example
```go
var once sync.Once

func setup() {
    fmt.Println("Init")
}

func doStuff(wg *sync.WaitGroup) {
    once.Do(setup)
    fmt.Println("Hello")
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go doStuff(&wg)
    go doStuff(&wg)
    wg.Wait()
}
```
- Output: Init \n Hello \n Hello
- `setup()` runs once before either goroutine proceeds.

---

## Topic 3.2: Deadlock

### Definition
- A **deadlock** occurs when multiple goroutines are waiting on each other to release resources, resulting in a standstill.

### Causes
- Circular dependencies due to:
  - Channels
  - Mutexes

### Example: Channel Deadlock
```go
func doStuff(c1, c2 chan int) {
    <-c1
    c2 <- 1
}

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    go doStuff(ch1, ch2)
    go doStuff(ch2, ch1)
    select {} // Blocks forever
}
```
- Each goroutine is waiting on the other to send first → **deadlock**

### Golang Runtime
- Detects full program deadlock (`all goroutines are asleep`)
- Cannot detect **partial** goroutine deadlocks

### Prevention
- Avoid circular wait conditions
- Establish clear acquisition order of resources

---

## Topic 3.3: Dining Philosophers Problem

### Problem Overview
- 5 philosophers sitting at a table
- 5 chopsticks placed between them
- Each needs **two chopsticks** to eat
- If all pick up their left chopstick simultaneously, no one can eat → **deadlock**

### Model
- Chopstick = `sync.Mutex`
- Philosopher = `goroutine`

### Naive Implementation
```go
type Chopstick struct{ sync.Mutex }

type Philosopher struct {
    leftCS  *Chopstick
    rightCS *Chopstick
}

func (p *Philosopher) eat() {
    for {
        p.leftCS.Lock()
        p.rightCS.Lock()
        fmt.Println("eating")
        p.rightCS.Unlock()
        p.leftCS.Unlock()
    }
}
```

### Main Setup
```go
func main() {
    chopsticks := make([]*Chopstick, 5)
    for i := range chopsticks {
        chopsticks[i] = new(Chopstick)
    }

    philosophers := make([]*Philosopher, 5)
    for i := range philosophers {
        philosophers[i] = &Philosopher{
            leftCS:  chopsticks[i],
            rightCS: chopsticks[(i+1)%5],
        }
    }

    for i := range philosophers {
        go philosophers[i].eat()
    }

    select{} // Prevent main from exiting
}
```

### Deadlock Scenario
- All philosophers lock their **left** chopstick → all 5 chopsticks locked
- Next, all try to lock **right** chopstick → all blocked

### Fix (Dijkstra’s Solution)
- Pick **lower-numbered chopstick first**:
```go
if philosopherID == 4 {
    rightCS.Lock()
    leftCS.Lock()
} else {
    leftCS.Lock()
    rightCS.Lock()
}
```

### Trade-off
- Prevents deadlock
- May cause **starvation** (e.g., philosopher 4 always waits)

### Conclusion
- Deadlock arises from circular resource wait
- Fix via ordering, timeout, or breaking circular dependency
- Be cautious: concurrency correctness depends on programmer discipline

Thank you.

