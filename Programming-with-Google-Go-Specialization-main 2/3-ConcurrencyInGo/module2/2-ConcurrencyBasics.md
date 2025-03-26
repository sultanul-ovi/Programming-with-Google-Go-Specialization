# Concurrency Basics

## Interleaving

- order of execution within a task is known

### Possible Interleavings

- many interleavings are possible
- must consider all possibilities
- ordering is *non-deterministic*

## Race Conditions

- outcome depends on non-deterministic ordering
- races occur due to *communication*

## Communication Between Tasks

- threads are largely independent but not completely independent
- web server, one thread per client
  - it makes sense to make a web application multithreaded
    - web page data -> web server -> {client, client, client}
- image processing, 1 thread per pixel block
  - thread 1 -> [pixel][pixel] -> thread 2
