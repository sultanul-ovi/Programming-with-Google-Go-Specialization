# Welcome to Module 2: Go Data Types

Welcome to the second module! In this module, we'll start exploring the **Go programming language** in more detail—specifically focusing on **basic data types**.

## What You'll Learn

In this module, we will cover:

- An introduction to Go data types
- How to **declare and create** variables of different types
- How to **use built-in functions** associated with each type

## Topics Covered

You'll learn about the following **common data types** in Go:

- **Integers**
- **Floats**
- **Booleans**
- **Strings**

We will also explore **variations** of these types, such as:

- Changing the **bit-length** of integers and floats
- Working with **different formats** of strings and booleans

## Practical Skills

By the end of this module, you’ll be able to:

✅ Declare and use basic data types in Go  
✅ Apply functions to manipulate and interact with integers, strings, booleans, and floats  
✅ Understand how Go handles data types compared to other programming languages

---

Let's dive in and get comfortable using data types in Go!

## Go Basics: Understanding Pointers

## 🧠 What Are Pointers?

Pointers are variables that **store the memory address** of another variable. Every variable in Go resides somewhere in memory, and pointers help you work with that memory directly.

### 📍 Two Important Operators

- `&` (Ampersand): **Returns the memory address** of a variable
- `*` (Asterisk): **Dereferences a pointer**, i.e., gets the value stored at the memory address

### 📌 Example1

```go
var x int = 1
var y int
var ip *int // ip is a pointer to an int

ip = &x      // ip now stores the address of x
y = *ip      // y is set to the value at the address stored in ip (which is 1)
```

In this example:

- `ip` holds the address of `x`
- `*ip` retrieves the value of `x`, and assigns it to `y`

## 🔧 The `new` Function

Go provides the `new()` function to create variables and return **pointers** to them.

### 📌 Example2

```go
ptr := new(int) // ptr is a pointer to a new int initialized to 0
*ptr = 3        // dereference ptr and assign the value 3
```

- `new(int)` allocates memory for an int and returns a pointer to it
- `*ptr = 3` sets the allocated int's value to 3

## 🎯 Summary

- A **pointer** is just an address pointing to data in memory
- `&` gives you the address, `*` gives you the data
- `new()` creates a new variable and returns a pointer to it, defaulting the value to zero

By mastering pointers, you gain better control over memory and variable management in Go!

## Go Basics: Variable Scope

## 🔍 What is Variable Scope?

**Scope** refers to the regions of the code where a variable is **accessible**. It determines how variable names are resolved during compilation and runtime.

---

## 🧪 Example: Global vs Local Scope

### Version 1: Global `x`

```go
var x = 4

func f(){
  fmt.Printf("%d", x)
}

func g(){
  fmt.Printf("%d", x)
}
```

- `x` is declared outside the functions.
- Both `f()` and `g()` print **4**.
- `x` has **global (file) scope**.

### Version 2: Local `x`

```go
func f(){
  var x = 4
  fmt.Printf("%d", x)
}

func g(){
  fmt.Printf("%d", x) // Error: x is undefined here
}
```

- `x` is local to `f()`.
- `g()` cannot access `x`, causing a **compilation error**.

---

## 🧱 Blocks in Go

A **block** is a sequence of declarations and statements inside curly braces `{}`.

### Explicit Blocks

- Defined with `{}` in code, e.g., function bodies, loops

### Implicit Blocks (Hierarchical)

- **Universe Block**: All Go source
- **Package Block**: All files in a package
- **File Block**: One file’s source
- **Function Block**: Each function
- **Conditional Blocks**: `if`, `for`, `switch`, `select`, etc.
- **Case/Clause Blocks**: Each clause in `switch`/`select`

Blocks can be **nested**, and variables declared in outer blocks are accessible to inner blocks.

---

## 🔒 Lexical Scoping in Go

Go uses **lexical (static) scoping** based on block structure:

- A variable is **accessible** in a block `bj` **if it is declared in** `bj` **or any containing block** `bi`
- If a variable is not found in the current block, Go looks in the **outer block** until it finds it

### Block Relationship Notation

