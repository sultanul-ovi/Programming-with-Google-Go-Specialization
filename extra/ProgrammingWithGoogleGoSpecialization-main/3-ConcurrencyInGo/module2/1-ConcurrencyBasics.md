# Concurrency Basics

## Processes

- an instance of a running program
- things unique to a process

1. memory
    - virtual address space
    - code, stack, head, shared libraries

2. registers
    - program counter, data regs, stack ptr ...

## Operating Systems

- allows many processes to execute concurrently
- processes are switched quickly
  - 20ms

## Scheduling Processes

- operating system schedules processes for execution
- gives the illusion of parallel execution
- OS gives fair access to CPU, memory, etc.

## Context Switch

- control flow changes from one process to another
- process "context" must be swapped
- air cooling (fans) can only remove so much heat

## Threads and Goroutines

### Threads vs. Processes

- threads share some context
- many threads can exist in one process
- OS schedules threads rather than processes

### Goroutines

- like a thread in go
- many goroutines execute within a single OS thread

### Go Runtime Scheduler

- schedules goroutines inside an OS thread
- like a little OS inside a single OS thread
- logical processor is mapped to a thread
