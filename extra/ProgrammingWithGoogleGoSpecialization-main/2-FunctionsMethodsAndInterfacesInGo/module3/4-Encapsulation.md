# Encapsulation

## Controlling Access

- can define public functions to allow access to hidden data

```golang
package data
var x int = 1
func Printx(){fmt.Println(x)} // functions that start with an uppercase letter gets exported and can be imported 
```

```golang
package main
import "data"
func main(){
  data.Printx()
}
```

## Controlling Access to Structs

- hide fields of structs by starting field name with a lower-case letter
- define public methods which access hidden data

```golang
package data
type Point struct{
  x float64
  y float64
}
func (p *Point) InitMe(xn, yn float64){
  p.x = xn
  p.y = yn
}
func (p *Point) Scale(v float64){
  p.x *= v
  p.y *= v
}
func (p *Point) PrintMe(){
  fmt.Println(p.x, p.y)
}
```

- needs InitMe() to assign hidden data fields

```golang
package main
import "data"
func main(){
  var p data.Point
  p.InitMe(3,4)
  p.Scale(2)
  p.PrintMe() // 6, 8
}
```

- access to hidden fields only through public methods
