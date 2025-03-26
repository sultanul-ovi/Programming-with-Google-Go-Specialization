# Variable Initialization

## Type Declarations

- defining an alias (alternate name) for a type
- may improve clarity

```golang
type Celsius float64
```

```golang
type IDnum int
```

- can declare variables using the type alias instead

```golang
var temp Celsius
```

```golang
var pid IDnum
```

## Initializing Variables

- initialize in the declaration

```golang
var x int = 100
var x = 100
```

- initialize after the declaration

```golang
var x int
x = 100
  ```

- uninitialized variables have a zero value

```golang
var x int // x = 0
```

```golang
var x string // x = ""
```

- short variable declarations
  - can perform a declaration and initialization together with the `:=` operator
  - variable is declared as the type of expression on the right hand side
  - can only do this inside a function

```golang
x:=100
```
