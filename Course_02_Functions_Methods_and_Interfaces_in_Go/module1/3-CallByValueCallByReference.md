# Call by Value and Call by Reference

## Call by Value

- passed arguments are copied to parameters
- modifying parameters has no effect outside the function

```golang
func foo(y int){
  y = y + 1
}
func main(){
  x := 2
  foo(x) // passes a copy of var x
  fmt.Println(x) // prints 2
}
```

## Tradeoffs of Call by Value

- advantage: data encapsulation
- function variables only changed inside the function
- disadvantage: copying time
- large objects may take a long time to copy

## Call by Reference

- programmer can pass a pointer as an argument
- called function has direct access to caller variable in memory

```golang
func foo(y *int){
  *y = *y + 1
}
func main(){
  x := 2
  foo(&x) // passes a pointer to x
  fmt.Println(x) // prints 3
}
```

## Tradeoffs of Call by Reference

- advantage: copying time
  - don't need to copy arguments
- disadvantage: data encapsulation
  - function variables may be changed in called functions
