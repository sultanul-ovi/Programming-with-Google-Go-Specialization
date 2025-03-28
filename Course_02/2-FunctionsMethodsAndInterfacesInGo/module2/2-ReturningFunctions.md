# Returning Functions

## Functions as Return Values

- functions can return functions
- might create a function with controllable parameters
- example: distance to origin function
  - takes a point (x, y, coordinates)
  - returns distance to origin
- what if i want to change the origin?
  - option 1: pass origin as argument
  - option 2: define function with new origin

```golang
// MakeDistOrigin returns a function which returns a float value
func MakeDistOrigin (o_x, o_y float64) func (float64, float64) float64{
  fn:=func (x, y float64) float64{
    return math.Sqrt(math.Pow(x - o_x, 2) + math.Pow(y - o_y, 2))
  }
  return fn
}
```

- origin location is passed as an argument
- origin is built into the returned function

## Special-Purpose Functions

```golang
func main(){
  Dist1:=MakeDistOrigin(0,0) // make Dist1 be a function returned by MakeDistOrigin()
  Dist2:=MakeDistOrigin(2,3)
  fmt.Println(Dist1(2,2))
  fmt.Println(Dist2(2,2))
}
```

- Dist1() and Dist2() have different origins

## Environment of a Function

- set of all names that are valid inside a function
- names defined locally in the function
- lexical scoping
- environment includes names define in block where the function is defined

```golang
var x int
func foo(y int){
  z:=1
  //...
}
```

## Closure

- function + its environment
- when functions are passed/returned, their environment comes with them!