- If `bj` is nested inside `bi`, then `bi >= bj`
- Variable reference is resolved by scanning **upward** in the block hierarchy

### Scope Resolution Example

```go
// File Block (b1)
var x = 4

func f() { // Block b2
  fmt.Println(x) // Resolved in b1
}

func g() { // Block b3
  fmt.Println(x) // Resolved in b1
}
```

- Both `f` and `g` refer to `x` declared in block `b1`
- Since `b2` and `b3` are inside `b1`, they have access

```go
func f() { // Block b2
  var x = 4
  fmt.Println(x) // Resolved in b2
}

func g() { // Block b3
  fmt.Println(x) // ❌ Error: x not found in b3 or any enclosing block
}
```

---

## 📌 Final Notes

- Be careful where you declare your variables
- Use outer blocks for shared variables
- Prefer smaller scopes when possible to avoid accidental modification or conflicts

> 🛠️ **Correction from Lecture M2.1.2**: The instructor said `x = 1`, but it should be `x = 4` — the screen text is correct.

Understanding scope ensures better **code organization**, **error prevention**, and **logical clarity**.

## Go Memory Management: Deallocating Space

## 🧠 Why Deallocation Matters

When a variable is no longer needed, its **memory must be deallocated** to free up space. Failing to do this results in **memory leaks**, eventually consuming all available memory.

```go
func f() {
  var x = 4
  fmt.Printf("%d", x)
}
```

Each call to `f()` creates a **new instance of `x`**, which must be deallocated when the function ends.

---

## 🗃️ Stack vs. Heap

Go manages memory using two main areas: **stack** and **heap**.

### Version 1: Global Variable (Heap)

```go
var x = 4 // allocated on heap

func f() {
  fmt.Printf("%d", x)
}

func g() {
  fmt.Printf("%d", x)
}
```

### Version 2: Local Variable (Stack)

```go
func f() {
  var x = 4 // allocated on stack
  fmt.Printf("%d", x)
}

func g() {
  fmt.Printf("%d", x) // ❌ Error: x is undefined here
}
```

### Key Differences

- 🧱 **Stack**:

  - Used for **function calls**
  - Stores **local variables**
  - Automatically deallocated when function completes

- 🧠 **Heap**:
  - Used for **persistent data**
  - Variables must be **explicitly deallocated** in many languages
  - Memory does **not disappear** when the allocating function ends

---

## 🧹 Manual Deallocation (In C)

```c
x = malloc(32);  // allocate 32 bytes
free(x);         // deallocate memory
```

- Used in compiled languages like C
- **Fast**, but **error-prone**
- Forgetting to free memory = **memory leak**

---

## 🧠 How Go Handles This

Go takes a **hybrid approach**:

- Local variables may be placed on **stack** and deallocated automatically
- Other variables may be **heap-allocated**, and Go uses **garbage collection** to handle deallocation

### Debug Tip:

Use this command to check variable allocation:

```bash
go build -o demo.exe -gcflags -m
```

This shows whether a variable was allocated on the **stack** or **heap** by the Go compiler.

---

## 📌 Final Thoughts

- Memory must be deallocated when no longer needed
- Stack = automatic deallocation; Heap = manual (in some languages)
- Go simplifies this with **garbage collection**, but understanding memory layout helps with performance and debugging

## Go Memory Management: Garbage Collection

## 🗑️ What is Garbage Collection?

**Garbage collection (GC)** is the process of automatically identifying and **deallocating memory** that is no longer in use.

In programming, it’s often hard to determine when a variable is truly unused. Especially when pointers are involved, deallocation becomes complex and error-prone.

---

## 📌 Problem with Manual Deallocation

```go
func foo() *int {
  x := 1
  return &x
}

func main() {
  var y *int
  y = foo()
  fmt.Printf("%d", *y)
}
```

- `foo()` returns a **pointer to a local variable `x`**.
- Even after `foo()` ends, `main()` still uses the pointer.
- Deallocating `x` when `foo()` finishes would cause undefined behavior.

---

## ⚖️ Manual vs Automatic Deallocation

