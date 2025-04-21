# 📝 Task List

1. **(🟢 Easy)** Write a function `SayHello()` that prints "Hello from function!" and call it from `main()`.
2. **(🟢 Easy)** Write a function `Add(x int, y int) int` that returns the sum of two integers.
3. **(🟢 Easy)** Write a function `Double(x int) int` and use it to double a number.
4. **(🟡 Medium)** Write a function `Swap(x int, y int) (int, int)` that returns the two values in reverse order.
5. **(🟡 Medium)** Write a function `Increment(ptr *int)` that increments the value at a given memory address.
6. **(🟡 Medium)** Write a function `UpdateFirst(arr []int)` that updates the first element of a slice to `100`.
7. **(🔴 Hard)** Create a struct `Point{x, y float64}` and a function `Distance(p1, p2 Point) float64` to compute Euclidean distance.
8. **(🔴 Hard)** Create a struct `Triangle{p1, p2, p3 Point}` and a function `TriangleArea(t Triangle) float64` to compute its area.
9. **(🔴 Hard)** Write a function `ProcessSlice(data []int) []int` that doubles only even values and returns a new slice.
10. **(🟡 Medium)** Demonstrate call by value vs call by reference using a function `Modify(x int)` and `ModifyPtr(x *int)`.
11. **(🟢 Easy)** Assign a named function that adds 1 to an integer to a variable and call it.
12. **(🟢 Easy)** Write a function `applyFunc(f func(int) int, x int) int` that takes a function and a value, applies the function to the value, and returns the result.
13. **(🟡 Medium)** Use anonymous functions as inline arguments to apply a square operation to a number.
14. **(🟡 Medium)** Create a function `MakeMultiplier(n int)` that returns a function which multiplies its input by `n`. Demonstrate the closure property.
15. **(🟡 Medium)** Use a closure to create a custom greeter function with a preconfigured name and call it.
16. **(🔴 Hard)** Write a function `MakeDistOrigin(o_x, o_y float64) func(float64, float64) float64` that returns a custom distance calculator using closures.
17. **(🔴 Hard)** Create two different distance functions using `MakeDistOrigin` and show that they calculate distances relative to different origins.
18. **(🟡 Medium)** Demonstrate how the closure remembers the environment by printing internal values.
19. **(🟢 Easy)** Create a variadic function `SumAll(nums ...int) int` that returns the sum of all input integers.
20. **(🟡 Medium)** Write a variadic function `MaxAll(nums ...int) int` that finds the maximum of all values.
21. **(🟢 Easy)** Call a variadic function by passing a slice using the `...` syntax.
22. **(🟢 Easy)** Write a program that uses `defer` to print "Goodbye" after "Hello".
23. **(🟡 Medium)** Show that arguments in deferred functions are evaluated immediately by changing a variable after the defer call.
24. **(🟡 Medium)** Demonstrate using `defer` for cleanup by opening and closing a dummy resource.
25. **(🔴 Hard)** Write a function that returns a function and uses both variadic parameters and closures to configure a custom aggregator.
26. **(🟢 Easy)** Define a custom type `MyInt` and add a method `Double()` that returns twice its value.
27. **(🟢 Easy)** Create a struct `Point` with fields `x` and `y`. Add a method `DistToOrig()` that returns the distance from origin.
28. **(🟡 Medium)** Implement pointer receiver method `OffsetX()` for `Point` to increase `x` by a value. Call it from `main()`.
29. **(🟡 Medium)** Create methods `SetX()` and `SetY()` for `Point` using pointer receivers. Use them to update coordinates.
30. **(🔴 Hard)** Define a `Point` struct with private fields and public methods `InitMe`, `Scale`, and `PrintMe` from another package. Demonstrate usage from `main` package.
31. **(🔴 Hard)** Create a `Rectangle` struct with `length` and `width`. Add methods `Area()` and `Scale(factor float64)` using pointer receivers.
32. **(🟡 Medium)** Demonstrate automatic referencing and dereferencing by calling pointer receiver methods on value objects.
33. **(🟢 Easy)** Write a method `Quadrant()` on `Point` to determine which quadrant it lies in.
34. **(🔴 Hard)** Create a struct `Image` with a large array and show why pointer receivers are more efficient.
35. **(🟡 Medium)** Add a method `DoubleDist()` to `Point` that doubles both `x` and `y` to double the distance from origin.
36. **(🟢 Easy)** Define an interface `Speaker` with a method `Speak()` and implement it using a `Cat` and a `Dog` struct.
37. **(🟢 Easy)** Write a function `MakeItSpeak` that accepts a `Speaker` and calls its `Speak()` method.
38. **(🟡 Medium)** Create a `Shape` interface with an `Area()` method and implement it using `Rectangle` and `Triangle`.
39. **(🟡 Medium)** Implement both `Area()` and `Perimeter()` in a `Shape2D` interface, then define a `printShapeInfo` function.
40. **(🔴 Hard)** Write a Go routine that takes user input to create and query animals using interfaces.
41. **(🟡 Medium)** Define a function `FitInYard(s Shape2D) bool` that checks area and perimeter constraints.
42. **(🟢 Easy)** Demonstrate use of the empty interface `interface{}` to create a function `PrintAny(val interface{})`.
43. **(🔴 Hard)** Implement a `DrawShape` function using type assertions to call specific draw methods based on shape type.
44. **(🔴 Hard)** Rewrite `DrawShape` using a type switch to simplify the code and support multiple types.
45. **(🟢 Easy)** Create a struct `Dog` with field `name` and method `Speak()` to print its name.
46. **(🔴 Hard)** Show how calling a method on an interface with nil dynamic value is valid, but nil dynamic type is not.
47. **(🟡 Medium)** Define and use the built-in `error` interface to return and handle a custom error.
48. **(🟡 Medium)** Implement a `Divide(a, b int) (int, error)` function that handles divide-by-zero using error interface.
49. **(🟢 Easy)** Demonstrate how interface values hold both dynamic type and dynamic value.
50. **(🟡 Medium)** Create an interface `Notifier` with a method `Notify()` and show how a `User` struct implements it.

