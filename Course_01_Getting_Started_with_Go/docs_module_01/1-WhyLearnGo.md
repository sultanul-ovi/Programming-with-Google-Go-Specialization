# Why should I learn Go?

## Advantages of Go

1. Code runs fast
2. Garbage collection
3. Simpler objects
4. Efficient concurrency

### Software Translation

- machine language: CPU instructions represented in binary
- assembly language: CPU instructions with mnemonics
  - easier to read
  - equivalent to machine language
- high-level language: commonly used languages (C/C++, Java, Python, Go, etc.)
- much easier to use

### Compiled vs Intepreted

#### Compilation

- translate instructions once before running the code
  - C/C++, Java(partially)
  - translation occurs only once, saves time
    - runs machine language instructions directly because they're already compiled into machine language by the compiler

#### Interpretation

- translate instructions while code is executed
  - python, Java (partially)
  - translation occurs every execution
    - in addition to actually executing the instruction, translation from the instruction into the equivalent machine code needs to happen
  - requires an interpreter

#### Efficiency vs Ease-of-Use

- compiled code is fast
- interpreters make coding easier
  - manage memory automatically
  - infer variable types
- Go is a good compromise

### Garbage Collection

- automatic memory management
  - where should memory be allocated?
  - when can memory be deallocated?
- manual memory management is hard
  - deallocate too early, false memory accesses
  - deallocate too late, wasted memory

#### Go includes garbage collection

- typically only done by interpreters
  - when go code is compiled, garbage collection also gets compiled automatically
  - this slows down execution but not much, and there's a lot of advantage of having this automatic garbage collection
