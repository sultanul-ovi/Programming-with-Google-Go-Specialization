# Garbage Collection

## Pointers and Deallocation

- hard to determine when a variable is no longer in use

```golang
func foo() *int{
  x:=1
  return &x
} // foo() returns pointer to x

func main(){
  var y *int
  y = foo()
  fmt.Printf("%d", *y)
}
```

- in interpreted languages, this is done by the interpreter
  - java virtual machine
  - machine interpreter
- easy for the programmer
- slow (need an interpreter)

- Go is a compiled language which enables garbage collection
  - it's built into it
- implementation is fast
- compiler determines stack vs heap
- garbage collection in the background
