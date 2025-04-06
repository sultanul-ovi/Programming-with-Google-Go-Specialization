# Interface vs. Concrete Types

## Concrete Types

- specify the exact representation of the data and methods
- complete method implementation is included

## Interface Types

- specifies some method signatures
- implementations are abstracted
- can be treated like other values
  - assigned to variables
  - passed, returned
- interface values have two components

1. dynamic type: concrete type which it is assigned to
2. dynamic value: value of the dynamic type

## Defining an Interface Type

```golang
type Speaker interface {Speak()}

type Dog Struct {name string}

func (d Dog) Speak(){
  fmt.Println(d.name)
}

func main(){
  var s1 Speaker
  var d1 Dog{"Brian"}
  s1 = d1
  s1.Speak()
}
```

- dyamic type is `Dog`, dynamic value is `d1`

## Interface with Nil Dynamic Value

- an interface can have a nil dynamic value

```golang
var s1 Speaker
var d1 *Dog
s1 = d1
```

- `d1` has no conrete value yet
- `s1` has a dynamic type but no dynamic value

## Nil Dynamic Value

- can still cal the `Speak()` method of `s1`
- doesn't need a dynamic value to call
- need to check inside the method

```golang
func (d *Dog) Speak(){
  if d == nil{
    fmt.Println("<noise>")
  } else{
    fmt.Println(d.name)
  }
}

func main(){
  var s1 Speaker
  var d1 *Dog
  s1 = d1
  s1.Speak()
}
```

- interface with `nil dynamic type`
- very different from an interface with a `nil dynamic value`

```golang
func main(){
  var s1 Speaker
  var d1 *Dog
  s1 = d1
  s1.Speak()
}
```

- *Nil dynamic value* and *valid dynamic type*
- can call a method since type is known

- *Nil dynamic type*

```golang
func main(){
  var s1 Speaker // no method associated so it'll be a runtime error
}
```

- cannot call a method, runtime error
