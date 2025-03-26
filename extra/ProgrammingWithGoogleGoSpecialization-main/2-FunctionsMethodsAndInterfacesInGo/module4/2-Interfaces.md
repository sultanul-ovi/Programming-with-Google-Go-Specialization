# Interfaces

- set of method signatures
  - name, parameters, return values
  - implementation is NOT defined
- used to express conceptual similarity between types
- example: `Shape2D` interface
- all 2D shapes must have `Area()` and `Perimeter()`

## Satisfying an Interface ;)

- type satisfies an interface if type defines all methods specificed in the interface
  - same method signatures
- `Rectangle` and `Triangle` types satisfy the `Shape2D` interface
  - must have `Area()` and `Perimeter()` methods
  - additional methods are OK
- similar to inheritance with overriding

## Defining an Interface Type

```golang
type Shape2D interface{
  Area() float64
  Perimeter() float64
}
type Triangle{
  //...
}

func (t Triangle) Area() float64 {
  // ...
}
func (t Triangle) Perimeter() float64 {
  // ...
}
```

- `Triangle` type satisfies the `Shape2D` interface
- no need to state it explicitly, i.e. explicitly "connecting" `Triangle` to `Shape2D`
