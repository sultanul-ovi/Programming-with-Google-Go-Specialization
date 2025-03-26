# Arrays

- fixed-length series of elements of a chosen type
- elements accessed using subscript notation []
- indices start at 0
- elements initialized to zero value
  - zero, empty string, etc.

## Array Literal

- an array pre-defined with values
  - `var x [5]int = [5]{1,2,3,4,5}`
- length of literal must be length of array
- `...` for size in array literal infers size from number of initializers
  - `x:=[...]int{1,2,3,4} // the size is 4`

## Iterating through Arrays

- use a for loop with the range

```golang
x:= [3]int{1,2,3}
for i, v range x{
  fmt.Printf("i:%d, val:%d, i,v)
}
```

- range returns two values
- index and element at index
