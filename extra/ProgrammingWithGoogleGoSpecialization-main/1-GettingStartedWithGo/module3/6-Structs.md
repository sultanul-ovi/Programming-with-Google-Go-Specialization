# Structs

- aggregate data type
- groups together other objects of arbitrary type

## Example

- Name, Address, Phone
- Option 1: Have 3 separate variables, programmer remembers that they are related
- Option 2: Make a single struct which contains all 3 vars

```golang
type struct Person{
  name string
  addr string
  phone string
}
var p1 Person
```

- each property is a field
- `p1` contains values for all fields

## Accessing Struct Fields

- use dot notation

```golang
p1.name = "joe"
x = p1.addr
```

## Initializing Structs

- can use new()
- initializes fields to zero

```golang
p1:= new(Person)
```

- can initialize using a struct literal

```golang
p1 := Person(
  name:"joe",
  addr:"a st.",
  phone:"123")
```
