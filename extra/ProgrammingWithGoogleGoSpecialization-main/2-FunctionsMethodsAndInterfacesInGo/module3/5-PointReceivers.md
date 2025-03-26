# Point Receivers

## Limitations of Methods

- receiver is passed implicitly as an argument to the method
- method cannot modify the data inside the receiver
- example: `OffsetX()` should increase coordinate

```golang
func main(){
  p1 := Point(3,4)
  p1.OffsetX(5) // 3,4 are copies
}
```

### Large Receivers

- if receiver is large, lots of copying is required

```golang
type Image[100][100]int
func main(){
  i1:=GrabImage()
  i1.BlurImage()
}
```

- 10,000 ints copied to BlurImage()

## Pointer Receivers Ex

```golang
func (p * Point) OffsetX(v float64) { // now, this doesn't create a copy but directly accesses values in a Point instance
  p.x = p.x+v
}
```

- receiver can be a pointer to a type
- call by reference, pointer is passed to the method
