# Module 2: Concurrency Basics

## Topic 1.1: Processes

- A **process** is an _instance_ of a running program.
- Each process has its **own memory**:
  - **Code** (instructions to execute)
  - **Stack** (function calls and local variables)
  - **Heap** (dynamically allocated memory)
  - **Virtual address space** (e.g., 0 to 2GB per process)
- **Shared libraries** may be shared between processes.

### Registers (process-specific):

- **Program counter**: holds the address of the next instruction.
- **Data registers**: store temporary data.
- **Stack pointer**: tracks the top of the stack.
- The **context** of a process = its memory + registers.

### Operating System's Role:

- Allows **concurrent execution** of multiple processes.
- Ensures processes do **not interfere** with each other.
  - Different processes may think they have the same memory (e.g., address 1000), but OS maps them to different physical memory.
- Provides **fair access** to CPU and other resources.
  - This fairness is achieved by **scheduling**.

### Scheduling:

- OS gives CPU to each process for a time slice (e.g., 20 ms in Linux).
- Then switches to another process.
- Rapid context switching creates **illusion of parallelism**.

### Viewing Processes:

- Windows: Use **Task Manager**.
- Linux/macOS: Use terminal command `ps` or `ps -l`.

## Topic 1.2: Scheduling

- **Scheduling** = deciding which process/thread runs when.
- Basic scheduling strategy: **Round Robin**
  - Each process gets equal CPU time in cycles.

### Prioritization:

- Some processes are **more important** than others.
  - E.g., Anti-lock braking > Music playback
- **Priority-based scheduling**: High-priority tasks get more CPU time.

### Context Switching:

- Switching between processes or threads.
- Involves saving current process **state** (context) to memory and loading next process's state from memory.
- State includes: registers, stack pointer, program counter, virtual memory info.
- The switch is executed by the **kernel**.
- A **timer interrupt** typically triggers the switch (e.g., every 10–20 ms).

## Topic 1.3: Threads and Goroutines

### Threads:

- **Thread** = lightweight version of a process.
- Threads share:
  - Virtual memory
  - File descriptors
- Threads have their own:
  - Stack
  - Registers
  - Program counter

### Benefits:

- **Faster context switches** than between processes.
- Operating systems now **schedule threads** instead of processes.

### Goroutines (Go Language):

- **Goroutine** = Go's version of a thread.
- Multiple goroutines run **inside a single OS thread**.
- Managed by the **Go Runtime Scheduler**.

### Hierarchy:

- Process
  - OS Thread
    - Goroutines

### Logical Processors:

- Go uses **logical processors**, each mapped to an OS thread.
- Default: 1 logical processor → only **concurrent** execution.
- Programmer can increase logical processors (e.g., to match CPU cores) for **parallel execution**.

### Summary:

- OS does **thread-level** scheduling.
- Go's runtime does **goroutine-level** scheduling.
- Parallelism possible if multiple logical processors are used.

---

# Module 2: Concurrency Basics

## Topic 2.1: Interleaving

- **Concurrent code is hard** because it is difficult to track program state at any moment.
- In **sequential code**, execution is predictable (e.g., line 1 → 2 → 3).
- In **concurrent code**, multiple tasks (threads) may execute **in different interleavings** every time.

### Key Concept: Interleaving

- Tasks have a fixed instruction order **within themselves**, but interleaving between tasks is **non-deterministic**.
- Example: Task 1 and Task 2 each have 3 instructions (1, 2, 3). Possible interleavings:
  - Task1.1 → Task2.1 → Task1.2 → Task2.2 → ...
  - Task1.1 → Task1.2 → Task1.3 → Task2.1 → ...
- The interleaving defines **overall program behavior** and makes reasoning complex.

### Additional Complexity

- Interleaving happens at the **machine instruction level**, not source code level.
  - A single Go/C statement (e.g., `a = b + c`) may compile to multiple machine instructions.
  - Interleaving can occur between these instructions.
- Debugging becomes harder: a crash at the same source line may happen in different machine states each time.

## Topic 2.2: Race Conditions

### What is a Race Condition?

- **Definition**: A bug where the program output **depends on the interleaving**.
- Interleaving is **non-deterministic**, so the output becomes unpredictable.
- A race condition leads to **non-deterministic behavior**, which is generally **undesirable and incorrect**.

### Example:

Task 1:

```
x = 1
x = x + 1
```

Task 2:

```
print(x)
```

- Interleaving A: print runs after first assignment → prints `1`
- Interleaving B: print runs after second assignment → prints `2`
- Output depends on interleaving → **race condition**

### Cause: Communication

- Tasks **communicate via shared variables** (e.g., `x`).
- **If no communication**, there are **no race conditions**.
- Communication is necessary but must be **safely managed**.

### Threads Are Not Completely Independent

- **Threads share memory** (virtual space, files, etc.)
- Communication is **common** in real-world applications

#### Example: Web Server

- One thread per client connection
- Threads access shared data (e.g., web page content)
- Clients may write (e.g., post) and read shared pages → communication
- May lead to race conditions if not synchronized

#### Example: Image Processing

- Large image → split into regions → each processed by a thread
- Threads need **neighbor pixel values** for filters (e.g., blur)
- Requires communication at region boundaries
- **Common in GPU programming** with many cores and threads

### Summary

- Communication between threads is **inevitable** and must be **carefully synchronized**.
- Poor handling leads to **race conditions** and **unreliable behavior**.
- Even "parallel-friendly" tasks like image processing need **controlled communication** to avoid errors.

---
