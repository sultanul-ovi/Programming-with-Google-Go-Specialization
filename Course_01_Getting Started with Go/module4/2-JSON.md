# JSON

- all unicode
- human readable
- fairly compact representation
- types can be combined recursively
  - array of structs, struct in struct, ...

## JSON Marshalling

- generating JSON representation from an object

```golang
type struct Person{
  name string
  addr string
  phone string
}
p1 := Person(
  name:"joe",
  addr:"a st.",
  phone: "123"
)
barr, err:= json.Marshal(p1)
```

- `Marshall()` takes a golang object and returns JSON representation as []byte

## JSON Unmarshalling

```golang
var p2 Person
err := json.Unmarshall(barr, &p2)
```

- `Unmarshall()` converts a JSON []byte into a go object
- pointer to Go object is passed to `Unmarshal()`
- object must "fit" JSON []byte
