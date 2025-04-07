# Module 1: Getting Started with Go

---

We begin by exploring:

- What makes **Go unique**?
- Why is Go worth learning over other programming languages?
- What are the benefits of using Go in real-world software development?

Next, we’ll setup go:

- **Install the Go environment** on our system.
- **Compile our first Go program** to ensure everything is working.

In this section, we’ll discuss:

- How to set up our **workspace**.
- Recommended **code organization** practices.
- How Go uses **packages** to structure and share code effectively.

> Sharing and collaborating are key goals of Go. Since most real software is developed in teams, understanding packages and modularity is essential.

Finally, we’ll begin working with the Go language directly:

- Learn about **variables** and **data types**.
- Understand **scoping rules**.
- See how Go resolves variable values based on **where they are defined and used**.

---

## Why Go Stands Out

> Go is not just fast or easy—it’s **engineered for productivity and scalability**. Let’s look at some design features that set Go apart from traditional languages.

### Built-in Concurrency

- Go provides **concurrency primitives** like `goroutines` and `channels`. These allow developers to:
  - Run multiple tasks **in parallel**.
  - Coordinate tasks **safely and efficiently**.
- Unlike other languages that rely on external libraries or OS-level threads, **Go's concurrency is part of the language itself**.
- **Ideal for modern applications** such as: Web servers, Real-time systems, Microservices, Cloud-native applications.

### Simplified Object-Oriented Programming

- Go supports a **lightweight form of object orientation**. No need for: Deep inheritance trees, Overloaded operators and Complex class hierarchies.
- Instead, Go uses: **Interfaces** for abstraction and **Structs with methods** for encapsulation.
- This keeps code **clean, readable, and maintainable**.

### Developer-Centric Language

- Go was designed by experienced engineers at Google who wanted to fix frustrations they had with other languages: Long compile times, Complex dependency management, Difficult concurrency models, Bloated syntax.
- **Go’s response:** Fast compile times, Simple syntax, Built-in tools like `go fmt`, `go run`, `go mod`, Excellent documentation practices.

## Conclusion: Why we Should Learn Go

| Feature              | Why It Matters                                             |
| -------------------- | ---------------------------------------------------------- |
| Compiled + Fast      | Execution speed comparable to C/C++                        |
| Garbage Collected    | Safer and easier memory management                         |
| Concurrency Built-in | Efficient handling of parallel tasks                       |
| Simple & Readable    | Easier to learn and maintain code                          |
| Modern & Scalable    | Perfect for microservices, APIs, and scalable architecture |

> **Bottom line**: Go combines the **speed of low-level languages** with the **developer-friendly features of high-level interpreted languages**—a perfect balance for modern software engineering.

---

## Object-Oriented Programming in Go

> Object-Oriented Programming (OOP) is a method of structuring code to improve readability, modularity, and reusability. It’s based on the concept of **encapsulation**, grouping together related data and the functions that operate on that data.

### Key Concepts of OOP

- **Encapsulation**: Organize code by bundling data and associated functions.
- **User-defined types**: Create types tailored to specific applications. Example: An `int` has both **data** (the value) and **functions** (e.g., `+`, `-`, `*`, `/`, `%`).

### Object Example – Geometry in 3D

Suppose you're building an application that performs 3D geometric calculations. You'll likely work with **points**.

#### Each `Point` has

- **Data**:
  - `x`, `y`, and `z` coordinates
- **Functions**:
  - `DisToOrigin()`: Calculates the distance from the origin
  - `Quadrant()`: Determines the spatial quadrant

### In Traditional OOP

- A **`class`** defines:
  - The structure (data fields)
  - The behavior (methods)
- **`Objects`** are instances of a class with actual values

## Objects in Go

Go supports OOP concepts but with a simpler approach:

- **No `class` keyword**
- Uses **`struct`** to define custom data types
- Methods can be associated with structs to define behavior

### Simplified OOP Model

| Feature      | Go Implementation         |
| ------------ | ------------------------- |
| Inheritance  | ❌ Not supported          |
| Constructors | ❌ Not supported          |
| Generics     | ❌ Not supported          |
| Methods      | ✅ Supported with structs |

> This makes Go **easy to learn, code, and execute efficiently**—ideal for developers who prefer simplicity over heavy OOP frameworks.

---

## Concurrency

> Modern computing demands handling many tasks efficiently and simultaneously. This section explores the **limits of traditional performance scaling**, the role of **parallelism**, and how **Go simplifies concurrent programming**.

### Performance Limits

- **Moore's Law**: Number of transistors used to double every 18 months _(now outdated)_.
- More transistors once meant: Higher **clock speeds**, Better performance automatically
- **Current challenge**: **Power and heat constraints** prevent further clock speed increases and Chips generate too much heat at high frequencies—limits are reached.

### Parallelism

- While clock speeds have plateaued, the number of **cores per processor is increasing**.
- Multiple cores = possibility of running **multiple tasks simultaneously**.

#### Challenges in Parallel Programming