### Interpreted Languages (e.g., Java, Python)

- Garbage collection is handled by the **interpreter** (e.g., JVM, Python Interpreter)
- ✅ Easy for the programmer
- ❌ Slower due to interpretation overhead

### Compiled Languages (e.g., C)

- No garbage collection
- ✅ Fast
- ❌ Programmer must **manually** allocate and deallocate memory (e.g., `malloc()` and `free()`)

---

## 🚀 Go's Advantage: Compiled + Garbage Collection

Go is a **compiled language** that includes built-in **garbage collection**:

- ✅ No need to manually manage memory
- ✅ Faster than interpreted languages
- ✅ Compiler decides **stack vs heap** allocation
- ✅ Garbage collection happens in the background

### How Go GC Works

- Keeps track of all **references (pointers)** to a variable
- When **no references remain**, memory is safely deallocated
- Automatically differentiates between stack and heap at **compile-time**

> ⚙️ Run this to see variable memory allocation:

```bash
go build -gcflags -m
```

---

## ⚠️ Trade-off

- Garbage collection incurs **some performance overhead**
- But the convenience and safety it brings are usually worth it

---

## ✅ Summary

- Deallocating memory is essential, but tricky when pointers are involved
- Go simplifies this with a **built-in garbage collector**
- It combines the performance of a compiled language with the ease of memory management from interpreted languages

Garbage collection is one of Go’s **most helpful features**, making it easier and safer to write efficient code.

## Go Basics: Comments, Printing, and Integers

## 💬 Comments

Comments improve code readability and are **ignored by the compiler**.

### Single-line Comments

```go
// This is a single-line comment
var x int // Comment after a statement
```

### Block Comments

```go
/*
  This is a block comment
  that spans multiple lines
*/
var x int
```

---

## 🖨️ Printing in Go

Printing is handled using the `fmt` package.

### Import the Package

```go
import "fmt"
```

### Basic Printing

```go
fmt.Println("Hello, World!")
```

### Formatted Output

Use `fmt.Printf()` with **format strings** and **conversion characters**.

```go
name := "Joe"
fmt.Printf("Hi %s\n", name) // Output: Hi Joe
```

- `%s`: string
- `%d`: integer
- `%f`: float

Format strings let you insert values neatly into output using `%` codes.

---

## 🔢 Integers in Go

### Declaration

```go
var x int // Generic integer
```

### Integer Types

- Signed:
  - `int8`, `int16`, `int32`, `int64`
- Unsigned:
  - `uint8`, `uint16`, `uint32`, `uint64`

> The number indicates **bit length**; `u` = unsigned (no negative values).

### Integer Ranges

- `int8`: -128 to 127
- `uint8`: 0 to 255
- `int16`: -32,768 to 32,767
- `uint16`: 0 to 65,535

### Operators

- **Arithmetic**: `+`, `-`, `*`, `/`, `%`
- **Comparison**: `==`, `!=`, `<`, `>`, `<=`, `>=`
- **Boolean**: `&&`, `||`
- **Bitwise**: `&`, `|`, `^`, `<<`, `>>`

---

- Use `//` and `/* */` for adding helpful comments
- Print output using `fmt.Printf()` and `fmt.Println()`
- Choose integer types based on size and sign requirements
- Operators allow arithmetic, comparison, and logic operations on integers

---

# Go Basics: Ints, Floats, and Strings

## 🔄 Type Conversions

- Binary operations require operands to be **of the same type**, including assignment

### Example (❌ Invalid Assignment)

```go
var x int32 = 1
var y int16 = 2
x = y // ERROR: mismatched types
```

### ✅ Type Conversion Syntax

```go
x = int32(y) // Correct: convert y to int32
```

---

## 🌊 Floating Point Types

- **float32**: ~6 decimal digits of precision
- **float64**: ~15 decimal digits of precision

### Examples

```go
var x float64 = 123.45
var y float64 = 1.234e2 // scientific notation = 123.4
```

### Complex Numbers

```go
var z complex128 = complex(2, 3) // 2 + 3i
```

---

## 🔡 ASCII and Unicode

