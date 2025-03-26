# Comments, Printing, Integers

## Comments

- comments are text for understandability
- ignored by the compiler
- single-line comments

```golang

// this is a comment
var x int // this is a comment
```

- block comments

```golang
/*
comment 1
comment 2
*/
var x int
```

## Printing

- import from the `fmt` package
- `fmt.Printf()` (fmt.Println)
- format strings are good for formatting
- conversion characters for each argument `fmt.Printf("%s", x)`

## Integers

- generic int declaraction
  - var x int
- different lengths and signs
  - int8, int16, int32, int64
  - uint8, uint16, uint32, uint64
    - the number represents the bit
    - u- means unsigned
- binary operators
  - arithmetic: +,-,*
  - comparison: ==, <=, >=
  - boolean: && ||
