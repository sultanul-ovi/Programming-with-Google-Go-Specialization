# Functions and Object-Oriented Programming in Go

## Learning Objectives

- Understand the basic concepts of objects and structs in Go
- Differentiate structs in Go from classes in OOP languages
- Learn to use methods with different data types
- Build an interactive Go routine using structs, methods, and instances

---

## Classes and Encapsulation

- A class represents a group of related data fields and functions
- Example: A `Point` class might include:
  - **Fields:** `x`, `y`
  - **Methods:** `DistToOrigin()`, `Quadrant()`, `AddXOffset()`, `AddYOffset()`, `SetX()`, `SetY()`

### Object

- An instance of a class/struct containing actual data

### Encapsulation

- Protects internal data
- Enforces access via methods to avoid misuse
- Example:
  - Instead of allowing manual changes to `x` and `y`, define `DoubleDist()`

---

## Structs in Go (Class Equivalent)

- Go does not have a `class` keyword
- Structs define composite types with multiple fields

```go
package main

import "fmt"

type MyInt int

func (mi MyInt) Double() int {
  return int(mi * 2)
}

func main() {
  v := MyInt(2)
  fmt.Println(v.Double())
}
```

### Implicit Method Argument

- The receiver (`MyInt`) is implicitly passed by value

---

## Structs and Methods

```go
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p Point) DistToOrig() float64 {
    t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
    return math.Sqrt(t)
}

func main() {
    p1 := Point{3, 4}
    fmt.Println(p1.DistToOrig())
}
```

---

## Encapsulation in Go

- Use lowercase field names for private access
- Only accessible within the package

### data package

```go
package data

import "fmt"

type Point struct {
	x, y float64
}

func (p *Point) InitMe(xn, yn float64) {
	p.x = xn
	p.y = yn
}

func (p *Point) Scale(v float64) {
	p.x *= v
	p.y *= v
}

func (p *Point) PrintMe() {
	fmt.Println(p.x, p.y)
}
```

### main package

```go
package main

import "data"

func main() {
	var p data.Point
	p.InitMe(3, 4)
	p.Scale(2)
	p.PrintMe()
}
```

---

## Pointer Receivers

- Avoids copying the object
- Allows method to modify the original data

```go
func (p *Point) OffsetX(v float64) {
	p.x += v
}
```

### Usage

```go
func main() {
	p := Point{3, 4}
	p.OffsetX(5)
	fmt.Println(p.x)
}
```

- No need for explicit referencing or dereferencing

### Good Practice

- Use either pointer or non-pointer receivers consistently
- Pointer receivers allow mutation

---

## Final Go Routine Task (Interactive)

Create a Go routine that:

- Uses structs to define a geometric object
- Initializes instances
- Queries/modifies properties via methods
- Enforces encapsulation using pointer receivers and lowercase fields

---

# Module 3: Object-Orientation in Go

## Topic 1.1: Classes and Encapsulation

### Classes

- A class is a **collection of data fields and methods** that serve a well-defined responsibility.
- **Example:** `Point` class
  - Use case: Geometry program
  - **Data Fields:**
    - `x` coordinate
    - `y` coordinate
  - **Methods:**
    - `DistToOrigin()`
    - `Quadrant()`
    - `AddXOffset()`
    - `AddYOffset()`
    - `SetX()`
    - `SetY()`
- A class is a **template** for objects:
  - It specifies what data and functions an object should have
  - Does **not contain actual data**, only field definitions

### Object

- An **object is an instance** of a class
- Contains actual data values
- Based on the class template
- **Example:**
  - Class: `Point`
  - Objects: Points with coordinates (0,0), (6,0), (5,5)
  - Each object has actual `x` and `y` values

### Key Distinction

- **Class** → Template with data fields and method definitions
- **Object** → Instance with real data that fills in the class template

---

### Encapsulation

- Hides internal data from the external programmer
- Data can only be modified through defined methods
- Protects data consistency and integrity

#### Why Use Encapsulation?

- Prevent misuse or accidental errors by programmers
- Relieves the user from maintaining data consistency
- Ensures correct usage by enforcing access via methods

#### Example: Doubling Distance to Origin

- **Correct Approach:**
  - Define method `DoubleDist()` that modifies both `x` and `y`