---

## ✅ Task Solutions

### ✅ Task 1: Write a function named `SayHello()` that prints "Hello from function!" to the console. Call this function from the `main()` function to display the message

```go
package main
import "fmt"

func SayHello() {
    fmt.Println("Hello from function!")
}

func main() {
    SayHello()
}
```

### ✅ Task 2: Write a function named `Add` that accepts two integer parameters `x` and `y`, and returns their sum as an integer. Call the function from `main()` and print the result

```go
package main
import "fmt"

func Add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println("Sum:", Add(3, 4))
}
```

### ✅ Task 3: Write a function named `Double` that accepts an integer and returns the result of multiplying that number by 2. Call it from `main()` and print the output

```go
package main
import "fmt"

func Double(x int) int {
    return x * 2
}

func main() {
    fmt.Println("Double of 5:", Double(5))
}
```

### ✅ Task 4: Write a function named `Swap` that takes two integers `x` and `y` and returns the two values in reverse order. In `main()`, call this function with sample values and print the result

```go
package main
import "fmt"

func Swap(x int, y int) (int, int) {
    return y, x
}

func main() {
    a, b := Swap(10, 20)
    fmt.Println("Swapped:", a, b)
}
```

### ✅ Task 5: Write a function named `Increment` that takes a pointer to an integer as a parameter. Increment the value pointed to by the pointer. Call this function from `main()` and print the updated value

```go
package main
import "fmt"

func Increment(ptr *int) {
    *ptr++
}

func main() {
    x := 5
    Increment(&x)
    fmt.Println("Incremented:", x)
}
```

### ✅ Task 6: Write a function named `UpdateFirst` that takes a slice of integers as input and modifies the first element in the slice to be `100`. Print the updated slice from the `main()` function

```go
package main
import "fmt"

func UpdateFirst(arr []int) {
    arr[0] = 100
}

func main() {
    nums := []int{1, 2, 3}
    UpdateFirst(nums)
    fmt.Println(nums)
}
```

### ✅ Task 7: Define a struct named `Point` with `x` and `y` float64 fields. Write a function `Distance(p1, p2 Point) float64` that computes the Euclidean distance between two points. Call this function from `main()` and print the result

```go
package main
import (
    "fmt"
    "math"
)

type Point struct {
    x, y float64
}

func Distance(p1, p2 Point) float64 {
    dx := p2.x - p1.x
    dy := p2.y - p1.y
    return math.Sqrt(dx*dx + dy*dy)
}

func main() {
    p1 := Point{0, 0}
    p2 := Point{3, 4}
    fmt.Println("Distance:", Distance(p1, p2))
}
```