- **ASCII**: 7/8-bit character code (e.g., `'A' = 0x41`)
- **Unicode**: 32-bit character set supporting global languages
- **UTF-8**: Variable-length encoding of Unicode
  - First 128 UTF-8 codes = ASCII
  - Allows encoding characters using 8, 16, or 32 bits

### Definitions

- **Code Point**: A Unicode character
- **Rune**: Go's term for a Unicode code point

🔗 [Learn more about Unicode](http://unicode.org/faq/utf_bom.html)

---

## 📦 Strings in Go

- A **string** is a **read-only** sequence of **UTF-8 encoded bytes**
- **Immutable**: cannot be changed in-place
- Often used for display/output

### Example

```go
x := "Hi there"
```

- Each character is stored as a **rune**
- Enclosed in double quotes (`"`)

---

## ✅ Summary

- Convert types explicitly using the `T(value)` format
- Use `float32` or `float64` based on precision needs
- UTF-8 strings allow representation of global characters
- Strings in Go are immutable and stored as sequences of UTF-8 runes

---

# Go Standard Library: String Packages

## ✨ Unicode Package

Go's `unicode` package provides functions to analyze and manipulate **runes** (Unicode characters).

### 🔍 Rune Classification

All functions return `true` or `false` based on the category of the rune:

```go
unicode.IsDigit(r)
unicode.IsSpace(r)
unicode.IsLetter(r)
unicode.IsLower(r)
unicode.IsPunct(r)
```

### 🔄 Rune Conversion

```go
unicode.ToUpper(r)
unicode.ToLower(r)
```

Convert a rune to uppercase or lowercase respectively.

---

## 🧰 Strings Package

The `strings` package offers tools to **search** and **manipulate** entire strings.

### 🔎 Search Functions

```go
strings.Compare(a, b)     // 0 if a == b, -1 if a < b, 1 if a > b
strings.Contains(s, sub)  // true if sub is in s
strings.HasPrefix(s, pre) // true if s starts with pre
strings.Index(s, sub)     // index of first sub in s, or -1
```

### 🛠️ String Manipulation

All functions return a **new string**, as strings are immutable:

```go
strings.Replace(s, old, new, n) // replaces up to n instances of old with new
strings.ToLower(s)              // lowercase version of s
strings.ToUpper(s)              // uppercase version of s
strings.TrimSpace(s)           // removes leading/trailing white space
```

---

## 🔁 Strconv Package

The `strconv` package is for **converting between strings and basic data types**.

### 🔄 Conversion Functions

```go
strconv.Atoi(s)                      // string to int
strconv.Itoa(i)                      // int to string
strconv.FormatFloat(f, fmt, prec, bitSize) // float to string
strconv.ParseFloat(s, bitSize)       // string to float
```

> ⚠️ Lecture Note: In Lecture 2.2.3, at 5:14, the screen says `itoa(s)`, but it should be `itoa(i)` — `i` is the int being converted.

---

## ✅ Summary

- Use `unicode` for character-level analysis and conversions
- Use `strings` for full string search and manipulation
- Use `strconv` for converting between strings and numeric types

These packages make Go powerful for handling **text, encoding, and user input** operations efficiently.

---

# Go Basics: Constants and `iota`

## 🧱 Constants

A **constant** is an expression whose value is known at **compile time** and cannot be changed.

### Features

- Type is inferred from the **right-hand side**
- Can declare **multiple constants** in a block

### Example

```go
const x = 1.3

const (
  y = 4
  z = "Hi"
)
```

---

## 🔁 `iota`

`iota` is a special identifier used in constant declarations to generate a set of related but **distinct values**.

### Use Cases

Ideal for representing **enumerated types**:

- Days of the week
- Months of the year
- Grades, states, modes, etc.

You don’t care about the **specific values**, just that each constant is **different**.

### Example: Enum for Grades

```go
type Grades int

const (
  A Grades = iota
  B
  C
  D
  F
)
```

- Starts from `0` by default and **increments** by `1`
- Equivalent to:
  ```go
  A = 0, B = 1, C = 2, D = 3, F = 4
  ```

> 🧠 **Note**: Although `iota` starts at `0` and increments, you **should not rely** on the actual numeric values.

---

## ✅ Summary

- Use `const` to define unchanging values known at compile time
- `iota` is great for generating distinct constants automatically
- Enables clean, readable enum-like declarations

`iota` gives you a **shorthand** for enumerations without worrying about assigning explicit values.

---

# Go Basics: Control Flow

## 🔁 What is Control Flow?

Control flow describes the **order in which statements are executed** inside a program.

- Default flow is **top-down** (one line at a time)
- Can be altered using **control structures**

---

## ⚖️ `if` Statements

Use `if` to conditionally execute code blocks based on a boolean condition.

### Syntax

```go
if condition {
  // code executes only if condition is true
}
```

### With `else`

```go
if condition {
  // executes if true
} else {
  // executes if false
}
```

---

## 🔁 `for` Loops

The only loop keyword in Go is `for`, but it supports multiple loop forms.

### 1. Classic Loop (like in C)

```go
for i := 0; i < 10; i++ {
  fmt.Println(i)
}
```

- `init`: runs once
- `condition`: checked before each iteration
- `update`: runs after each iteration

### 2. Condition-only Loop (like `while`)

```go
i := 0
for i < 10 {
  fmt.Println(i)
  i++
}
```

### 3. Infinite Loop

```go
for {
  // runs forever unless you break out
}
```

> ⚠️ Make sure the loop condition eventually becomes false to avoid infinite loops (unless intentional).

---

## 🔄 `switch` Statement

A `switch` provides a cleaner way to handle multiple conditional branches.

### Syntax

```go
switch x {
case 1:
  fmt.Println("One")
case 2:
  fmt.Println("Two")
default:
  fmt.Println("Default")
}
```

### Key Notes

- Compares the value of `x` to each case
- **Only one case is executed**
- **No need for `break`** like in C — it auto-breaks
- `default` is optional, runs if no case matches

---

## ✅ Summary

- Use `if` and `if-else` for conditional branches
- Use `for` for looping (supports C-style, condition-only, infinite)
- Use `switch` for multi-way conditional branching

Control flow is essential for writing dynamic and responsive Go programs!

---

# Go Control Flow and Input

## 🔁 Control Flow

Control flow describes the **order in which statements execute** in a program.

- Default flow is **top-down**, one statement after another.
- Programmers insert **control structures** to modify this flow.

### 🧩 Control Structures

- `if`, `if-else` statements
- `for` loops
- `switch/case` statements

---

## 🔄 Tagless Switch

- A **switch** does not need a tag.
- Each `case` contains a **boolean condition**.
- The **first case** that evaluates to `true` is executed.

### Example

```go
switch {
  case x > 1:
    fmt.Println("x is greater than 1")
  case x < -1:
    fmt.Println("x is less than -1")
  default:
    fmt.Println("x is in range")
}
```

---

## ⛔ Break and 🔁 Continue

### Break

- Immediately **exits** the current loop.

```go
for i := 0; i < 10; i++ {
  if i == 5 {
    break
  }
  fmt.Println(i)
}
```

### Continue

- **Skips** the current iteration of the loop.

```go
for i := 0; i < 10; i++ {
  if i == 5 {
    continue
  }
  fmt.Println(i)
}
```

---

## 📥 `fmt.Scan` — Reading User Input

- Reads user input from the terminal
- Takes a **pointer** as an argument
- Blocks until **Enter** is pressed

### Example

```go
var appleNum int
fmt.Print("Number of apples: ")
fmt.Scan(&appleNum)
fmt.Printf("You entered: %d\n", appleNum)
```

### Details

- Stores the value in the location pointed to by the argument
- Returns two values:
  - Number of scanned items
  - An `error` value (`nil` if successful)

---

## ✅ Summary

- Go supports structured control flow via `if`, `for`, and `switch`
- Tagless `switch` provides a clean `if-else` alternative
- Use `break` and `continue` to fine-tune loop execution
- Use `fmt.Scan` to gather user input with pointer-based assignments

---
