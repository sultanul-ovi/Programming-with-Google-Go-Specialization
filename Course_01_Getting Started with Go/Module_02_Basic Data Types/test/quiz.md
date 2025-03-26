# 📝 **List of Questions**

**Q1.** Which of the following expressions does **NOT** compute the average of two integers `a` and `b`?

- ○ `avg := 2 % (a + b)`
- ○ `avg := float64(a + b) / 2`
- ○ `avg := float64(a + b) / 2.0`
- ○ `avg := float64(float64(a + b) / 2.0)`

**Q2.** What is printed when the following program is executed?

```go
func main() {
    i, _ := strconv.Atoi("10")
    y := i * 2
    fmt.Println(y)
}
```

- ○ 1010
- ○ 10
- ○ 102
- ○ 20

**Q3.** What is printed when the following program is executed?

```go
func main() {
    s := strings.Replace("ianianian", "ni", "in", 2)
    fmt.Println(s)
}
```

- ○ ianianian
- ○ iainainan
- ○ iainanian
- ○ nianiania

---

**Q4.** What is printed by this code?

```go
func main() {
    x := 7
    switch {
    case x > 3:
        fmt.Printf("1")
    case x > 5:
        fmt.Printf("2")
    case x == 7:
        fmt.Printf("3")
    default:
        fmt.Printf("4")
    }
}
```

- ○ 1
- ○ 2
- ○ 3
- ○ 4

**Q5.** What is printed by this code?

```go
func main() {
    var xtemp int
    x1 := 0
    x2 := 1
    for x := 0; x < 5; x++ {
        xtemp = x2
        x2 = x2 + x1
        x1 = xtemp
    }
    fmt.Println(x2)
}
```

- ○ 5
- ○ 13
- ○ 8
- ○ 4

---

**Q6. True or False:**  
This code compiles correctly.

```go
func main() {
    var x int
    var y *int
    z := 3
    y = &z
    x = &y
}
```

- ○ True
- ○ False

**Q7.** Which integer type provides higher accuracy?

- ○ int16
- ○ int32
- ○ int
- ○ All of these types provide the same accuracy

---

### ✅ **Answers with Explanations**

**Question 1:**  
**Which of the following expressions does NOT compute the average of two integers `a` and `b`?**  
✅ **Correct Answer:** `avg := 2 % (a + b)`  
📘 **Explanation:**  
This uses the modulo operator `%`, which returns the remainder, not an average. The other options correctly convert the sum to a float and divide by 2.

---

**Question 2:**  
**What is printed when the following program is executed?**

```go
i, _ := strconv.Atoi("10")
y := i * 2
fmt.Println(y)
```

✅ **Correct Answer:** `20`  
📘 **Explanation:**  
`Atoi("10")` converts the string to the integer 10. Multiplying it by 2 gives `20`, which is printed.

---

**Question 3:**  
**What is printed when the following program is executed?**

```go
s := strings.Replace("ianianian", "ni", "in", 2)
fmt.Println(s)
```

✅ **Correct Answer:** `iainainan`  
📘 **Explanation:**  
The `Replace()` function replaces the first 2 instances of `"ni"` with `"in"`. There are exactly 2 such instances in `"ianianian"`, and both are replaced.

---

**Question 4:**  
**What is printed by this code?**

```go
x := 7
switch {
  case x > 3:
    fmt.Printf("1")
  case x > 5:
    fmt.Printf("2")
  case x == 7:
    fmt.Printf("3")
  default:
    fmt.Printf("4")
}
```

✅ **Correct Answer:** `1`  
📘 **Explanation:**  
Since the switch has no tag, it checks conditions sequentially. The first condition `x > 3` is true, so it prints `"1"` and exits the switch.

---

**Question 5:**  
**What is printed by this code?**

```go
x1 := 0
x2 := 1
for x := 0; x < 5; x++ {
  xtemp = x2
  x2 = x2 + x1
  x1 = xtemp
}
fmt.Println(x2)
```

✅ **Correct Answer:** `8`  
📘 **Explanation:**  
This loop calculates the Fibonacci sequence. After 5 iterations, the 7th Fibonacci number is reached (0-indexed), which is `8`.

---

**Question 6:**  
**True or False: This code compiles correctly**

```go
var x int
var y *int
z := 3
y = &z
x = &y
```

✅ **Correct Answer:** **False**  
📘 **Explanation:**  
`x` is an `int`, but `&y` is a `**int` (pointer to a pointer). Assigning a pointer to an int variable causes a compile-time type error.

---

**Question 7:**  
**Which integer type provides higher accuracy?**  
✅ **Correct Answer:** **All of these types provide the same accuracy**  
📘 **Explanation:**  
Accuracy refers to whether the values represented are exact. All integer types represent exact values, so accuracy is the same. However, their **range** differs based on the number of bits.

---
