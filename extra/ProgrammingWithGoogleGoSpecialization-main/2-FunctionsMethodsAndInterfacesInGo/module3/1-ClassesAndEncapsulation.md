# Classes and Encapsulation

## Classes

- collection of data fields and functions that share a well-define responsibility
- example: `Point` class
  - used in a geometry program
  - data: x, y coordinates
  - functions:
    - `DistToOrigin()`, `Quadrant()`
    - `AddXOffSet()`, `AddYOffSet()`
    - `SetX()`, `SetY()`
- Classes are a template
- contain data fields, not data

### Object

- instance of a class
- contains real data
- example: point class

## Encapsulation

- data can be protected from the programmer
- data can be accessed only using methods
- maybe we don't trust the programmer to keep data consistent
- example: double distance to origin
  - option 1: make method DoubleDist()
  - option 2: trust programmer to double X and Y directly
