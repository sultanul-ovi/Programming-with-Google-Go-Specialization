# Support for Classes

## No `Class` Keyword

- most OO languages have a class keyword
- data fields and methods are defined inside a class block

## Associating Methods with Data

- method has a receiver type that it is associated with
- use dot notation to call the method

```golang
type MyInt int
func (mi MyInt) Double() int{ // mi is an implicit argument
  return int(mi*2)
}
func main(){
  v := MyInt(3)
  fmt.Println(v.Double()) // 6
}
```

- object `v` is an implicit argument to the method
  - call by value
