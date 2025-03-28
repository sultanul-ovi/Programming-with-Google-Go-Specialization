# Type Assertions

## Concealing Type Differences

- interfaces hide the differences between types

```golang
func FitInYard(s Shape2D) bool{
  if (s.Area() > 100 && s.Perimeter() > 100){
    return true
  }
   return false
}
```

- sometimes you need to treat different types in different ways

## Exposing Type Differences

- example: graphics program
- `DrawShape()` will draw any shape
  - `func DrawShape (s Shape2d){}`
- underlying API has different drawing functions for each shape
  - `func DrawRect(r Rectangle){}`
  - `func DrawTriangle(t Triangle){}`
- concrete type of shape s must be determined

## Type Assertions for Disambiguation

- type assertions can be used to determine and extract the underlying concrete type

```golang
func DrawShape(s Shape2D) bool{
  rect, ok := s.(Rectangle)
  if ok{
    DrawRect(rect)
  }
  tri, ok := s.(Triangle)
  if ok{
    DrawRect(tri)
  }
}
```

- type assertion extracts Rectangle from Shape2D
  - concrete type in parentheses
- if interface contains concrete type
  - rect == concrete type, ok == true
- if interface does not contain concrete type
  - rect == zero, ok == false


## Type Switch

- switch statement used with a type assertion

```golang
func DrawShape(s Shape2D) bool{
  switch := sh := s.(type){
    case Rectangle:
      DrawRect(sh)
    case Triangle:
      DrawRect(sh)
  }
}
```