### ✅ Task 8: Create a struct named `Triangle` composed of three `Point` fields: `p1`, `p2`, and `p3`. Write a function `TriangleArea(t Triangle) float64` that computes the area of the triangle using the determinant method. Print the result from `main()`

```go
package main
import "fmt"

type Point struct {
    x, y float64
}

type Triangle struct {
    p1, p2, p3 Point
}

func TriangleArea(t Triangle) float64 {
    return 0.5 * ((t.p1.x*(t.p2.y-t.p3.y) + t.p2.x*(t.p3.y-t.p1.y) + t.p3.x*(t.p1.y-t.p2.y)))
}

func main() {
    t := Triangle{Point{0, 0}, Point{4, 0}, Point{0, 3}}
    fmt.Println("Area:", TriangleArea(t))
}
```

### ✅ Task 9: Write a function `ProcessSlice(data []int) []int` that returns a new slice where all even numbers in the input slice are doubled, while odd numbers remain unchanged. Print the resulting slice

```go
package main
import "fmt"

func ProcessSlice(data []int) []int {
    result := make([]int, len(data))
    for i, v := range data {
        if v%2 == 0 {
            result[i] = v * 2
        } else {
            result[i] = v
        }
    }
    return result
}

func main() {
    input := []int{1, 2, 3, 4, 5}
    fmt.Println("Processed:", ProcessSlice(input))
}
```

### ✅ Task 10: Write two functions: `Modify(x int)` which attempts to modify its input value, and `ModifyPtr(x *int)` which modifies the value pointed to. Demonstrate from `main()` how passing by value and by reference behave differently

```go
package main
import "fmt"

func Modify(x int) {
    x = 10
}

func ModifyPtr(x *int) {
    *x = 10
}

func main() {
    a := 5
    b := 5
    Modify(a)
    ModifyPtr(&b)
    fmt.Println("a (by value):", a) // 5
    fmt.Println("b (by reference):", b) // 10
}
```

### ✅ Task 11: Assign a named function that adds 1 to an integer to a variable and call it

```go
package main
import "fmt"

func inc(x int) int {
    return x + 1
}

func main() {
    var funcVar func(int) int
    funcVar = inc
    fmt.Println(funcVar(5))
}
```

### ✅ Task 12: Write a function `applyFunc(f func(int) int, x int) int` that takes a function and a value, applies the function to the value, and returns the result

```go
package main
import "fmt"

func applyFunc(f func(int) int, x int) int {
    return f(x)
}

func double(n int) int {
    return n * 2
}

func main() {
    fmt.Println(applyFunc(double, 4))
}
```

### ✅ Task 13: Use anonymous functions as inline arguments to apply a square operation to a number

```go
package main
import "fmt"

func applyFunc(f func(int) int, x int) int {
    return f(x)
}

func main() {
    result := applyFunc(func(n int) int {
        return n * n
    }, 5)
    fmt.Println(result)
}
```

### ✅ Task 14: Create a function `MakeMultiplier(n int)` that returns a function which multiplies its input by `n`. Demonstrate the closure property

```go
package main
import "fmt"

func MakeMultiplier(n int) func(int) int {
    return func(x int) int {
        return x * n
    }
}

func main() {
    times3 := MakeMultiplier(3)
    fmt.Println(times3(4))
}
```

### ✅ Task 15: Use a closure to create a custom greeter function with a preconfigured name and call it

```go
package main
import "fmt"

func MakeGreeter(name string) func() {
    return func() {
        fmt.Println("Hello,", name)
    }
}

func main() {
    greetAlice := MakeGreeter("Alice")
    greetAlice()
}
```

### ✅ Task 16: Write a function `MakeDistOrigin(o_x, o_y float64) func(float64, float64) float64` that returns a custom distance calculator using closures

```go
package main
import (
    "fmt"
    "math"
)

func MakeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
    return func(x, y float64) float64 {
        return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
    }
}

func main() {
    distFromZero := MakeDistOrigin(0, 0)
    fmt.Println(distFromZero(3, 4))
}
```

### ✅ Task 17: Create two different distance functions using `MakeDistOrigin` and show that they calculate distances relative to different origins

