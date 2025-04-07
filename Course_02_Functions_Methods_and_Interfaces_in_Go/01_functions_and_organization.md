# Go Programming - Functions, Methods, and Interfaces

This module introduces key Go programming concepts such as functions, function types, methods, and interfaces. You'll also explore object-oriented programming in Go through class-like struct instantiation and interface implementation.

## Learning Objectives

- Understand basic features and purposes of functions
- Learn benefits of using pointers with functions
- Differentiate between slices and arrays as function arguments
- Use functions and slices to implement a sorting routine

---

## Topic 1.1: Why Use Functions?

### What is a Function?

- A named group of instructions
- Can be executed by _calling_ the function

```go
func main() {
  PrintHello()
}

func PrintHello() {
  fmt.Printf("Hello, world!\n")
}
```

### Key Characteristics

- All Go programs must have a `main()` function (execution entry point)
- Other functions are defined using the `func` keyword and called explicitly

---

## Reusability

- Functions can be written once and called many times
- Encourages modular code
- Example use cases:
  - `ThresholdImage()` in graphics editing
  - `QueryDBase()` in database apps
  - `ChangeKey()` in music apps

---

## Abstraction

- Hides internal implementation details
- Focus is on what the function does, not how
- Enhances readability and maintainability

### Example:

```go
func FindPupil() {
  GrabImage()
  FilterImage()
  FindEllipses()
}
```

- Even without internal logic, names provide understanding
- Naming matters: `x()`, `y()` are less meaningful than `GrabImage()`

---

## Function Anatomy in Go

```go
func FunctionName(arg1 Type1, arg2 Type2) ReturnType {
  // code block
  return result
}
```

- Starts with `func` keyword
- Includes name, parameters, return type, and code block

---

## Summary

Functions in Go:

- Provide structure and modularity
- Promote code reuse
- Enable abstraction

They are essential for building scalable and maintainable Go applications.

---

## Topic 1.2: Function Parameters and Return Values

## Function Parameters

- Functions often need input data to operate on
- Parameters are listed in parentheses after the function name in the declaration
- Arguments are supplied during the function call

### Example:

```go
func foo(x int, y int) {
  fmt.Print(x * y)
}

func main() {
  foo(2, 3) // prints 6
}
```

### Notes:

- Parameters become local variables inside the function
- `2` is bound to `x`, `3` is bound to `y`

---

## Parameter Variations

- **No parameters:** use empty parentheses

```go
func foo() {}
```

- **Same type arguments:** list them together

```go
func foo(x, y int) {}
```

---

## Return Values

- Functions can return one or more values
- Return type(s) follow the parameter list in the declaration
- Use `return` to specify what value to send back

### Single Return Example:

```go
func foo(x int) int {
  return x + 1
}

func main() {
  y := foo(1) // y = 2
}
```

### Notes:

- `return x + 1` sends value back
- `y := foo(1)` captures the return value

---

## Multiple Return Values

- List multiple return types in parentheses
- Return values separated by commas

### Example:

```go
func foo(x int) (int, int) {
  return x, x + 1
}

func main() {
  a, b := foo(1) // a = 1, b = 2
}
```

---

## Summary

- Parameters supply input to functions
- Arguments are actual values passed during calls
- Return values capture results from function execution
- Go allows multiple return values with clear syntax

---

## Topic 1.3: Call by Value and Call by Reference

## Call by Value

- Arguments are **copied** to parameters
- Modifying parameters does **not** affect caller variables

### Example:

```go
func foo(y int) {
  y = y + 1
}

func main() {
  x := 2
  foo(x)          // passes a copy of x
  fmt.Println(x)  // prints 2
}
```

### Tradeoffs

- ✅ Advantage: Data encapsulation
  - Functions cannot modify caller variables
- ❌ Disadvantage: Copying time
  - Large data structures are costly to copy

---

## Call by Reference

- Pass a **pointer** to the function
- Function has **direct access** to caller variable in memory

### Example:

```go
func foo(y *int) {
  *y = *y + 1
}

func main() {
  x := 2
  foo(&x)         // passes pointer to x
  fmt.Println(x)  // prints 3
}
```

### Tradeoffs

- ✅ Advantage: Reduced copying
  - Only the pointer is copied
- ❌ Disadvantage: Less encapsulation
  - Functions can modify external variables

---

## Summary

- **Call by value**: default in Go, safe but potentially inefficient
- **Call by reference**: done via pointers, efficient but riskier

Use call by reference only when the function **must modify** the caller's variable.

---

## Topic 1.4: Passing Arrays and Slices

## Passing Array Arguments

- Arrays are passed by value (copied)
- Copying large arrays can be inefficient

### Example:

```go
func foo(x [3]int) int {
  return x[0]
}

func main() {
  a := [3]int{1, 2, 3}
  fmt.Print(foo(a)) // prints 1
}
```

---

## Passing Array Pointers

