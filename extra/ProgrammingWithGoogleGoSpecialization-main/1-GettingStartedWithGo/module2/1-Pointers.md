# Pointers

- a pointer is an address to data in memory
- `&` operator returns the address of a variable/function
- `*` operator returns data at an address (dereferencing)

```golang
var x int = 1
var y int
var ip *int // ip is pointer to int

ip = &x // ip now points to x
y = *ip // y is now whatever the value of ip, which is 1
```

## `New`

- Alternate way to create a variable/function
- `new()` function creates a variable and returns a pointer to the variable
- variable is initialized to zero

```golang
ptr := new (int)
*ptr = 3
```