```go
package main
import (
    "fmt"
    "math"
)

func MakeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
    return func(x, y float64) float64 {
        return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
    }
}

func main() {
    dist1 := MakeDistOrigin(0, 0)
    dist2 := MakeDistOrigin(2, 3)

    fmt.Println("From (0,0):", dist1(2, 2))
    fmt.Println("From (2,3):", dist2(2, 2))
}
```

### ✅ Task 18: Demonstrate how the closure remembers the environment by printing internal values

```go
package main
import "fmt"

func RememberValue(val int) func() {
    return func() {
        fmt.Println("Captured:", val)
    }
}

func main() {
    remember := RememberValue(42)
    remember()
}
```

### ✅ Task 19: Create a variadic function `SumAll(nums ...int) int` that returns the sum of all input integers

```go
package main
import "fmt"

func SumAll(nums ...int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    return sum
}

func main() {
    fmt.Println(SumAll(1, 2, 3, 4))
}
```

### ✅ Task 20: Write a variadic function `MaxAll(nums ...int) int` that finds the maximum of all values

```go
package main
import "fmt"

func MaxAll(nums ...int) int {
    max := nums[0]
    for _, v := range nums {
        if v > max {
            max = v
        }
    }
    return max
}

func main() {
    fmt.Println(MaxAll(1, 5, 3, 7, 2))
}
```

### ✅ Task 21: Call a variadic function by passing a slice using the `...` syntax

```go
package main
import "fmt"

func SumAll(nums ...int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    return sum
}

func main() {
    values := []int{2, 4, 6}
    fmt.Println(SumAll(values...))
}
```

### ✅ Task 22: Write a program that uses `defer` to print "Goodbye" after "Hello"

```go
package main
import "fmt"

func main() {
    defer fmt.Println("Goodbye")
    fmt.Println("Hello")
}
```

### ✅ Task 23: Show that arguments in deferred functions are evaluated immediately by changing a variable after the defer call

```go
package main
import "fmt"

func main() {
    x := 10
    defer fmt.Println("Deferred value:", x+1)
    x = 20
    fmt.Println("Current value:", x)
}
```

### ✅ Task 24: Demonstrate using `defer` for cleanup by opening and closing a dummy resource

```go
package main
import "fmt"

type Resource struct{}

func (r *Resource) Close() {
    fmt.Println("Resource closed")
}

func useResource() {
    r := &Resource{}
    defer r.Close()
    fmt.Println("Using resource")
}

func main() {
    useResource()
}
```

### ✅ Task 25: Write a function that returns a function and uses both variadic parameters and closures to configure a custom aggregator

```go
package main
import "fmt"

func MakeAggregator(multiplier int) func(...int) int {
    return func(nums ...int) int {
        sum := 0
        for _, v := range nums {
            sum += v
        }
        return sum * multiplier
    }
}

func main() {
    agg := MakeAggregator(2)
    fmt.Println(agg(1, 2, 3)) // (1+2+3)*2 = 12
}
```

### ✅ Task 26: Define a custom type `MyInt` and add a method `Double()` that returns twice its value

```go
package main
import "fmt"

type MyInt int

func (m MyInt) Double() int {
    return int(m * 2)
}

func main() {
    val := MyInt(5)
    fmt.Println(val.Double()) // Output: 10
}
```

### ✅ Task 27: Create a struct `Point` with fields `x` and `y`. Add a method `DistToOrig()` that returns the distance from origin

```go
package main
import (
    "fmt"
    "math"
)

type Point struct {
    x, y float64
}

func (p Point) DistToOrig() float64 {
    return math.Sqrt(p.x*p.x + p.y*p.y)
}

func main() {
    p := Point{3, 4}
    fmt.Println(p.DistToOrig()) // Output: 5
}
```

### ✅ Task 28: Implement pointer receiver method `OffsetX()` for `Point` to increase `x` by a value. Call it from `main()`

```go
package main
import "fmt"

type Point struct {
    x, y float64
}

func (p *Point) OffsetX(v float64) {
    p.x += v
}

func main() {
    p := Point{2, 3}
    p.OffsetX(5)
    fmt.Println(p.x) // Output: 7
}
```

### ✅ Task 29: Create methods `SetX()` and `SetY()` for `Point` using pointer receivers. Use them to update coordinates

