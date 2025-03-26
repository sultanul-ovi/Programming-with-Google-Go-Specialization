# Using Interfaces

## Ways to Use an Interface

- need a function which takes multiple taypes of parameter
- function `foo()` parameter
  - type X or type Y
- define interface Z
- `foo()` parameter is interface Z
- type X and Y satisfy Z
- interface methods must be those needed by `foo()`

## FitInYard()

- many possible shape types
  - rectangle, triangle, circle, etc.
- `FitInyard()` should take many shape types
- valid shape types must have:
  - `Area()`
  - `Perimeter()`
- any shape with these methods is OK

## Interface for Shapes

```golang
type Shape2d interface{
  Area() float64
  Perimeter() float64
}

type Triangle{
  //...
}
func (t Triangle) Area() float64{}
func (t Triangle) Perimeter() float64{}

type Rectangle{
  //...
}
func (t Rectangle) Area() float64{}
func (t Rectangle) Perimeter() float64{}
```

- rectangle and triangle satisfy Shape2D interface

## FitInYard() Implementation

```golang
func FitInYard(s Shape2D) bool{
  if (s.Area() > 100 && s.Perimeter() > 100){
    return true
  }
   return false
}
```

- parameter is any type that satisfies the interface

## Empty Interface

- empty interface specifies no methods
- all types satisfy the empty interface
- use it to have a function accept any type as a parameter

```golang
func PrintMe(val interface{}){
  fmt.Println(val)
}
```
