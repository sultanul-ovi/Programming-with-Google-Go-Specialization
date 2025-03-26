# Function Parameters and Return Values

## Function Parameters

- functions may need input data to perform their operations
- Parameters are listed in parenthesis after function name
- Arguments are supplied in the call to function

```golang
func foo(x int, y int) {
  fmt.Print(x*y)
}
func main(){
  foo(2,3)
}
```

## Parameters Options

- if no parameters are needed, put nothing in parentheses
- still need parentheses

```golang
func foo(){}
```

- list arguments of same type

```golang
func foo(x,y int){}
```

## Return values

- functions can return a value as a result
- type of return value after parameters in declaration
- function call used on right-hand side of an assignment

```golang
func foo(x int) int{
  return x+1
}
func main(){
  y := foo(1) // returns 2
}
```

## Multiple Return Values

- multiple value types must be listed in the declaration

```golang
func foo(x int) (int, int){
  return x, x+1
}
func main(){
  a, b := foo(1) // returns 1, 2
}
```