```go
package main
import "fmt"

type Point struct {
    x, y float64
}

func (p *Point) SetX(x float64) {
    p.x = x
}

func (p *Point) SetY(y float64) {
    p.y = y
}

func main() {
    p := Point{1, 1}
    p.SetX(4)
    p.SetY(5)
    fmt.Println(p) // Output: {4 5}
}
```

### ✅ Task 30: Define a `Point` struct with private fields and public methods `InitMe`, `Scale`, and `PrintMe` from another package. Demonstrate usage from `main` package

```go
// data/point.go
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

```go
// main.go
package main

import (
    "data"
)

func main() {
    var p data.Point
    p.InitMe(3, 4)
    p.Scale(2)
    p.PrintMe() // Output: 6 8
}
```

### ✅ Task 31: Create a `Rectangle` struct with `length` and `width`. Add methods `Area()` and `Scale(factor float64)` using pointer receivers

```go
package main
import "fmt"

type Rectangle struct {
    length, width float64
}

func (r *Rectangle) Area() float64 {
    return r.length * r.width
}

func (r *Rectangle) Scale(f float64) {
    r.length *= f
    r.width *= f
}

func main() {
    rect := Rectangle{4, 5}
    fmt.Println("Area:", rect.Area())
    rect.Scale(2)
    fmt.Println("Scaled Area:", rect.Area()) // 80
}
```

### ✅ Task 32: Demonstrate automatic referencing and dereferencing by calling pointer receiver methods on value objects

```go
package main
import "fmt"

type Point struct {
    x float64
}

func (p *Point) OffsetX(v float64) {
    p.x += v
}

func main() {
    p := Point{3}
    p.OffsetX(5) // Automatically gets address
    fmt.Println(p.x) // 8
}
```

### ✅ Task 33: Write a method `Quadrant()` on `Point` to determine which quadrant it lies in

```go
package main
import "fmt"

type Point struct {
    x, y float64
}

func (p Point) Quadrant() string {
    switch {
    case p.x > 0 && p.y > 0:
        return "Quadrant I"
    case p.x < 0 && p.y > 0:
        return "Quadrant II"
    case p.x < 0 && p.y < 0:
        return "Quadrant III"
    case p.x > 0 && p.y < 0:
        return "Quadrant IV"
    default:
        return "On Axis"
    }
}

func main() {
    p := Point{-2, 3}
    fmt.Println(p.Quadrant()) // Quadrant II
}
```

### ✅ Task 34: Create a struct `Image` with a large array and show why pointer receivers are more efficient

```go
package main
import "fmt"

type Image struct {
    data [10000]int
}

func (img *Image) Clear() {
    for i := range img.data {
        img.data[i] = 0
    }
}

func main() {
    var img Image
    img.Clear()
    fmt.Println("Image cleared.")
}
```

### ✅ Task 35: Add a method `DoubleDist()` to `Point` that doubles both `x` and `y` to double the distance from origin

```go
package main
import "fmt"

type Point struct {
    x, y float64
}

func (p *Point) DoubleDist() {
    p.x *= 2
    p.y *= 2
}

func main() {
    p := Point{3, 4}
    p.DoubleDist()
    fmt.Println(p) // {6 8}
}
```

### ✅ Task 36: Define an interface `Speaker` with a method `Speak()` and implement it using a `Cat` and a `Dog` struct

```go
package main
import "fmt"

type Speaker interface {
    Speak()
}

type Cat struct{}
func (c Cat) Speak() { fmt.Println("meow") }

type Dog struct{}
func (d Dog) Speak() { fmt.Println("woof") }

func main() {
    var s Speaker
    s = Cat{}
    s.Speak()
    s = Dog{}
    s.Speak()
}
```

### ✅ Task 37: Write a function `MakeItSpeak` that accepts a `Speaker` and calls its `Speak()` method

```go
package main
import "fmt"

type Speaker interface {
    Speak()
}

type Dog struct{}
func (d Dog) Speak() { fmt.Println("woof") }

func MakeItSpeak(s Speaker) {
    s.Speak()
}

func main() {
    MakeItSpeak(Dog{})
}
```

### ✅ Task 38: Create a `Shape` interface with an `Area()` method and implement it using `Rectangle` and `Triangle`

```go
package main
import "fmt"

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

func main() {
    var s Shape
    s = Rectangle{4, 5}
    fmt.Println(s.Area())
    s = Triangle{4, 5}
    fmt.Println(s.Area())
}
```

### ✅ Task 39: Implement both `Area()` and `Perimeter()` in a `Shape2D` interface, then define a `printShapeInfo` function

```go
package main
import "fmt"

