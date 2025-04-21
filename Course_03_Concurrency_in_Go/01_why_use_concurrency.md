# Module 1: Why Use Concurrency?

## Topic 1.1: Parallel Execution

- **Go** has **concurrency built into the language**.
  - Unlike languages like C or Python, Go's concurrency tools are native.

### Parallel Execution:

- **Definition**: Two tasks execute _at exactly the same time_.
- Requires **multiple hardware units**:
  - Multiple processors or **multi-core** CPUs.
  - One instruction per core per time unit.

### Benefits:

- **Improved throughput**: Multiple tasks complete faster _together_.
- **Example**: Two people washing/drying dishes simultaneously increase overall speed.

### Limitations:

- **Sequential dependencies** prevent full parallelism.
  - E.g., a dish must be washed before it can be dried.
- **Hardware constraints**:
  - Need multiple sinks/racks to enable true parallelism.
- **Not all tasks are parallelizable** — due to nature of the computation.

## Topic 1.2: Von Neumann Bottleneck

### Traditional Speedup:

- **Faster processors** = faster programs (historically).
- Speedups were mainly from:
  - Increased **clock speed**
  - **Larger caches**
  - Improved **memory access** times

### The Bottleneck:

- CPUs depend on **memory** for instructions/data.
- **Memory access is slower** than CPU speeds → delays execution.
- Even fast CPUs **wait on memory**, slowing actual performance.

### Cache Memory:

- Built on-chip to mitigate memory delays.
- Provides fast-access buffer between CPU and main memory.

### Moore's Law:

- **Transistor density doubles every ~2 years**.
- Denser = smaller transistors → faster switching → faster processors
- Was the dominant driver of speedup — **now tapering off**.

### Consequence:

- Hardware improvements **alone are no longer enough**.
- Programmers must adapt by writing **concurrent code**.

## Topic 1.3: Power Wall

### Limitation on Speedup:

- As transistor count increases, so does **power consumption**.
- Leads to the **Power Wall**:
  - More power → more **heat** → need for **cooling**
  - **Thermal limits** prevent increasing frequency further

### Power Equation:

- **P = α · C · F · V^2**
  - α: activity factor (how often switching)
  - C: capacitance (related to transistor size)
  - F: frequency
  - V: voltage swing

### Scaling Issues:

- **Dennard Scaling**: voltage should scale with transistor size
- Problem: voltage can't drop below physical limits (threshold + noise)
- **Leakage Power**: Even idle transistors consume power as they shrink

### Cooling:

- Standard: **air cooling** (fans, heat sinks)
- Extreme: **liquid cooling** in supercomputers
- Too much heat = **chip damage or failure**

### Solution: Multicore

- Instead of increasing frequency → add **more cores**
- Multi-core processors = multiple units of execution

### Key Point:

- To exploit multicore processors:
  - Must write **parallel/concurrent programs**
  - Otherwise, only 1 core is used → no benefit

### Programmer's Role:

- **Compiler automation** (parallelizing sequential code) is limited
- Programmers must explicitly decompose tasks into concurrent pieces
- **Concurrency is now essential** for performance improvements

---

# Module 1: Why Use Concurrency?

## Topic 2.1: Concurrent vs Parallel

### Definitions

- **Parallel execution**: Two or more tasks execute at the _same exact time_ on different hardware (cores).
- **Concurrent execution**: Tasks have **overlapping lifetimes**, even if not executing simultaneously.
  - Tasks can _alternate_ execution on the same core.
  - Not necessarily parallel.

### Diagram Explanation

- **Concurrent but not Parallel**:
  - Task 1 starts, pauses → Task 2 runs → Task 1 resumes.
  - Total execution time is longer.
- **Parallel**:
  - Tasks 1 and 2 run at the same time on separate cores.
  - Shorter completion time overall.

### Why Use Concurrency?

- Even without parallel hardware, concurrency can:
  - **Hide I/O latency**
  - **Utilize CPU while waiting for slow events** (e.g., memory, network, disk access)

### Real-World Delays

- CPU communicates with:
  - **Memory** (slow)
  - **Network card / WiFi**
  - **Display adapter**
- Example: `x = y + z` may take 100 cycles to fetch from memory, 1 cycle to compute

### Latency Hiding

- When task 1 is waiting (e.g., memory access), task 2 can run
- Overall execution gets faster, even on a **single core**

### Task-to-Hardware Mapping

- **Parallel execution**: Tasks mapped to different cores
- **Concurrent execution**: Tasks mapped to the same core, alternate execution
- **Mapping is automatic**, controlled by:
  - **Operating system**
  - **Go runtime scheduler**

### Programmer Responsibility

- Defines:
  - What **can** run in parallel
  - Communication and synchronization rules
- **Does not** define:
  - What **will** run in parallel
  - Which task runs on which core

### Why Mapping is Hard

- Depends on **hardware architecture**:
  - Number of cores
  - Local vs. shared caches
  - Shared memory architecture
  - Communication costs
- Programmers would need full awareness of system layout → impractical

### Typical Hardware Architecture

- Multi-core system:
  - 4 cores with **private caches**
  - Shared memory
  - Optimal mapping considers **data locality** and **cache access**

### Summary

- **Concurrency ≠ Parallelism**
- Concurrency is about **structure and potential** for parallelism
- Parallelism is about **actual execution** on multiple cores
- Concurrency is still **valuable on single-core systems** due to I/O latency hiding
- **Programmers focus on concurrency**, system handles the rest

---