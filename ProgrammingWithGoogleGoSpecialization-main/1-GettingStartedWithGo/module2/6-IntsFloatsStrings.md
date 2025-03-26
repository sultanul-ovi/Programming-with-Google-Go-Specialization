# Ints, Floats, Strings

## Type conversions

- most binary operations ne ed operands of the same type
  - including assignments

  ```golang
  var x int32 = 1
  var y int16 = 2
  x = y // ERROR
  ```

- convert type with T() operations

  ```golang
  x = int32(y);
  ```

## Floating Point

- float32 ~6 digits of precision
- float64 ~15 digits of precision
- expressed using decimals or scientific notation
  - var x float64 = 123.45
  - var y float64 = 1.234e2
- complex numbers represented as two floats: real and imaginary numbers
  - var z complex128 = complex(2,3)

## ASCII and Unicode

- american standard code for information exchange
- character coding - each charatcter is associated with an (7) 8-bit number
  - 'A' = 0x41
- unicode is a 32-bit character code
- UTF-8 is variable length
  -8-bit UTF codes are same as ASCII
- code points - unicode characters
- run - a code point in Go

## Strings

- sequence of arbitraty bytes
  - read-only
    - can only be mutated through a copy
  - often meant to be printed
- string literal - notated by double quotes
  - `x:= "Hi there"`
- each byte is a rune (UTF-8 code point)
