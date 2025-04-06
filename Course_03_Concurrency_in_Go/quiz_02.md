## ✅ Questions

**1. What unique resources does a process have?**  
I. Program counter value.  
II. Virtual address space.  
III. Shared libraries

- ○ I only.
- ○ I and II, NOT III
- ○ II and III, NOT I
- ○ II only

**2. What does an operating system do to enable concurrency on a single processor machine?**

- ○ Provides a graphic interface to allow the user to control system functions.
- ○ Partitions processor hardware to allow parallel execution of multiple processes.
- ○ Combines different processes into a single process.
- ○ Interleaves the execution of different processes.

**3. What is the "context" that is referred to in the term "context switch"?**

- ○ Shared libraries used by a process.
- ○ Memory and register values unique to a process.
- ○ The parameters specific to the operating system.
- ○ The set of executing processes.

**4. What is the difference between a thread and a process?**

- ○ A thread has less unique context than a process.
- ○ Threads do not have unique program counter values.
- ○ Processes can execute programs but threads cannot.
- ○ A thread is another name for a process.

**5. What is the main function of the Go runtime scheduler?**

- ○ Schedules operating systems processes.
- ○ Schedules goroutines inside an operating system thread.
- ○ Schedules operating system threads inside a process.
- ○ Assigns operating system threads to hardware processors.

**6. Suppose that there are two goroutines executing, g1 and g2. Assume that g1 requires 1 second to complete when it is executed alone, and g2 requires 2 seconds to complete when it is executed alone. Assume that there is no synchronization between g1 and g2. Assume that g1 and g2 are executed concurrently and that g1 begins executing first. What do we know about the relative completion times of g1 and g2?**

- ○ g2 will complete before g1.
- ○ Nothing!
- ○ g1 will complete before g2.
- ○ g2 and g1 will complete at virtually the same time.

**7. Assume that two goroutines are executed concurrently. Which of the following statements is true?**

- ○ The relative order of the execution of their instructions is deterministic.
- ○ The relative order of the execution of their instructions is unknown, but it is the same each time they are executed together.
- ○ The relative order of the execution of their instructions can be different every time that they are executed together.
- ○ The relative order of the execution of their instructions can be determined from startup conditions.

---

## ✅ Answers with Explanations

### Q1. What unique resources does a process have?

**Answer:** I and II, NOT III  
**Explanation:** A process has its own program counter and virtual address space, but shared libraries can be used by multiple processes.

---

### Q2. What does an operating system do to enable concurrency on a single processor machine?

**Answer:** Interleaves the execution of different processes.  
**Explanation:** Concurrency on a single processor is achieved through time-slicing, where the OS switches between processes.

---

### Q3. What is the "context" that is referred to in the term "context switch"?

**Answer:** Memory and register values unique to a process.  
**Explanation:** Context refers to the CPU state, which must be saved and restored during process switches.

---

### Q4. What is the difference between a thread and a process?

**Answer:** A thread has less unique context than a process.  
**Explanation:** Threads share process resources but maintain individual registers and stack, making their context lighter.

---

### Q5. What is the main function of the Go runtime scheduler?

**Answer:** Schedules goroutines inside an operating system thread.  
**Explanation:** The Go scheduler manages goroutines, which are lightweight threads, mapping them onto OS threads.

---

### Q6. What do we know about the relative completion times of g1 and g2?

**Answer:** g1 will complete before g2.  
**Explanation:** Since g1 starts first and requires less time, it will finish before g2 assuming no synchronization.

---

### Q7. Which of the following statements is true regarding concurrent execution of goroutines?

**Answer:** The relative order of the execution of their instructions can be different every time that they are executed together.  
**Explanation:** Concurrency introduces nondeterminism, meaning goroutines may interleave differently in each run.
