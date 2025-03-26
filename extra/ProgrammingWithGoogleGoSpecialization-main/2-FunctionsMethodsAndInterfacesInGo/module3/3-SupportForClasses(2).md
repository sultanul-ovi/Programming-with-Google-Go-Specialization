# Support for Classes

## Struct, Again

- struct types compose data fields

```golang
type Point struct{
  y float64
  x float64
}
```

- traditional feature of classes

## Structs with Methods

- structs and methods together allow arbitrary data and functions to be composed

```golang
func (p Point) DistToOrig(){
  t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
  return math.Sqrt(t)
}
func main(){
  p1 := Point(3,4)
  fmt.Println(p1.DistToOrig())
}
```
