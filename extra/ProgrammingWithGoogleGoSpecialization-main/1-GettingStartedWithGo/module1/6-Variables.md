# Variables

## Naming

- names are needed to refer to programming constructs
  - variables, functions...
- names must start with a letter
- any number of letterse, digits, underscores
- case sensitive
- don't use keywords
  - i.e. `if`, `case`, `package`, ...

## More on Variables

- data stored in memory
- must have a name and a type
- all variables must have declarations
- most basic declaration:

```golang
var x int // keyword name type
```

- can declare many on the same line:

```golang
var x, y int
```

## Variable Types

- type defines the values a variable may take and operations that can be performed on it
- Integer
  - only integral values
  - integer arithmetic (+,-,*...)
- Floating Point
  - fractional (decimal) values
  - floating point arithmetic (may use different hardware)
- Strings
  - byte (character) sequences
  - string comparison, search, concatenation, ...