- Possible to pass a pointer to an array
- Allows function to modify the original array

### Example:

```go
func foo(x *[3]int) {
  (*x)[0] = (*x)[0] + 1
}

func main() {
  a := [3]int{1, 2, 3}
  foo(&a)
  fmt.Print(a) // prints [2 2 3]
}
```

- ✅ Avoids copying large data
- ❌ Requires manual dereferencing — messy

---

## Pass Slices Instead

- Slices contain a pointer, length, and capacity
- Passing a slice copies the **structure**, not the underlying array
- Cleaner and preferred in Go

### Example:

```go
func foo(sli []int) {
  sli[0] = sli[0] + 1
}

func main() {
  a := []int{1, 2, 3}
  foo(a)
  fmt.Print(a) // prints [2 2 3]
}
```

---

## Summary

- Arrays are copied when passed — avoid for large data
- Pointer-to-array enables modification but is clunky
- Use **slices** for efficiency and simplicity

Next: Sorting slices using functions.

---

## Topic 1.5: Well-Written Functions

## Understandability

- Code is composed of **functions** and **data**
- Good organization means:
  - You can find the function for a feature quickly
    - Example: "Where is the function that blurs the image?"
  - You can trace where data is used or modified
    - Example: "Where do you modify the record list?"

---

## Debugging Principles

- When code crashes in a function:
  - Two possible root causes:
    1. **Function logic is wrong**
       - e.g. Sorting a slice in the wrong order
    2. **Input data is wrong**
       - e.g. Slice has bad data before sorting

---

## Supporting Debugging

- Functions must be easy to read and verify
  - Check if actual behavior matches expected behavior
- Data must be traceable
  - You should know:
    - Where it was defined
    - Where it was modified
    - Where it was used

### Pitfall: Global Variables

- ❌ Global variables make it hard to trace data flow
- ✅ Prefer passing data via parameters

---

## Summary

- A well-written function is clear, concise, and easy to verify
- Good organization helps both **feature discovery** and **bug tracing**
- Avoid globals unless necessary to keep data flows predictable

Next: Methods in Go - Attaching Functions to Types

---

## Topic 1.6: Guidelines for Functions

## Function Naming

- Names should clearly convey behavior
- Parameter names matter too

### Example:

```go
func ProcessArray(a []int) float {}
func ComputeRMS(samples []float64) float64 {}
```

- "ComputeRMS" is clear for domain experts (RMS = Root Mean Square)
- "samples" is more informative than "a"

---

## Functional Cohesion

- Functions should do only **one logical operation**
- Context matters—operations should align with application domain

### Good Examples:

- `PointDist()`
- `DrawCircle()`
- `TriangleArea()`

### Bad Example:

- Merging `DrawCircle()` and `TriangleArea()` into one function
  - Unclear name
  - Poor separation of concerns

---

## Few Parameters

- More parameters = more complexity
- Harder to debug and trace
- May indicate poor cohesion

### Example:

- Merging unrelated operations results in needing more arguments

---

## Reducing Parameter Count

- **Use structs** to group related data

### Bad Solution:

```go
func TriangleArea(x1, y1, z1, x2, y2, z2, x3, y3, z3 float64) float64 {}
```

### Good Solution:

```go
type Point struct {
  x, y, z float64
}

func TriangleArea(p1, p2, p3 Point) float64 {}
```

### Even Better:

```go
type Triangle struct {
  p1, p2, p3 Point
}

func TriangleArea(t Triangle) float64 {}
```

---

## Summary

- Clear function and parameter names improve understandability
- Functions should do **one thing only**
- Keep parameter lists short
- Group related data using structs

Next: Go Methods – Attaching Functions to Types

---

## Topic 1.7: Function Guidelines

## Function Complexity

- Simpler functions are easier to debug
- **Length** is a basic indicator of complexity
- But even short functions can be complex

---

## Function Length

- Goal: Keep individual functions simple
- Use **Function Call Hierarchy** to manage complex logic

### Example Hierarchy:

- **Option 1:** One function, 100 lines – hard to debug
- **Option 2:** One controller function (few lines) calling two 50-line functions – easier to debug and understand

---

## Control-Flow Complexity

- Describes the number of conditional paths in a function

### Example:

```go
func foo() {
  if a == 1 {
    if b == 1 {
      // ...
    }
  }
  // ...
}
```

- 3 control-flow paths:
  - a != 1
  - a == 1 and b != 1
  - a == 1 and b == 1

---

## Partitioning Conditionals

- Use helper functions to separate logic

### Example:

```go
func foo() {
  if a == 1 {
    CheckB()
  }
  // ...
}

func CheckB() {
  if b == 1 {
    // ...
  }
}
```

- Now each function has only 2 control-flow paths
- Improves readability and debugging

---

## Summary

- Avoid large, complex functions
- Use function hierarchies to break down tasks
- Partition logic to reduce control-flow complexity

---
