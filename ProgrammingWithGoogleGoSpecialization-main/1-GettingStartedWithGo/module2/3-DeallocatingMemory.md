# Deallocating Space

- when a variable is no longer needed, it should be deallocated
  - memory space is made available
- otherwise, we will eventuially run out of memory

```golang
func f(){
  var x = 4
  fmt.Printf("%d", x)
}
```

- each call f() creates an integer

## Stack vs. Heap

```golang
// version 1
var x = 4 // heap
func f(){
  fmt.Printf("%d", x)
}

func g(){
  fmt.Printf("%d", x)
}

```

vs.

```golang
// version 2
func f(){
  var x = 4 // stack
  fmt.Printf("%d", x)
}

func g(){
  fmt.Printf("%d", x)
}
```

- stack is dedicated to function calls
  - local variables are stored here
  - deallocated after function completes
- heap is persistent region of memory
  - when you allocate something on the heap, it doesn't go away just because the function that allocated it is complete
  - in another language, memory deallocation must be EXPLICIT

## Manual Deallocation

- data onthe heap must be deallocated when it is done being used
- in most compiled languages (i.e. C), this is done manually

```C
x = malloc(32);
free(x);
```

- error-prone but fast
- in an interpreted language, the interpreter does the deallocation
