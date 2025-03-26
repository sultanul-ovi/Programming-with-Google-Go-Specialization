# Parallel Execution

- two programs execute in parallel if they execute at exactly the same time
- at time t, an instruction is performed for both P1 and P2

## Why Use Parallel Execution

- tasks may complete more quickly
- example: two piles of dishes to wash
  - two dishwashers can complete twice as fast as one
- some tasks must be performed sequentially
- examples: wash dish, dry dish
  - must wash before you can dry

## Speedup without Parallelism

- can we achieve speedup without parallelism?
- design faster processors
  - get speedup without changing software
- design processor with more memory
  - reduces the Von Neumann bottleneck
  - cache access time = 1 clock cycle
  - main memory access time = ~100 clock cycles
  - increasing on-chip cache improves performance

## Moore's Law

- predicted that transistor density would double every two years
- smaller transistors switch faster
- not a physical law, just an observation
- exponential increase in density would lead to exponential increase in speed

### Power/Temperature Problem

- transistors consume power when they switch
- increasing transistor density leads to increase power consumption
  - smaller transistors use less power, but density scaling is much faster
- high power leads to high temperature
- air cooling (fans) can only remove so much heat

### Dynamic Power

- P = α * CFV^2
- α is percent of time switching
- C is capacitance (related to size)
- F is the clock frequency
- V is voltage swich (from low to high)
  - voltage is important
    - 0-5V uses much more power than 0-1.3V etc

### Dennard Scaling

- voltage should scale with transistor size
- keeps power consumption, and temprature, low
- problem: voltage can't go too low
  - must stay above threshold voltage
  - noise prblems occur
- problem: doesn't consider leakage power

### Multi-Core Systems

- P = α * CFV^2
- cannot increase frequency
- can still add processor cores, without increasing frequency
  - trend is apparent today
- parallel execution is needed to exploit multi-core systems
- code made toexecute onmultiple cores
- different programs on different cores
