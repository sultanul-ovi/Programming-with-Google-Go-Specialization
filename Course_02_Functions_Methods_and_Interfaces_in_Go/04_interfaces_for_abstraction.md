# Go Interfaces and Polymorphism

## Module Overview

This module wraps up the course by covering:

- Interfaces in Go
- Polymorphism using interfaces
- Concepts of overriding (Go-style)
- Implementing polymorphism through structs and methods
- Final interactive Go routine to create and query class instances

---

## Learning Objectives

- Understand interfaces in Go
- Understand polymorphism and method overriding
- Create and use structs, methods, and interfaces
- Develop a Go routine for creating and querying object properties

---

## Interfaces in Go

- Define a set of method signatures.
- A type implicitly satisfies an interface by implementing all its methods.

```go
// Example
package main

import "fmt"

type Speaker interface {
    Speak()
}

type Cat struct{}
type Dog struct{}

func (c Cat) Speak() {
    fmt.Println("meow")
}

func (d Dog) Speak() {
    fmt.Println("woof")
}

func MakeItSpeak(s Speaker) {
    s.Speak()
}

func main() {
    c := Cat{}
    d := Dog{}
    MakeItSpeak(c)
    MakeItSpeak(d)
}
```

---

## Polymorphism in Go

- Different implementations of a method for different types.
- Example: `Area()` behaves differently for `Rectangle` and `Triangle`.

```go
type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width, Height float64
}

type Triangle struct {
    Base, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (t Triangle) Area() float64 {
    return 0.5 * t.Base * t.Height
}
```

---

## Overriding in Go (via Interface Implementation)

- Go doesn't support class-based inheritance.
- Instead, types implement interfaces with methods.
- Method redefinition = overriding.

```go
type Speaker interface {
    Speak()
}

func (c Cat) Speak() {
    fmt.Println("meow")
}

func (d Dog) Speak() {
    fmt.Println("woof")
}
```

---

## Final Go Routine

- Create instances (objects) of types (structs)
- Query their properties interactively

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Animal interface {
    Speak()
    Type() string
}

type Cat struct{}
func (c Cat) Speak() { fmt.Println("meow") }
func (c Cat) Type() string { return "Cat" }

type Dog struct{}
func (d Dog) Speak() { fmt.Println("woof") }
func (d Dog) Type() string { return "Dog" }

func main() {
    animals := make(map[string]Animal)
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Println("Enter commands: create [name] [type] | query [name] | exit")
    for {
        fmt.Print("> ")
        scanner.Scan()
        input := strings.Fields(scanner.Text())
        if len(input) == 0 {
            continue
        }
        cmd := input[0]

        switch cmd {
        case "create":
            if len(input) != 3 {
                fmt.Println("Usage: create [name] [cat/dog]")
                continue
            }
            name, kind := input[1], input[2]
            switch strings.ToLower(kind) {
            case "cat":
                animals[name] = Cat{}
            case "dog":
                animals[name] = Dog{}
            default:
                fmt.Println("Unknown type")
            }

        case "query":
            if len(input) != 2 {
                fmt.Println("Usage: query [name]")
                continue
            }
            name := input[1]
            if animal, ok := animals[name]; ok {
                fmt.Printf("%s says: ", animal.Type())
                animal.Speak()
            } else {
                fmt.Println("Animal not found")
            }

        case "exit":
            return

        default:
            fmt.Println("Unknown command")
        }
    }
}
```

---

## Summary

- Go uses interfaces for abstraction and polymorphism.
- Polymorphism in Go = different structs implementing the same interface.
- Go doesn’t support inheritance but achieves similar functionality with interfaces and composition.
- Final program lets users create and interact with different object types dynamically.

---

# Go Interfaces: Shape2D Example

## Module 4: Interfaces for Abstraction

### Topic 1.2: Interfaces

---

## Interfaces

- A **set of method signatures**:
  - method names
  - parameters and return types
- **No implementation** is defined in the interface
- Used to express **conceptual similarity**

### Example: Shape2D Interface

All 2D shapes must implement:

- `Area() float64`
- `Perimeter() float64`

```go
// Defining an interface
// Shape2D represents any 2D shape with area and perimeter methods

package main

