# First-Class Values

## Functions are First-Class

- functions can be treated like other types
  - variables can be declared with a function type
  - can be created dynamically
  - can be passed as arguments and returned as values
  - can be stored in data structures

## Variables as Functions

- declare a variable as a func

```golang
var funcVar func(int) int
func incFn(x int) int {
  return x + 1
}

func main() {
  funcVar = incFn // can also just be funcVar := incFn
  fmt.Println(funcVar(1))
}
```

- function is on right-hand side, without ()

## Functions as Arguments

- function can be passed to another function as an argument

```golang
func applyIt(afunct func(int) int, val int){
  return afunct(val)
}
func incFn(x int) int {return x+1}
func decFn(x int) int {return x-1}
func main(){
  fmt.Println(applyIt(incFn,2))
  fmt.Println(applyIt(decFn,2))
}
```

## Anonymous Functions

- don't need to name a function

```golang
func applyIt(afunct func(int) int, val int){
  return afunct(val)
}
func main(){
  v:=applyIt(func(x int)int{return x+1},2)
  fmt.Prinln(v)
}
```
