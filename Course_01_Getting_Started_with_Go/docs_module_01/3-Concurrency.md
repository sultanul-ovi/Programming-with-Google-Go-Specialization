# Concurrency

## Performance Limits

- Moore's Law used to help performance
  - number of transistors doubles every 18 months (outdated)
- more transistors used to lead to higher clock frequencies
- power/temperature constraints limit clock frequencies now

## Parallelism

- number of cores still increases over time
- multiple tasks may be performed at the same time on different cores
- difficulties with parallelism
  - when do tasks start/stop?
  - what if one task needs data from another task?
  - do tasks conflict in memory?

## Concurrent Programming

- Concurrency is the management of multiple tasks at the same time
- key requirement for large systems
  - essentially, not all executing sequentially or at the same time
- concurrent programming enables parallelism (if there's available hardware)
  - management of task execution
  - communication between tasks
  - synchronization between tasks

## Concurrency in Go

- Go includes concurrency primitives
- `Goroutines` represent concurrent tasks
- `Channels` are used to communicate between tasks
- `Select` enables task synchronization
- concurrency primitives are efficient and easy to use
