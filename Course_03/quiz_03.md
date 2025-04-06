## ✅ Questions

**1. Suppose you want to start a goroutine which executes a function called `test1()`. What code would create this goroutine?**

- ○ test1() go
- ○ start test1()
- ○ goroutine test1()
- ○ go test1()

**2. When does a goroutine complete?**  
I. When its code completes.  
II. When all goroutines complete.  
III. When the main goroutine completes.

- ○ I and II, NOT III.
- ○ I and III, NOT II.
- ○ I, II, and III.
- ○ I only.

**3. Synchronization is useful for what purpose?**  
I. Restrict illegal interleavings.  
II. Force events in different goroutines to occur in sequence.  
III. Allow a goroutine to continue to execute after the main goroutine has completed.

- ○ I, II, and III.
- ○ I only.
- ○ I and III, NOT II.
- ○ I and II, NOT III.

**4. If a goroutine g1 is using a WaitGroup `wg` to wait until another goroutine g2 completes a task, what method of the WaitGroup should be called when g2 has finished the task?**

- ○ wg.Done()
- ○ wg.End()
- ○ wg.Finished()
- ○ wg.Alarm()

**5. If a goroutine g1 is using a WaitGroup `wg` to wait until another goroutine g2 completes a task, what method of the WaitGroup should be called _before_ g2 starts its task?**

- ○ wg.Fork()
- ○ wg.Start()
- ○ wg.Add()
- ○ wg.Begin()

**6. How might you write code to allow a goroutine to receive data from a channel `c`?**

- ○ x <- c
- ○ x = <- c
- ○ x = c
- ○ x <-- c

**7. What is the difference between a buffered channel and an unbuffered channel?**

- ○ A buffered channel can hold multiple objects until they are read. An unbuffered channel cannot.
- ○ A buffered channel delays the transmission of data. An unbuffered channel does not.
- ○ A buffered channel delays the reception of data. An unbuffered channel does not.
- ○ A buffered channel can communicate between more than 2 goroutines. An unbuffered channel cannot.

---

## ✅ Answers with Explanations

### Q1. What code would create this goroutine?

**Answer:** go test1()  
**Explanation:** In Go, the `go` keyword is used before a function call to start a goroutine.

---

### Q2. When does a goroutine complete?

**Answer:** I only.  
**Explanation:** A goroutine completes when its function finishes executing, independently of other goroutines.

---

### Q3. Synchronization is useful for what purpose?

**Answer:** I and II, NOT III.  
**Explanation:** Synchronization ensures correct ordering of operations and prevents data races. A goroutine cannot continue after the main goroutine exits unless explicitly coordinated.

---

### Q4. What method should be called when g2 has finished the task?

**Answer:** wg.Done()  
**Explanation:** `wg.Done()` decrements the WaitGroup counter, signaling that a goroutine has completed.

---

### Q5. What method should be called before g2 starts its task?

**Answer:** wg.Add()  
**Explanation:** `wg.Add(1)` increments the WaitGroup counter before launching a new goroutine.

---

### Q6. How to receive data from a channel `c`?

**Answer:** x = <- c  
**Explanation:** The `<-` operator is used for channel communication. `x = <-c` assigns the received value to `x`.

---

### Q7. What is the difference between a buffered and an unbuffered channel?

**Answer:** A buffered channel can hold multiple objects until they are read. An unbuffered channel cannot.  
**Explanation:** Buffered channels allow sending without blocking until the buffer is full. Unbuffered channels require both sender and receiver to be ready.