type Shape2D interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

func printShapeInfo(s Shape2D) {
    fmt.Println("Area:", s.Area())
    fmt.Println("Perimeter:", s.Perimeter())
}

func main() {
    r := Rectangle{10, 5}
    printShapeInfo(r)
}
```

### ✅ Task 40: Write a Go routine that takes user input to create and query animals using interfaces

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

    for {
        fmt.Print("> ")
        scanner.Scan()
        input := strings.Fields(scanner.Text())
        if len(input) == 0 { continue }
        switch input[0] {
        case "create":
            if len(input) == 3 {
                name, kind := input[1], input[2]
                switch kind {
                case "cat": animals[name] = Cat{}
                case "dog": animals[name] = Dog{}
                }
            }
        case "query":
            if len(input) == 2 {
                if a, ok := animals[input[1]]; ok {
                    fmt.Printf("%s says: ", a.Type())
                    a.Speak()
                }
            }
        case "exit":
            return
        }
    }
}
```

### ✅ Task 41: Define a function `FitInYard(s Shape2D) bool` that checks area and perimeter constraints

```go
func FitInYard(s Shape2D) bool {
    return s.Area() < 100 && s.Perimeter() < 100
}
```

### ✅ Task 42: Demonstrate use of the empty interface `interface{}` to create a function `PrintAny(val interface{})`

```go
package main
import "fmt"

func PrintAny(val interface{}) {
    fmt.Println(val)
}

func main() {
    PrintAny(42)
    PrintAny("hello")
    PrintAny(true)
}
```

### ✅ Task 43: Implement a `DrawShape` function using type assertions to call specific draw methods based on shape type

```go
func DrawShape(s Shape2D) {
    if r, ok := s.(Rectangle); ok {
        fmt.Println("Drawing Rectangle")
        _ = r
    }
    if t, ok := s.(Triangle); ok {
        fmt.Println("Drawing Triangle")
        _ = t
    }
}
```

### ✅ Task 44: Rewrite `DrawShape` using a type switch to simplify the code and support multiple types

```go
func DrawShape(s Shape2D) {
    switch sh := s.(type) {
    case Rectangle:
        fmt.Println("Drawing Rectangle")
    case Triangle:
        fmt.Println("Drawing Triangle")
    default:
        fmt.Println("Unknown shape")
    }
}
```

### ✅ Task 45: Create a struct `Dog` with field `name` and method `Speak()` to print its name

```go
package main
import "fmt"

type Dog struct {
    name string
}

func (d Dog) Speak() {
    fmt.Println(d.name)
}

func main() {
    d := Dog{"Buddy"}
    d.Speak()
}
```

### ✅ Task 46: Show how calling a method on an interface with nil dynamic value is valid, but nil dynamic type is not

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
    var s Speaker
    var d *Dog // nil
    s = d
    s.Speak() // Valid: prints <noise>

    var s2 Speaker
    s2.Speak() // RUNTIME ERROR
}
```

### ✅ Task 47: Define and use the built-in `error` interface to return and handle a custom error

```go
package main
import (
    "errors"
    "fmt"
)

func riskyThing(success bool) error {
    if !success {
        return errors.New("Something went wrong")
    }
    return nil
}

func main() {
    err := riskyThing(false)
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

### ✅ Task 48: Implement a `Divide(a, b int) (int, error)` function that handles divide-by-zero using error interface

```go
package main
import (
    "errors"
    "fmt"
)

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    if result, err := Divide(10, 0); err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

### ✅ Task 49: Demonstrate how interface values hold both dynamic type and dynamic value

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
    fmt.Println("Name is:", d.name)
}

func main() {
    var s Speaker
    d := Dog{"Max"}
    s = d
    s.Speak() // Dynamic type: Dog, Dynamic value: Max
}
```

### ✅ Task 50: Create an interface `Notifier` with a method `Notify()` and show how a `User` struct implements it

```go
package main
import "fmt"

type Notifier interface {
    Notify()
}

type User struct {
    email string
}

func (u User) Notify() {
    fmt.Println("Sending email to:", u.email)
}

func main() {
    u := User{"abc@example.com"}
    var n Notifier = u
    n.Notify()
}
```
