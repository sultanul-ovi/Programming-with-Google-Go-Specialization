## ✅ Questions

**1. What line of code could be used to define a loop which iteratively reads from a channel `ch1`?**

- ⭕ `for i <- ch1`
- ⭕ `for i := range ch1`
- ⭕ `for i, err <- range ch1`
- ⭕ `for i := ch1`

**2. What does the `select` keyword do?**

- ⭕ Executes a set of case statements.
- ⭕ Allows a choice of channels to wait on.
- ⭕ Chooses the greatest of a set of numbers.
- ⭕ Chooses an element from a list based on a user-defined criterion.

**3. What is the meaning of the `default` clause inside a `select`?**

- ⭕ The default clause is executed if all case clauses are blocked.
- ⭕ The default clause is executed before any case clause is executed.
- ⭕ The default clause is executed after any case clause is executed.
- ⭕ The default clause is executed only if a case clause is executed.

**4. Suppose that there are two goroutines, g1 and g2, which share a variable `x`. X is initialized to 0. The only instruction executed by g1 is `x = 4`. The only instruction executed by g2 is `x = x + 1`. What is a possible value for x after both goroutines are complete?**

- ⭕ I. 0
- ⭕ II. 1
- ⭕ III. 4
- ⭕ IV. 5
- ⭕ I and II only.
- ⭕ II and III only.
- ⭕ I, II, and III but not IV.
- ⭕ II, III, IV but not I.

**5. What is mutual exclusion?**

- ⭕ When a single goroutine can execute only one of two blocks of code.
- ⭕ When a single goroutine cannot execute a block of code.
- ⭕ When multiple goroutines cannot execute blocks of code concurrently.
- ⭕ When a single goroutine is the only goroutine which ever accesses a variable.

**6. What is true about deadlock?**  
I. It can always be detected by the Golang runtime  
II. It's caused by a circular dependency chain between goroutines  
III. It can be caused by waiting on channels

- ⭕ I and II only.
- ⭕ II and III only.
- ⭕ I and III only.
- ⭕ I, II, and III.

**7. What is the method of the `sync.Mutex` type which must be called at the beginning of a critical region?**

- ⭕ Lock()
- ⭕ Unlock()
- ⭕ Take()
- ⭕ Block()

---

## ✅ Answers with Explanations

### Q1. What line of code could be used to define a loop which iteratively reads from a channel `ch1`?

**Answer:** `for i := range ch1`  
**Explanation:** This idiomatic Go loop reads values from a channel until it is closed.

---

### Q2. What does the `select` keyword do?

**Answer:** Allows a choice of channels to wait on.  
**Explanation:** The `select` statement lets a goroutine wait on multiple channel operations.

---

### Q3. What is the meaning of the `default` clause inside a `select`?

**Answer:** The default clause is executed if all case clauses are blocked.  
**Explanation:** It provides a non-blocking alternative when no communication is ready.

---

### Q4. What is a possible value for `x` after both goroutines are complete?

**Answer:** I, II, and III but not IV.  
**Explanation:** Due to race conditions, `x` could be 0 (g2 finishes first), 1 (g2 runs before g1 sets), or 4 (g1 overwrites). But 5 isn't possible without specific ordering.

---

### Q5. What is mutual exclusion?

**Answer:** When multiple goroutines cannot execute blocks of code concurrently.  
**Explanation:** Mutual exclusion ensures only one goroutine enters a critical section at a time.

---

### Q6. What is true about deadlock?

**Answer:** I, II, and III.  
**Explanation:** Deadlocks can occur from circular waits, channel dependencies, and are sometimes detectable at runtime.

---

### Q7. What is the method of the `sync.Mutex` type which must be called at the beginning of a critical region?

**Answer:** Lock()  
**Explanation:** `Lock()` is used to enter a critical section guarded by a mutex, ensuring exclusive access.