- When should **tasks start or stop**?
- How should **tasks communicate**, especially if one depends on another?
- Do tasks **conflict in memory**?

### Concurrent Programming

- **Concurrency** = Managing multiple tasks that are alive at the same time. They may not run at the same time, but must be **handled together**. Essential for **large, complex systems** where many things happen "at once". **Concurrency enables parallelism**, but:
  - Tasks must be designed to run in parallel.
  - Hardware needs to be available (e.g., multiple cores).

#### Concurrency Involves

- Task execution management
- Communication between tasks
- Synchronization (e.g., one task must finish before another starts)

### Concurrency in Go

Go was designed from the ground up for concurrency:

| Feature      | Purpose                                   |
| ------------ | ----------------------------------------- |
| `Goroutines` | Lightweight concurrent threads            |
| `Channels`   | Safe communication between goroutines     |
| `Select`     | Coordination and synchronization of tasks |

> **Go makes concurrent programming efficient, readable, and scalable**—ideal for today’s multicore processors.

---

## Installing Go (Golang)

- Go to [https://golang.org](https://golang.org)
- Locate the button: **Download Go** and click it.

Once you click **Download Go**, you’ll land on a download page with precompiled versions for:

- ✅ **Windows** – Download the `.msi` installer
- ✅ **macOS** – Download the `.pkg` file
- ✅ **Linux** – Follow the tarball installation instructions

After installation:

- Your system’s environment will be updated (e.g., `GOPATH`, `GOROOT`, etc.)
- You can start writing and compiling Go programs locally.
- You can verify the installation by opening a terminal and typing:

```bash
go version
```

If installed successfully, it should return the installed version number.

---

## Go Tools

### Import

- The **_import_** keyword is used to access other packages.
- The Go standard library includes many built-in packages. Example: **_"fmt"_**
- When importing, Go searches the directories defined by: **_GOROOT_** and **_GOPATH_**

### Go Tool Overview

- The Go Tool is a command-line utility used to **build, run, format, test, and manage Go source code**.

### Common Go Tool Commands

- **_go build_** – Compiles the program.

  - Can accept a list of packages or `.go` files.
  - Creates an executable for the main package, named after the first `.go` file.
  - On Windows, the executable gets a `.exe` suffix.

- **_go doc_** – Prints documentation for a package.

- **_go fmt_** – Formats Go source code according to Go standards.

- **_go get_** – Downloads and installs external packages.

- **_go list_** – Lists all installed packages.

- **_go run_** – Compiles and runs `.go` files in one step.

- **_go test_** – Runs tests defined in files ending with `_test.go`.

---

## Variables

### Naming Rules

To refer to values in a program, you need **names**—for variables, functions, etc.

- Must start with a **letter**
- Can include **letters, digits, and underscores**
- **Case-sensitive** (e.g., `myVar` ≠ `MyVar`)
- Cannot use **Go reserved keywords** like: **_if_**, **_case_**, **_package_**, etc.

### Variable Basics

- A variable stores **data in memory**
- Every variable has: A **name** and A **type**
- All variables must be **declared before use**
- You can declare multiple variables of the same type on the same line

#### Basic Declaration

```golang
var x int // keyword name type
```

#### Declaring Multiple Variables

```golang
var x, y int
```

### Variable Types

A **type** defines: What **values** a variable can store and What **operations** can be performed on it.

#### Integer

- Whole numbers only
- Supports arithmetic: `+`, `-`, `*`, `/`, `%`

#### Floating Point

- Fractional (decimal) values
- Operations like addition, subtraction, division
- May be handled by **special hardware**

#### Strings

- A sequence of **Unicode bytes**
- Operations include: Comparison, Search, Concatenation.

> The compiler uses the variable's type to: Allocate the correct amount of memory and Select the proper machine-level instructions during compilation

### Type Declarations

- Every variable in Go must have a type.
- You can define **type aliases** to improve code clarity.

#### Examples

- `type Celsius float64` – Useful for temperature-related data.
- `type IDnum int` – Useful for labeling things like user ID numbers.

> Using aliases makes your code **semantically meaningful**—you understand the purpose of a variable from its type name.

### Initializing Variables

During Declaration: You can declare and initialize a variable in one line

```go
var x int = 100       //Explicit type
var x = 100          // Type is inferred as `int` from the value
```

> ⚠️ Be cautious when omitting the type. For example, `var x = 100` makes `x` an `int`. If you later assign `x = 100.1`, it will cause a type mismatch.

After Declaration: Declare first, initialize later

```go
var x int
x = 100
```

### Zero Values

If you declare a variable without initializing it, Go assigns a **default zero value**:

- `int` → 0
- `string` → `""` (empty string)

This ensures all variables are always in a valid state, even before assignment.

### Short Variable Declarations

Use `:=` to declare and initialize in one step:

```go
x := 100
```

This automatically infers the type of `x` from the right-hand side (`int` in this case).

> ✅ Can only be used **inside functions**. Using `:=` outside a function is not allowed.

---
