# Variadic and Deferred

## Variable Argument Number

- functions can take a variable number of arguments
- use `...` to indicate that the function can take any number of arguments
- treated as a slice inside function

```golang
func getMax(vals ...int)int{
  maxV := 1
  for _, v := range vals{
    if v > maxV{
      maxV = v
    }
  }
}
```

## Variadic Slice Argument

```golang
func main(){
  fmt.Println(getMax(1,3,6,4))
  vslice := []int{1,3,6,4}
  fmt.Println(getMax(vslice...))
}
```

- can pass a slice to a variadic function
- need the `...` suffix

## Deferred Function Calls

- call can be deferred until the surround function completes
- typically used for cleanup activities

```golang
func main(){
  defer fmt.Println("I am deferred") // gets printed after all other instructions are done
  fmt.Println("Hello") // gets printed first
}
```

## Deferred Call Arguments

- arguments of a deferred call are evaluated immediately

```golang
func main(){
  i := 1
  defer fmt.Println(i+1) // 2 because i+1 still gets executed just not printed
  i++ // 2
  fmt.Println("Hello")
}
```
