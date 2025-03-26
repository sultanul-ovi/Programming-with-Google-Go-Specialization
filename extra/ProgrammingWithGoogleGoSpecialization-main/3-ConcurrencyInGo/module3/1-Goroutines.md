# Goroutines

## Creating a Goroutine

- one goroutine is created automatically to execute the `main()`
- other goroutines are created using the `go` keyword

## Exiting a Goroutine

- a goroutine exits when its code is complete
- when the main goroutine is complete, all other goroutines exit
- a goroutine may not complete its execution because main completes early

### Early Exit

```golang
func main(){
  go fmt.Printf("New routine")
  fmt.Printf("Main routine") // only main routine is printed
}
```

- main finished before the new goroutine started

### Delayed Exit

```golang
func main(){
  go fmt.Printf("New routine")
  time.Sleep(100 * time.Millisecond)
  fmt.Printf("Main routine") // only main routine is printed without delay
}
```

- add a delay in the main routine to give the new routine a chance to complete
- "New RoutineMain Routine" is now printed

### Timing with Goroutines

- adding a delay to wait for a goroutine is BAD!
- timing assumptions may be wrong
  - assumption: delay of 100 ms will ensure that goroutine has time to execute
  - maybe the OS schedules another thread
- timing is nondeterministic
- need formal synchronization constructs
