# 📝 **Questions Section**

**Question 1:**  
What is printed when the following program is executed?

```go
x := []int{4, 8, 5}
y := -1
for _, elt := range x {
  if elt > y {
    y = elt
  }
}
fmt.Print(y)
```

---

**Question 2:**  
What is printed when the following program is executed?

```go
x := [...]int{4, 8, 5}
y := x[0:2]
z := x[1:3]
y[0] = 1
z[1] = 3
fmt.Print(x)
```

---

**Question 3:**  
What is printed when the following program is executed?

```go
x := [...]int{1, 2, 3, 4, 5}
y := x[0:2]
z := x[1:4]
fmt.Print(len(y), cap(y), len(z), cap(z))
```

---

**Question 4:**  
What is printed when the following program is executed?

```go
x := map[string]int {
  "ian": 1,
  "harris": 2,
}
for i, j := range x {
  if i == "harris" {
    fmt.Print(i, j)
  }
}
```

---

**Question 5:**  
What is printed when the following program is executed?

```go
type P struct {
  x string
  y int
}
b := P{"x", -1}
a := [...]P{P{"a", 10}, P{"b", 2}, P{"c", 3}}
for _, z := range a {
  if z.y > b.y {
    b = z
  }
}
fmt.Println(b.x)
```

---

**Question 6:**  
What is printed when the following program is executed?

```go
s := make([]int, 0, 3)
s = append(s, 100)
fmt.Println(len(s), cap(s))
```

---

### ✅ **Answers with Explanations**

---

**Answer 1:** ✔️ `8`  
📘 **Explanation:**  
The loop iterates through the slice and finds the maximum element. Initially, `y = -1`, and the largest element in `[4, 8, 5]` is `8`, so the final value of `y` is `8`.

---

**Answer 2:** ✔️ `[1 8 3]`  
📘 **Explanation:**  
Slices `y` and `z` share memory with the array `x`.

- `y[0] = 1` changes `x[0]`
- `z[1] = 3` changes `x[2]`  
  Final `x = [1, 8, 3]`

---

**Answer 3:** ✔️ `2 5 3 4`  
📘 **Explanation:**

- `len(y) = 2` (`x[0:2]`)
- `cap(y) = 5` (from index 0 to end of array)
- `len(z) = 3` (`x[1:4]`)
- `cap(z) = 4` (from index 1 to end)

---

**Answer 4:** ✔️ `harris2`  
📘 **Explanation:**  
The map is iterated, and when the key `"harris"` is found, it prints the key and value. Output: `harris2`.

---

**Answer 5:** ✔️ `a`  
📘 **Explanation:**  
Struct `P` has fields `x` and `y`. The loop finds the struct with the largest `y`.  
Among `a`, `"a"` has the largest `y = 10`, so it prints `"a"`.

---

**Answer 6:** ✔️ `1 3`  
📘 **Explanation:**  
`make([]int, 0, 3)` creates a slice with length 0 and capacity 3.  
After appending one element, `len = 1`, `cap = 3`.

---
