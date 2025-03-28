# Maps

- implementation of a hash table
- use make() to create a mapping

```golang
var idMap map[string]int // map[key][value] // a map variable
idMap = make(map[string][int]) // assigns and points to a map
```

- may define a map literal

```golang
idMap:=map[string]int{"joe":123}
```

## Accessing Maps

- referencing a value with [key]
- returns zero if key is not present

```golang
  fmt.Println(idMap["joe"]) // returns 123
```

- adding/replacing a key/value pair

```golang
idMap["jane"] = 456
idMap["jane"] = 789 // "jane" now has a value of 789
```

- deleting a key/value pair

```golang
  delete(idMap, "joe")
```

## Map Functions

- two-value assignment tests for existence of the key

```golang
id, p := idMap["joe"]
```

- `id` is value, `p` is presence of key
- `len()` returns number values

```golang
fmt.Println(len(idMap))
```

## Iterating Through a Map

- use a for loop with the range keyword
- two-value assignment with range

```golang
for key, val := range idMap{
  fmt.Println(key, val)
}
```