- **Unsafe Approach:**
  - Let programmer directly modify `x` and `y`
  - Risk of inconsistency (e.g., doubling one but not the other)

#### Summary

- Encapsulation uses **methods as a barrier** to access or modify internal state
- Encourages **safe, consistent programming practices**

---

# Module 3: Object-Orientation in Go

## Topic 1.2: Support for Classes

### No `class` Keyword

- Go does **not** have a `class` keyword
- Most OOP languages define data and methods inside a `class` block
- Example from Python (for contrast):
  ```python
  class Point:
      def __init__(self, x, y):
          self.x = x
          self.y = y
  ```
- In Go, classes are mimicked using **structs** and **methods** associated via receiver types

---

### Associating Methods with Data (Receiver Types)

- A **method** in Go is defined with a **receiver type** to bind it to a specific data type
- This is how Go achieves class-like behavior

#### Example:

```go
package main

import "fmt"

type MyInt int

func (mi MyInt) Double() int { // mi is the receiver (passed implicitly)
    return int(mi * 2)
}

func main() {
    v := MyInt(3)
    fmt.Println(v.Double()) // Output: 6
}
```

- `MyInt` is a custom type (alias for `int`)
- `Double()` is a method associated with `MyInt`
- `v` is an instance of `MyInt` and `v.Double()` invokes the method

---

### Key Concepts

#### Method Association

- A method must be defined in the **same package** as the type
- Cannot define methods on standard types like `int`, `string`, etc.

#### Dot Notation

- Method is invoked using the syntax: `object.method()`
- The **type** of the object determines which method gets called
- Dot notation allows the Go compiler to identify the correct method by type

#### Receiver Type

- Specified before the method name:
  ```go
  func (receiverType TypeName) MethodName() {...}
  ```
- Binds the method to the specified type

---

### Implicit Method Argument

- Although method signature shows no argument, the receiver object is passed implicitly
- In `v.Double()`, `v` is passed to `Double` as an **implicit argument**
- **Call by value**: Go makes a copy of `v` and passes it to the method

#### Why It Matters

- Since it's passed by value, the method cannot modify the original object
- If mutation is required, use **pointer receivers**

---

### Summary

- Go doesn't have classes but supports OOP through **structs** and **receiver methods**
- Methods are bound to data types using receiver syntax
- Method calls pass the object as an implicit argument (by value unless using a pointer)

---

# Module 3: Object-Orientation in Go

## Topic 1.3: Structs and Methods (Support for Classes)

### Structs in Go

- `struct` types are used to **compose multiple data fields** into one logical unit
- This mimics traditional class-like data aggregation

#### Example:

```go
package main

type Point struct {
  x float64
  y float64
}
```

- The `Point` struct combines two float64 fields: `x` and `y`
- Structs allow any amount of arbitrary data to be grouped

---

### Structs with Methods

- Go allows **methods to be associated** with struct types using receiver types
- Struct + methods = class-like behavior

#### Example:

```go
package main

import (
  "fmt"
  "math"
)

type Point struct {
  x float64
  y float64
}

func (p Point) DistToOrig() float64 {
  t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
  return math.Sqrt(t)
}

func main() {
  p1 := Point{3, 4}
  fmt.Println(p1.DistToOrig()) // Output: 5
}
```

- `DistToOrig` is a method bound to `Point` via the receiver `(p Point)`
- When `p1.DistToOrig()` is called:
  - `p1` is passed implicitly to the method as `p`
  - Method calculates Euclidean distance to origin using Pythagorean theorem

---

### Summary

- Structs let you group complex data
- Methods can be defined on structs using receiver types
- This creates **class-like** constructs in Go without having an actual `class` keyword

---

# Module 3: Object-Orientation in Go

## Topic 2.1: Encapsulation

### Controlling Access to Data

- Go supports **encapsulation** by hiding data and allowing **controlled access**
- Even when data is private, controlled interaction is enabled through **public methods**

#### Example: Private Variable with Public Access

```go
package data

import "fmt"

var x int = 1

func PrintX() {
  fmt.Println(x)
}
```

```go
package main

import "data"

func main() {
  data.PrintX() // Outputs: 1
}
```

- `x` is hidden (lowercase)
- `PrintX()` is exported (uppercase) and provides controlled access to `x`

