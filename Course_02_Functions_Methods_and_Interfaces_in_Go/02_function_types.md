# Go Programming - First-Class Functions (Module 2, Topic 1.1)

## Learning Objectives

- Identify advanced types, properties, and uses of functions
- Predict outputs of code blocks using function features
- Implement practical routines using Go functions

---

## First-Class Values

In Go, **functions are first-class citizens**, meaning they can be:

- Assigned to variables
- Created dynamically
- Passed as arguments
- Returned from other functions
- Stored in data structures (e.g., slices)

---

## Variables as Functions

You can assign a function to a variable using a function type.

### Example:

```go
var funcVar func(int) int

func incFn(x int) int {
  return x + 1
}

func main() {
  funcVar = incFn
  fmt.Println(funcVar(1)) // prints 2
}
```

---

## Functions as Arguments

Functions can be passed to other functions as parameters.

### Example:

```go
func applyIt(afunct func(int) int, val int) int {
  return afunct(val)
}

func incFn(x int) int { return x + 1 }
func decFn(x int) int { return x - 1 }

func main() {
  fmt.Println(applyIt(incFn, 2)) // prints 3
  fmt.Println(applyIt(decFn, 2)) // prints 1
}
```

---

## Anonymous Functions

You can define a function without naming it, often useful for passing it inline.

### Example:

```go
func applyIt(afunct func(int) int, val int) int {
  return afunct(val)
}

func main() {
  result := applyIt(func(x int) int {
    return x + 1
  }, 2)
  fmt.Println(result) // prints 3
}
```

---

## Summary

- Functions can be dynamically created, assigned, passed, and returned
- Use anonymous functions for inline or temporary behaviors
- These features are foundational for writing flexible, modular code

Next: Returning Functions from Functions

---

# Go Programming - Returning Functions (Module 2, Topic 1.2)

## Functions as Return Values

- Functions can return **other functions**
- Useful for generating functions with custom parameters
- Example: Create a distance-to-origin function where the origin is customizable

---

## Defining a Function That Returns a Function

```go
func MakeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
  fn := func(x, y float64) float64 {
    return math.Sqrt(math.Pow(x - o_x, 2) + math.Pow(y - o_y, 2))
  }
  return fn
}
```

- `o_x`, `o_y` are built into the returned function
- Creates a new function tailored to a specific origin

---

## Using Special-Purpose Functions

```go
func main() {
  Dist1 := MakeDistOrigin(0, 0)
  Dist2 := MakeDistOrigin(2, 3)

  fmt.Println(Dist1(2, 2)) // Distance from (2,2) to origin (0,0)
  fmt.Println(Dist2(2, 2)) // Distance from (2,2) to origin (2,3)
}
```

- `Dist1` and `Dist2` are customized distance calculators
- Behavior is controlled by their original parameters

---

## Function Environment

- The **environment** of a function includes all variable names that are accessible
- Lexical scoping: includes variables defined in the same block as the function

### Example:

```go
var x int
func foo(y int) {
  z := 1
  // x, y, z all accessible inside foo
}
```

---

## Closures

- **Closure** = function + its environment
- When functions are **passed or returned**, their environment follows them
- Allows functions to "remember" variables from their creation context

### Example Context:

In `MakeDistOrigin`, variables `o_x` and `o_y` are part of the closure for `fn()`.

```go
// When Dist1 is called later, it still has access to o_x and o_y used during creation
Dist1 := MakeDistOrigin(0, 0)
fmt.Println(Dist1(2, 2)) // Uses closure: o_x = 0, o_y = 0
```

---

## Summary

- Functions can return custom functions using closure-based behavior
- Helps modularize logic and simplify repetitive tasks
- Closure = powerful feature for encapsulating configuration state

Next: Higher-order functions and practical physics routines

---

# Variadic and Deferred

## Variadic Functions

### Variable Argument Number

- In Go, functions can accept a variable number of arguments using the `...` (ellipsis) syntax.
- This is useful when you don’t know in advance how many values will be passed.
- Inside the function, the variadic parameter is treated as a slice of the specified type.

```go
func getMax(vals ...int) int {
    maxV := vals[0]
    for _, v := range vals {
        if v > maxV {
            maxV = v
        }
    }
    return maxV
}
```

### Passing a Slice to a Variadic Function

- You can also pass a slice directly to a variadic function.
- You must use the `...` after the slice variable to indicate you’re passing multiple values, not a single slice.

```go
func main() {
    fmt.Println(getMax(1, 3, 6, 4))        // pass individual integers
    vslice := []int{1, 3, 6, 4}
    fmt.Println(getMax(vslice...))        // pass a slice using ...
}
```

### Summary

- Use `...Type` to define variadic parameters.
- Variadic parameters are treated like slices within the function.
- You can either:
  - Pass individual arguments: `getMax(1, 2, 3)`
  - Pass a slice using `...`: `getMax(mySlice...)`

---

## Deferred Function Calls

### What Is Defer?

- The `defer` keyword postpones a function call until the surrounding function (typically `main` or any other function) finishes.
- Deferred calls are commonly used for cleanup tasks like:
  - Closing files
  - Unlocking mutexes
  - Logging exit messages

```go
func main() {
    defer fmt.Println("I am deferred")  // prints after everything else in main
    fmt.Println("Hello")
}
```

**Output:**

```
Hello
I am deferred
```

### Argument Evaluation in Defer

- **Important**: Function arguments in a deferred call are evaluated immediately at the point where `defer` is encountered—not when the deferred function is actually run.

```go
func main() {
    i := 1
    defer fmt.Println(i + 1) // i+1 is evaluated now, which is 2
    i++                      // i becomes 2
    fmt.Println("Hello")
}
```

**Output:**

```
Hello
2
```

- Even though `i` is incremented later, the deferred `fmt.Println(i + 1)` still prints `2` because `i+1` was evaluated when the `defer` statement was encountered.

---

### When to Use Defer

- Use `defer` when cleanup or final operations must always run regardless of return paths or errors.
- Common uses:
  - `defer file.Close()`
  - `defer unlock()`
  - `defer fmt.Println("Done")`

---
