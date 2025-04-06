## ✅ Questions

**1. If two tasks are executing in parallel, which of the following statements is true?**

- ○ They are using exactly the same hardware at the same time.
- ○ They are using different hardware, but running at the same time.
- ○ Their executions are alternating in time.
- ○ One task executes immediately after the other finishes.

**2. What does the von Neumann bottleneck state about computer architectures?**

- ○ Power consumption is a limiting factor for performance.
- ○ Temperature is a performance bottleneck.
- ○ Clock frequency cannot be improved without considering temperature.
- ○ Memory access time is a performance bottleneck.

**3. What does Moore's law directly observe?**

- ○ Power consumption doubles every 18 months.
- ○ Transistor density doubles every 2 years.
- ○ Processor power doubles every 2 years.
- ○ Transistor switching delay is cut in half every year.

**4. How is dynamic power consumption related to voltage swing?**

- ○ Dynamic power is proportional to the square of the voltage swing.
- ○ Dynamic power is proportional to the cube of the voltage swing.
- ○ Dynamic power is proportional to the square root of the voltage swing.
- ○ Dynamic power is proportional to the capacitance.

**5. Why can't Dennard Scaling continue forever?**  
I. The speed of light limits the potential performance improvements.  
II. Voltage must remain above threshold voltage.  
III. Some noise margin must be maintained.

- ○ I only.
- ○ I and II, NOT III.
- ○ II and III, NOT I.
- ○ I, II, and III.

**6. What factor limits clock rates in future designs?**  
I. The speed of light.  
II. Excessive power consumption.  
III. Excessive temperature.

- ○ I only.
- ○ I and II, NOT III.
- ○ II and III, NOT I.
- ○ I, II, and III.

**7. One benefit of concurrent execution on a single processor is that it can hide latency. What does this mean?**

- ○ When tasks execute in parallel, only the delay of the slowest task matters.
- ○ One task can execute while another task is waiting on something.
- ○ The concurrent execution time of two tasks is less than the sum of their sequential execution times.
- ○ Total latency is reduced because two tasks can execute at the same time.

---

## ✅ Answers with Explanations

### Q1. If two tasks are executing in parallel, which of the following statements is true?

**Answer:** They are using different hardware, but running at the same time.  
**Explanation:** Parallelism implies true simultaneous execution on separate cores or processors.

---

### Q2. What does the von Neumann bottleneck state about computer architectures?

**Answer:** Memory access time is a performance bottleneck.  
**Explanation:** The limited bandwidth between the CPU and memory restricts overall system performance.

---

### Q3. What does Moore's law directly observe?

**Answer:** Transistor density doubles every 2 years.  
**Explanation:** Moore's law refers to the observation that the number of transistors on a chip doubles roughly every two years.

---

### Q4. How is dynamic power consumption related to voltage swing?

**Answer:** Dynamic power is proportional to the square of the voltage swing.  
**Explanation:** The power used in switching transistors depends on capacitance, frequency, and the square of the voltage (P ∝ V²).

---

### Q5. Why can't Dennard Scaling continue forever?

**Answer:** II and III, NOT I.  
**Explanation:** As scaling continues, voltage cannot be reduced infinitely and noise margins must be preserved, while speed of light is not a limiting factor in typical chip designs.

---

### Q6. What factor limits clock rates in future designs?

**Answer:** I, II, and III.  
**Explanation:** Physical and thermal limitations, including speed of light and power dissipation, constrain achievable clock speeds.

---

### Q7. One benefit of concurrent execution on a single processor is that it can hide latency. What does this mean?

**Answer:** One task can execute while another task is waiting on something.  
**Explanation:** Concurrent execution allows the processor to use idle time (e.g., waiting for I/O) to perform other computations, reducing perceived latency.