import "fmt"

type Shape2D interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width, Height float64
}

type Triangle struct {
    Base, Height float64
    A, B, C float64 // sides for perimeter
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func (t Triangle) Area() float64 {
    return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
    return t.A + t.B + t.C
}

func printShapeInfo(s Shape2D) {
    fmt.Printf("Area: %.2f\n", s.Area())
    fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
    r := Rectangle{Width: 10, Height: 5}
    t := Triangle{Base: 6, Height: 4, A: 5, B: 5, C: 6}

    fmt.Println("Rectangle:")
    printShapeInfo(r)

    fmt.Println("\nTriangle:")
    printShapeInfo(t)
}
```

---

## Satisfying an Interface

- A type satisfies an interface **implicitly** by implementing **all methods** with the correct **signatures**
- No need to explicitly declare satisfaction
- Types can have **additional data or methods**

### Comparison with Inheritance:

- Similar benefits to inheritance + overriding
- Achieved with interface + method implementations

---

## Summary

- Go uses **interfaces** to achieve polymorphism
- Interfaces define method signatures only
- Any type that implements the required methods satisfies the interface
- This provides a clean, implicit way to model behavior without inheritance

---

# Interface vs. Concrete Types in Go

## Module 4: Interfaces for Abstraction

### Topic 1.3: Interface vs. Concrete Types

---

## Concrete Types

- Fully specify **data** and **method implementations**
- Methods with the type as the receiver are fully defined

## Interface Types

- Contain only **method signatures**, no data
- Abstract the implementation
- Used to express **behavior** shared across types
- Interface values are composed of:
  1. **Dynamic Type**: the actual concrete type assigned
  2. **Dynamic Value**: the instance (value) of that concrete type

---

## Example: Concrete Type Assigned to Interface

```go
package main

import "fmt"

type Speaker interface {
    Speak()
}

type Dog struct {
    name string
}

func (d Dog) Speak() {
    fmt.Println(d.name)
}

func main() {
    var s1 Speaker
    d1 := Dog{"Brian"}
    s1 = d1
    s1.Speak() // Output: Brian
}
```

- `s1` has **dynamic type** `Dog`, **dynamic value** `d1`

---

## Interface with Nil Dynamic Value

```go
package main

import "fmt"

type Speaker interface {
    Speak()
}

type Dog struct {
    name string
}

func (d *Dog) Speak() {
    if d == nil {
        fmt.Println("<noise>")
    } else {
        fmt.Println(d.name)
    }
}

func main() {
    var s1 Speaker
    var d1 *Dog // nil pointer
    s1 = d1
    s1.Speak() // Output: <noise>
}
```

- `s1` has dynamic type `*Dog`, **but no dynamic value**
- Method can still be called, as long as the type is known
- **Check `nil` inside method** to avoid runtime panic

---

## Nil Dynamic Type

```go
package main

type Speaker interface {
    Speak()
}

func main() {
    var s1 Speaker
    s1.Speak() // RUNTIME ERROR: method not associated
}
```

- **No dynamic type**, **no dynamic value**
- **Cannot** call methods — compiler cannot determine the implementation

---

## Summary

- **Concrete types** = full definition (data + methods)
- **Interfaces** = abstract method contracts (no data)
- Interface values = pair of **dynamic type** + **dynamic value**
- Legal:
  - `dynamic type` present, `dynamic value` nil → method call allowed
- Illegal:
  - `dynamic type` missing → method call causes runtime error

---

# Using Interfaces in Go

## Module 4: Interfaces for Abstraction

### Topic 1.4: Practical Use of Interfaces

---

## Why Use Interfaces?

- Express **conceptual similarity** between types
- Allow functions to accept **multiple types** as arguments
- Enable writing **general-purpose functions**

---

## General Usage Pattern

- Want: `foo()` accepts either **type X** or **type Y**
- Define an **interface Z**
- Make both **X and Y** satisfy interface Z
- `foo(z Z)` → accepts any type that satisfies Z

---

## Example Use Case: `FitInYard()`

**Problem**: You want to check if a pool shape fits in a yard:

- Constraints: **Area** and **Perimeter** should be within limits
- Need to accept **any shape** (rectangle, triangle, etc.)

### Define `Shape2D` Interface

```go
package main

type Shape2D interface {
    Area() float64
    Perimeter() float64
}
```

### Implement for Triangle and Rectangle

```go
type Triangle struct {/* fields */}

func (t Triangle) Area() float64     { /* implementation */ return 0 }
func (t Triangle) Perimeter() float64 { /* implementation */ return 0 }

type Rectangle struct {/* fields */}

func (r Rectangle) Area() float64     { /* implementation */ return 0 }
func (r Rectangle) Perimeter() float64 { /* implementation */ return 0 }
```

---

## FitInYard Function

```go
func FitInYard(s Shape2D) bool {
    return s.Area() < 100 && s.Perimeter() < 100
}
```

- Takes **any type** that satisfies `Shape2D`
- Returns `true` if shape fits within area & perimeter constraints

---

## Empty Interface: `interface{}`

- Special predefined interface
- **Specifies no methods**
- **All types satisfy** it
- Used when:
  - You want a function to accept **any type**

### Example

```go
func PrintMe(val interface{}) {
    fmt.Println(val)
}
```

- Accepts `int`, `float`, `string`, `bool`, structs, etc.

---

## Summary

- Use interfaces to **generalize functions** over multiple types
- Define **method-based contracts** that types must satisfy
- `interface{}` is a fallback when **any type** is acceptable

---

# Type Assertions in Go

## Module 4: Interfaces for Abstraction

### Topic 2.2: Type Assertions

---

## Concealing Type Differences

- Interfaces **hide type differences** by focusing on shared method sets
- Example: `FitInYard()` treats all `Shape2D` values the same

```go
func FitInYard(s Shape2D) bool {
    return s.Area() < 100 && s.Perimeter() < 100
}
```

- You **don't need** to know the concrete type if only interface methods are used

---

## When Concrete Type Matters

- In some applications, you **must differentiate** between types
- Example: **drawing functions** in a graphics API may require specific shape types

```go
func DrawRect(r Rectangle) {}
func DrawTriangle(t Triangle) {}
```

- General function `DrawShape(s Shape2D)` needs to call specific draw function

---

## Type Assertions

- Use to **extract the concrete type** from an interface value

```go
func DrawShape(s Shape2D) {
    rect, ok := s.(Rectangle)
    if ok {
        DrawRect(rect)
    }

    tri, ok := s.(Triangle)
    if ok {
        DrawTriangle(tri)
    }
}
```

- `s.(Rectangle)` extracts `Rectangle` if possible
  - `ok == true`: extraction succeeded
  - `ok == false`: `s` is not of that type

---

## Type Switch

- Cleaner alternative when many types need to be handled

```go
func DrawShape(s Shape2D) {
    switch sh := s.(type) {
    case Rectangle:
        DrawRect(sh)
    case Triangle:
        DrawTriangle(sh)
    default:
        fmt.Println("Unknown shape")
    }
}
```

- `s.(type)` provides the concrete type of `s`
- Each `case` handles a different type
- More readable for **multiple concrete types**

---

## Summary

- Interfaces allow generalization by hiding type differences
- Type assertions let you **recover the original type** when needed
- Use `s.(Type)` or `switch s.(type)` to **disambiguate concrete types**
- Essential in cases where behavior depends on **specific type implementations**

---

# Error Handling with Interfaces in Go

## Module 4: Interfaces for Abstraction

### Topic 2.3: Error Handling

---

## The `error` Interface

- Built-in interface in Go for handling errors

```go
type error interface {
    Error() string
}
```

- Any type that defines an `Error() string` method satisfies the `error` interface

### Behavior

- **Success**: `error` value is `nil`
- **Failure**: `error` value is **non-nil**, and `Error()` returns a descriptive message

---

## Example: File Opening

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    f, err := os.Open("test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    fmt.Println("File opened successfully")
}
```

- `os.Open()` returns a file and an error
- Always check `err != nil`
- `fmt.Println(err)` internally calls `err.Error()`

---

## Summary

- `error` is a **standard interface** in Go
- Many functions return an `error` as the **second return value**
- Always **check** if `err != nil`
- Use `fmt.Println(err)` to print the error message via the `Error()` method

---
