# Point Receivers, Referencing, and Dereferencing

- no need to dereference the pointer to the passed type

```golang
func (p *Point) OffsetX(v int){
  p.x += v // no dereference needed
}
```

- point is referenced as p, not *p inside the method
- dereferencing is automatic with the dot operator

```golang
func main(){
  p := Point{3,4}
  p.OffsetX(5)
  fmt.Println(p.x) // 8
}
```

- do not need to reference when calling the method

## Using Pointer Receivers

- good programming practice:
  - all methods for a type have pointer receivers, OR
  - all methods for a type have non-pointer receiveres
- mixing pointer/non-pointer receivers for a type will get confusing
  - pointer receiver allows modification