---

### Encapsulation for Structs

- Struct fields can be **unexported** (private) using lowercase field names
- Use **public methods** to access or modify these fields

#### Example: Struct with Controlled Access

```go
package data

import "fmt"

type Point struct {
  x float64
  y float64
}

func (p *Point) InitMe(xn, yn float64) {
  p.x = xn
  p.y = yn
}

func (p *Point) Scale(v float64) {
  p.x *= v
  p.y *= v
}

func (p *Point) PrintMe() {
  fmt.Println(p.x, p.y)
}
```

- `x` and `y` are private
- `InitMe`, `Scale`, and `PrintMe` are public methods to interact with `x` and `y`

#### Usage in Main Package

```go
package main

import "data"

func main() {
  var p data.Point
  p.InitMe(3, 4)
  p.Scale(2)
  p.PrintMe() // Output: 6 8
}
```

- External code cannot access `p.x` or `p.y` directly
- Can only modify or view through public methods like `InitMe`, `Scale`, and `PrintMe`

---

### Summary

- Go uses **capitalization** to define visibility (public vs private)
- **Encapsulation** ensures that internal struct data is accessed in a safe and predictable way
- Encourages better abstraction and prevents misuse of internal fields

---

# Module 3: Object-Orientation in Go

## Topic 2.2: Pointer Receivers

### Limitations of Value Receivers

- Receiver object is **passed implicitly** to methods
- By default, Go uses **call by value**
- This means:
  - The method receives a **copy** of the object
  - Any changes to the receiver's data **do not affect** the original object

#### Example (Ineffective Modification):

```go
func main() {
  p1 := Point{3, 4}
  p1.OffsetX(5) // This will not modify p1.x
}
```

- `OffsetX` receives a **copy** of `p1`, so changes are lost after execution

---

### Performance Issue with Large Receivers

- Copying large structs is expensive in memory and time

#### Example:

```go
type Image [100][100]int

func main() {
  i1 := GrabImage()
  i1.BlurImage() // Copies 10,000 ints
}
```

- Struct `Image` has 10,000 integers
- Calling a method like `BlurImage()` creates a full copy on the stack

---

### Pointer Receivers

- Solution: Use **pointer receiver** in method definition
- Passes a reference instead of a copy
- Allows **modifying original data**

#### Example:

```go
func (p *Point) OffsetX(v float64) {
  p.x = p.x + v
}
```

- `*Point` indicates a pointer receiver
- Method modifies the actual `Point` instance
- No need for explicit dereferencing when calling method

#### Call:

```go
p := Point{3, 4}
p.OffsetX(5) // p.x is now 8
```

---

### Summary

- Use **pointer receivers** to:
  - Modify the original object
  - Avoid performance cost of copying large structs
- Receiver methods still use **dot notation** even when working with pointers

---

# Module 3: Object-Orientation in Go

## Topic 2.3: Point Receivers, Referencing, and Dereferencing

### No Need for Explicit Dereferencing

- When using **pointer receivers**, you **do not** need to dereference manually
- Go automatically dereferences pointer types with the dot (`.`) operator

#### Example:

```go
func (p *Point) OffsetX(v int) {
  p.x += v // No need to write (*p).x
}
```

- `p` is a pointer, but Go handles dereferencing
- Cleaner syntax, reduces error and clutter

---

### No Need for Explicit Referencing on Method Call

- Even if the receiver expects a pointer, Go **automatically references** a value when calling the method

#### Example:

```go
func main() {
  p := Point{3, 4}       // p is a value, not a pointer
  p.OffsetX(5)           // Works fine without &p
  fmt.Println(p.x)       // Output: 8
}
```

- Go recognizes that `OffsetX` expects a `*Point` and inserts `&p` automatically

---

### Best Practice

- For consistency, choose **one style per type**:
  - All methods use pointer receivers, OR
  - All methods use non-pointer receivers
- **Pointer receivers**:
  - Allow modification of the original object
  - Avoid unnecessary copying for large structs
- Mixing styles can lead to confusion and bugs

---

### Summary

- Go simplifies referencing and dereferencing when using methods with pointer receivers
- Stick to a **consistent receiver style** to improve readability and maintainability

---
