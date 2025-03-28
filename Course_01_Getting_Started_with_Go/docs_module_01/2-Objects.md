# Object-Oriented Programming

- organize your code through encapsulation
- group together related data and functions
- user-defined type which is specific to an application
  - ex. ints have data (the number) and functions (+,-,*,/,%)

### Object Example

- implementing an application performing geometry in 3D
- many functions will operate on points
  - each point has data:
    - x,y,z coordinates
  - points have functions:
    - DisToOrigin();
    - Quadrant();
- point `class` defins data and functions
- point `objects` are instances of class

## Objects in Go

- Go does not use the term `class`
- Go uses `structs` with associated methods
- simplified implementation of classes
  - no inheritance
  - no constructors
  - no generics
- easier to code
