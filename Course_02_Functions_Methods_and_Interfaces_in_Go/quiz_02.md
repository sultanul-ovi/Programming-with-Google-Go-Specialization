## ✅ Questions

**1. Is the highlighted assignment to f in the following code a legal variable assignment?**

```go
var f func(string) int
func test(x string) int {
    return len(x)
}
func main() {
    f = test
}
```

- ○ Yes
- ○ No

**2. Which of the following statements correctly declares a function whose argument is another function which takes an integer as an argument and returns a string?**

- ○ func fA(fB (int) string)
- ○ func fA(fB func (int) string)
- ○ func fA(fB func (int)) string
- ○ func fA(fB func (string) int)

**3. What is an anonymous function?**

- ○ A function with no return value
- ○ A function with multiple names
- ○ A function with no name
- ○ A function with no arguments

**4. Which of the following statements correctly declares a function whose return value is another function which takes a string as an argument and returns an integer?**

- ○ func fA(fB (int) string) func (string) int
- ○ func fA(fB func (int) string) {}
- ○ func fA(fB func (int) string) int
- ○ func fA() fB func (string) int{}

**5. What does the following code print on the screen?**

```go
func fA() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
func main() {
    fB := fA()
    fmt.Print(fB())
    fmt.Print(fB())
}
```

- ○ 12
- ○ 11
- ○ 01
- ○ 1

**6. What symbols are used in a function declaration to indicate that it is a variadic function?**

- ○ "->"
- ○ ".."
- ○ "..."
- ○ "[]"

**7. What does this routine produce?**

```go
package main
import "fmt"
func main() {
    i := 1
    fmt.Print(i)
    i++
    defer fmt.Print(i+1)
    fmt.Print(i)
}
```

- ○ 132
- ○ 134
- ○ 234
- ○ 123

---

## ✅ Answers with Explanations

### Q1. Is the highlighted assignment to f in the following code a legal variable assignment?

**Answer:** Yes  
**Explanation:** Assigning a function value to a variable of the same signature is legal in Go.

---

### Q2. Which of the following statements correctly declares a function whose argument is another function which takes an integer as an argument and returns a string?

**Answer:** func fA(fB func (int) string)  
**Explanation:** This syntax is correct for a function that takes another function as an argument, which itself takes an int and returns a string.

---

### Q3. What is an anonymous function?

**Answer:** A function with no name  
**Explanation:** Anonymous functions are defined without a name and are often used as arguments or for short-term logic.

---

### Q4. Which of the following statements correctly declares a function whose return value is another function which takes a string as an argument and returns an integer?

**Answer:** func fA() func (string) int  
**Explanation:** This defines a function `fA` which returns another function that takes a string and returns an int.

---

### Q5. What does the following code print on the screen?

**Answer:** 12  
**Explanation:** The closure retains the variable `i`. First call increments to 1, second call to 2. Hence, prints 12.

---

### Q6. What symbols are used in a function declaration to indicate that it is a variadic function?

**Answer:** "..."  
**Explanation:** The `...` symbol is used in variadic functions to indicate that the function can accept a variable number of arguments.

---

### Q7. What does this routine produce?

**Answer:** 123  
**Explanation:** The program prints `1`, then `2`, and finally defers `3` (i+1). The deferred value executes after the function returns.
