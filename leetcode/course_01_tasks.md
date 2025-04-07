# 📝 Task List

1. **(🟢 Easy)** Declare an integer variable named `age` and assign it the value `25`. Then print its value.
2. **(🟢 Easy)** Declare two integer variables named `width` and `height` with values `10` and `20`, respectively. Print both values.
3. **(🟢 Easy)** Declare an integer variable `score` and a string variable `name` without assigning them values. Print their zero values.
4. **(🟢 Easy)** Inside the `main` function, use short declaration syntax `:=` to declare a string variable `message` with value `"Hello, Go!"` and an integer `number` with value `42`. Print both.
5. **(🟢 Easy)** Declare a `float64` variable named `price` and assign it the value `19.99`. Print it.
6. **(🟢 Easy)** Declare a variable `temperature` with the value `36.5` and let Go infer its type. Then print its type and value.
7. **(🟢 Easy)** Declare an integer variable `count`, assign it the value `5`, then change it to `10`. Print the final value.
8. **(🟢 Easy)** Declare a variable `id` of type `int` with value `101` and another variable `name` of type `string` with value `"Alice"` in a single line. Print both.
9. **(🟡 Medium)** Declare a constant `pi` with value `3.1416` and a variable `radius` with value `5.0`. Calculate the area of a circle using the formula `area = pi * radius * radius` and print the result.
10. **(🟡 Medium)** Define a type alias `Celsius` for `float64`, then declare a variable `temp` of type `Celsius` with value `36.6`. Print its value.
11. **(🟢 Easy)** Declare a boolean variable `isAvailable` and assign it the value `true`. Print the value.
12. **(🟢 Easy)** Declare and print a string variable `greeting` that holds the value `"Hello, Gophers!"`.
13. **(🟢 Easy)** Use the `&` operator to get the address of an integer variable `a` initialized with `100`. Print the pointer value.
14. **(🟡 Medium)** Use the `*` operator to declare a pointer to an integer and assign the value `7` via dereferencing. Print the updated value.
15. **(🟡 Medium)** Use the `new()` function to create a pointer to an integer. Assign it the value `50` and print it using dereferencing.
16. **(🟢 Easy)** Declare a `float32` variable `discount` using scientific notation with the value `2.5e2`. Print it.
17. **(🟢 Easy)** Demonstrate type conversion by converting an `int16` variable with value `8` to `int32`. Print the converted value.
18. **(🟢 Easy)** Create a `string` from a sequence of `rune` values and print it (e.g., `72, 105` for "Hi").
19. **(🟡 Medium)** Declare a variable `x` in a global scope and access it from two different functions `f()` and `g()`. Print the value in both.
20. **(🟡 Medium)** Declare a local variable `y` inside function `h()` and attempt to access it from function `k()` to demonstrate scope limitation. Handle the resulting error.
21. **(🟢 Easy)** Declare an array `nums` of 5 integers and initialize it with the values 1 through 5. Print the third element.
22. **(🟢 Easy)** Create an array of 4 strings named `colors`, use a loop to print each index and value using `range`.
23. **(🟢 Easy)** Use array literal syntax to declare an array `scores` with elements 90, 80, 70. Let the compiler infer the size.
24. **(🟢 Easy)** Create a slice `names` with the elements "Joe", "Jane", "Jill". Append the value "Jack" and print the slice.
25. **(🟢 Easy)** Create a slice from an existing array `arr := [5]int{10, 20, 30, 40, 50}`. Slice the first three elements and print them.
26. **(🟡 Medium)** Declare a slice using `make()` with length 3 and capacity 5. Print its length and capacity.
27. **(🟡 Medium)** Append five values to a slice that has capacity 3 and print the new length and capacity.
28. **(🟢 Easy)** Create a map `studentAges` that maps names to ages. Add three entries and print them using `range`.
29. **(🟢 Easy)** Use `delete()` to remove a key from a map. Then use two-value assignment to check if a key exists.
30. **(🟡 Medium)** Declare a map using a literal with key type `string` and value type `bool`. Initialize it with two entries and print.
31. **(🟢 Easy)** Define a struct `Book` with fields `title` and `author`, both strings. Create a variable `b1` and assign values manually.
32. **(🟢 Easy)** Create a struct `Point` with fields `x` and `y` (both `int`). Use a struct literal to initialize and print the struct.
33. **(🟡 Medium)** Use `new()` to create a pointer to a `Person` struct (fields: `name`, `email`). Assign values and print them.
34. **(🟢 Easy)** Define a struct `Rectangle` with `length` and `width`. Write a function that takes a `Rectangle` and returns the area.
35. **(🟡 Medium)** Create a slice of `Book` structs with 2 books. Use a loop to print each book’s title and author.
36. **(🟢 Easy)** Use `json.Marshal()` to convert a struct `Person{Name: "Alice", Addr: "Baker St", Phone: "456"}` to JSON and print the result.
37. **(🟢 Easy)** Convert a JSON string `{"name": "Bob", "addr": "Main St", "phone": "789"}` to a `Person` struct using `json.Unmarshal()` and print the struct.
38. **(🟡 Medium)** Create a Go struct with a nested struct inside (e.g., `User` with nested `Profile`). Marshal the struct to JSON.
39. **(🟡 Medium)** Unmarshal a JSON string representing an array of `Person` structs into a Go slice and print each person's name.
40. **(🟢 Easy)** Use `http.Get()` to make a GET request to "https://api.github.com" and print the status code.
41. **(🟡 Medium)** Perform an HTTP GET request, decode the JSON response into a Go `map[string]interface{}` and print selected fields.
42. **(🔴 Hard)** Write a small HTTP server using `net/http` that responds with a JSON object when accessed at `/status`.
43. **(🟡 Medium)** Use the `os` package to create a file named `log.txt` and write the string "Session started" into it.
44. **(🟢 Easy)** Read contents of a file `info.txt` using `ioutil.ReadFile()` and print the content as a string.
45. **(🟡 Medium)** Use `os.Open()` and `Read()` to read and print the first 20 bytes of a file called `data.txt`.
46. **(🔴 Hard)** Write a program that opens a file, reads its content, converts it to uppercase, and writes the result to a new file.
47. **(🟡 Medium)** Use `json.MarshalIndent()` to pretty-print a struct as formatted JSON.
48. **(🟢 Easy)** Delete a key from a JSON-parsed Go `map[string]interface{}` and marshal it back to a string.
49. **(🟡 Medium)** Write a function to write a slice of `Person` structs into a CSV file using `encoding/csv`.
50. **(🔴 Hard)** Parse HTML from a string using `golang.org/x/net/html` and print all text inside `<h1>` tags.

---

## ✅ Task Solutions

### ✅ Task 1: Declare an integer variable named `age` and assign it the value `25`. Then print its value. (Solve using Go)

```go
package main
import "fmt"

func main() {
    var age int = 25
    fmt.Println("Age:", age)
}
```

### ✅ Task 2: Declare two integer variables named `width` and `height` with values `10` and `20`, respectively. Print both values. (Solve using Go)

```go
package main
import "fmt"

func main() {
    var width, height int = 10, 20
    fmt.Println("Width:", width, "Height:", height)
}
```

### ✅ Task 3: Declare an integer variable `score` and a string variable `name` without assigning them values. Print their zero values. (Solve using Go)

```go
package main
import "fmt"

func main() {
    var score int
    var name string
    fmt.Println("Score:", score)  // 0
    fmt.Println("Name:", name)    // ""
}
```

### ✅ Task 4: Use short declaration syntax to declare a string variable `message` and an integer variable `number`, then print them. (Solve using Go)

```go
package main
import "fmt"

func main() {
    message := "Hello, Go!"
    number := 42
    fmt.Println(message, number)
}
```

### ✅ Task 5: Declare a `float64` variable named `price` and assign it the value `19.99`. Print it. (Solve using Go)

```go
package main
import "fmt"

func main() {
    var price float64 = 19.99
    fmt.Println("Price:", price)
}
```

### ✅ Task 6: Declare a variable `temperature` with the value `36.5` and let Go infer its type. Then print its type and value. (Solve using Go)

```go
package main
import "fmt"

func main() {
    var temperature = 36.5  // inferred as float64
    fmt.Printf("Type: %T, Value: %v\n", temperature, temperature)
}
```

### ✅ Task 7: Declare an integer variable `count`, assign it the value `5`, then reassign it to `10`. Print the final value. (Solve using Go)

```go
package main
import "fmt"

func main() {
    var count int
    count = 5
    count = 10
    fmt.Println("Count:", count)
}
```

### ✅ Task 8: Declare a variable `id` (int) and a variable `name` (string) in a single line and print them. (Solve using Go)

```go
package main
import "fmt"

func main() {
    var id, name = 101, "Alice"
    fmt.Println("ID:", id, "Name:", name)
}
```

### ✅ Task 9: Declare a constant `pi` and use it to calculate and print the area of a circle. (Solve using Go)

```go
package main
import "fmt"

func main() {
    const pi = 3.1416
    var radius = 5.0
    var area = pi * radius * radius
    fmt.Println("Area:", area)
}
```

### ✅ Task 10: Define a type alias `Celsius` and use it to declare a variable `temp`, then print it. (Solve using Go)

```go
package main
import "fmt"

type Celsius float64

func main() {
    var temp Celsius = 36.6
    fmt.Println("Body temperature:", temp)
}
```

### ✅ Task 11: Declare a boolean variable `isAvailable` and assign it the value `true`. Print the value. (Solve using Go)

```go
package main
import "fmt"

func main() {
    var isAvailable bool = true
    fmt.Println("Available:", isAvailable)
}
```

### ✅ Task 12: Declare and print a string variable `greeting` that holds the value `"Hello, Gophers!"`. (Solve using Go)

```go
package main
import "fmt"

func main() {
    greeting := "Hello, Gophers!"
    fmt.Println(greeting)
}
```

### ✅ Task 13: Use the `&` operator to get the address of an integer variable `a` initialized with `100`. Print the pointer value. (Solve using Go)

```go
package main
import "fmt"

func main() {
    a := 100
    ptr := &a
    fmt.Println("Address of a:", ptr)
}
```

### ✅ Task 14: Use the `*` operator to declare a pointer to an integer and assign the value `7` via dereferencing. Print the updated value. (Solve using Go)

```go
package main
import "fmt"

func main() {
    var num int
    ptr := &num
    *ptr = 7
    fmt.Println("Dereferenced value:", *ptr)
}
```

### ✅ Task 15: Use the `new()` function to create a pointer to an integer. Assign it the value `50` and print it using dereferencing. (Solve using Go)

```go
package main
import "fmt"

func main() {
    ptr := new(int)
    *ptr = 50
    fmt.Println("Value using new():", *ptr)
}
```

### ✅ Task 16: Declare a `float32` variable `discount` using scientific notation with the value `2.5e2`. Print it. (Solve using Go)

```go
package main
import "fmt"

func main() {
    var discount float32 = 2.5e2
    fmt.Println("Discount:", discount)
}
```

### ✅ Task 17: Demonstrate type conversion by converting an `int16` variable with value `8` to `int32`. Print the converted value. (Solve using Go)

```go
package main
import "fmt"

func main() {
    var small int16 = 8
    var large int32 = int32(small)
    fmt.Println("Converted:", large)
}
```

### ✅ Task 18: Create a `string` from a sequence of `rune` values and print it (e.g., `72, 105` for "Hi"). (Solve using Go)

```go
package main
import "fmt"

func main() {
    runes := []rune{72, 105}
    str := string(runes)
    fmt.Println("String from runes:", str)
}
```

### ✅ Task 19: Declare a variable `x` in a global scope and access it from two different functions `f()` and `g()`. Print the value in both. (Solve using Go)

```go
package main
import "fmt"

var x = 10

func f() {
    fmt.Println("From f():", x)
}

func g() {
    fmt.Println("From g():", x)
}

func main() {
    f()
    g()
}
```

### ✅ Task 20: Declare a local variable `y` inside function `h()` and attempt to access it from function `k()` to demonstrate scope limitation. Handle the resulting error. (Solve using Go)

```go
package main
import "fmt"

func h() {
    var y = 20
    fmt.Println("From h():", y)
}

func k() {
    // fmt.Println("From k():", y) // ❌ Error: undefined y
    fmt.Println("From k(): cannot access y declared in h()")
}

func main() {
    h()
    k()
}
```

### ✅ Task 21: Declare an array `nums` of 5 integers and initialize it with the values 1 through 5. Print the third element

```go
package main
import "fmt"

func main() {
    var nums = [5]int{1, 2, 3, 4, 5}
    fmt.Println("Third element:", nums[2])
}
```

### ✅ Task 22: Create an array of 4 strings named `colors`, use a loop to print each index and value using `range`

```go
package main
import "fmt"

func main() {
    colors := [4]string{"Red", "Green", "Blue", "Yellow"}
    for i, v := range colors {
        fmt.Printf("Index %d: %s\n", i, v)
    }
}
```

### ✅ Task 23: Use array literal syntax to declare an array `scores` with elements 90, 80, 70. Let the compiler infer the size

```go
package main
import "fmt"

func main() {
    scores := [...]int{90, 80, 70}
    fmt.Println("Scores:", scores)
}
```

### ✅ Task 24: Create a slice `names` with the elements "Joe", "Jane", "Jill". Append the value "Jack" and print the slice

```go
package main
import "fmt"

func main() {
    names := []string{"Joe", "Jane", "Jill"}
    names = append(names, "Jack")
    fmt.Println("Names:", names)
}
```

### ✅ Task 25: Create a slice from an existing array `arr := [5]int{10, 20, 30, 40, 50}`. Slice the first three elements and print them

```go
package main
import "fmt"

func main() {
    arr := [5]int{10, 20, 30, 40, 50}
    s := arr[0:3]
    fmt.Println("Sliced array:", s)
}
```

### ✅ Task 26: Declare a slice using `make()` with length 3 and capacity 5. Print its length and capacity

```go
package main
import "fmt"

func main() {
    sli := make([]int, 3, 5)
    fmt.Println("Length:", len(sli), "Capacity:", cap(sli))
}
```

### ✅ Task 27: Append five values to a slice that has capacity 3 and print the new length and capacity

```go
package main
import "fmt"

func main() {
    sli := make([]int, 0, 3)
    for i := 1; i <= 5; i++ {
        sli = append(sli, i*10)
    }
    fmt.Println("Slice:", sli)
    fmt.Println("Length:", len(sli), "Capacity:", cap(sli))
}
```

### ✅ Task 28: Create a map `studentAges` that maps names to ages. Add three entries and print them using `range`

```go
package main
import "fmt"

func main() {
    studentAges := map[string]int{"Alice": 21, "Bob": 22, "Carol": 23}
    for name, age := range studentAges {
        fmt.Printf("%s is %d years old\n", name, age)
    }
}
```

### ✅ Task 29: Use `delete()` to remove a key from a map. Then use two-value assignment to check if a key exists

```go
package main
import "fmt"

func main() {
    studentAges := map[string]int{"Alice": 21, "Bob": 22}
    delete(studentAges, "Bob")

    age, ok := studentAges["Bob"]
    if ok {
        fmt.Println("Bob's age:", age)
    } else {
        fmt.Println("Bob not found")
    }
}
```

### ✅ Task 30: Declare a map using a literal with key type `string` and value type `bool`. Initialize it with two entries and print

```go
package main
import "fmt"

func main() {
    status := map[string]bool{"online": true, "offline": false}
    fmt.Println("Status Map:", status)
}
```

### ✅ Task 31: Define a struct `Book` with fields `title` and `author`, both strings. Create a variable `b1` and assign values manually

```go
package main
import "fmt"

type Book struct {
    title  string
    author string
}

func main() {
    var b1 Book
    b1.title = "Go Programming"
    b1.author = "John Doe"
    fmt.Println(b1)
}
```

### ✅ Task 32: Create a struct `Point` with fields `x` and `y` (both `int`). Use a struct literal to initialize and print the struct

```go
package main
import "fmt"

type Point struct {
    x int
    y int
}

func main() {
    p := Point{x: 3, y: 4}
    fmt.Println("Point:", p)
}
```

### ✅ Task 33: Use `new()` to create a pointer to a `Person` struct (fields: `name`, `email`). Assign values and print them

```go
package main
import "fmt"

type Person struct {
    name  string
    email string
}

func main() {
    p := new(Person)
    p.name = "Alice"
    p.email = "alice@example.com"
    fmt.Println(*p)
}
```

### ✅ Task 34: Define a struct `Rectangle` with `length` and `width`. Write a function that takes a `Rectangle` and returns the area

```go
package main
import "fmt"

type Rectangle struct {
    length float64
    width  float64
}

func area(r Rectangle) float64 {
    return r.length * r.width
}

func main() {
    r := Rectangle{length: 5.0, width: 3.0}
    fmt.Println("Area:", area(r))
}
```

### ✅ Task 35: Create a slice of `Book` structs with 2 books. Use a loop to print each book’s title and author

```go
package main
import "fmt"

type Book struct {
    title  string
    author string
}

func main() {
    books := []Book{
        {title: "Go in Action", author: "William Kennedy"},
        {title: "The Go Programming Language", author: "Alan Donovan"},
    }
    for _, b := range books {
        fmt.Printf("%s by %s\n", b.title, b.author)
    }
}
```

### ✅ Task 36: Use `json.Marshal()` to convert a struct `Person{Name: "Alice", Addr: "Baker St", Phone: "456"}` to JSON and print the result

```go
package main
import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name  string
    Addr  string
    Phone string
}

func main() {
    p := Person{"Alice", "Baker St", "456"}
    jsonData, _ := json.Marshal(p)
    fmt.Println(string(jsonData))
}
```

### ✅ Task 37: Convert a JSON string `{"name": "Bob", "addr": "Main St", "phone": "789"}` to a `Person` struct using `json.Unmarshal()` and print the struct

```go
package main
import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name  string
    Addr  string
    Phone string
}

func main() {
    data := `{"name":"Bob","addr":"Main St","phone":"789"}`
    var p Person
    json.Unmarshal([]byte(data), &p)
    fmt.Println(p)
}
```

### ✅ Task 38: Create a Go struct with a nested struct inside (e.g., `User` with nested `Profile`). Marshal the struct to JSON

```go
package main
import (
    "encoding/json"
    "fmt"
)

type Profile struct {
    Age int
}

type User struct {
    Name    string
    Profile Profile
}

func main() {
    u := User{"Alice", Profile{30}}
    out, _ := json.Marshal(u)
    fmt.Println(string(out))
}
```

### ✅ Task 39: Unmarshal a JSON string representing an array of `Person` structs into a Go slice and print each person's name

```go
package main
import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string
}

func main() {
    jsonData := `[{"name":"A"},{"name":"B"}]`
    var people []Person
    json.Unmarshal([]byte(jsonData), &people)
    for _, p := range people {
        fmt.Println(p.Name)
    }
}
```

### ✅ Task 40: Use `http.Get()` to make a GET request to "https://api.github.com" and print the status code

```go
package main
import (
    "fmt"
    "net/http"
)

func main() {
    res, err := http.Get("https://api.github.com")
    if err == nil {
        fmt.Println("Status:", res.StatusCode)
    }
}
```

### ✅ Task 41: Perform an HTTP GET request, decode the JSON response into a Go `map[string]interface{}` and print selected fields

```go
package main
import (
    "encoding/json"
    "fmt"
    "net/http"
)

func main() {
    res, _ := http.Get("https://api.github.com")
    defer res.Body.Close()
    var data map[string]interface{}
    json.NewDecoder(res.Body).Decode(&data)
    fmt.Println(data["current_user_url"])
}
```

### ✅ Task 42: Write a small HTTP server using `net/http` that responds with a JSON object when accessed at `/status`

```go
package main
import (
    "encoding/json"
    "net/http"
)

func main() {
    http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
    })
    http.ListenAndServe(":8080", nil)
}
```

### ✅ Task 43: Use the `os` package to create a file named `log.txt` and write the string "Session started" into it

```go
package main
import (
    "os"
)

func main() {
    f, _ := os.Create("log.txt")
    defer f.Close()
    f.WriteString("Session started")
}
```

### ✅ Task 44: Read contents of a file `info.txt` using `ioutil.ReadFile()` and print the content as a string

```go
package main
import (
    "fmt"
    "io/ioutil"
)

func main() {
    data, _ := ioutil.ReadFile("info.txt")
    fmt.Println(string(data))
}
```

### ✅ Task 45: Use `os.Open()` and `Read()` to read and print the first 20 bytes of a file called `data.txt`

```go
package main
import (
    "fmt"
    "os"
)

func main() {
    f, _ := os.Open("data.txt")
    b := make([]byte, 20)
    n, _ := f.Read(b)
    f.Close()
    fmt.Println(string(b[:n]))
}
```

### ✅ Task 46: Write a program that opens a file, reads its content, converts it to uppercase, and writes the result to a new file

```go
package main
import (
    "bytes"
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    data, _ := ioutil.ReadFile("input.txt")
    upper := strings.ToUpper(string(data))
    ioutil.WriteFile("output.txt", []byte(upper), 0777)
}
```

### ✅ Task 47: Use `json.MarshalIndent()` to pretty-print a struct as formatted JSON

```go
package main
import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{"Jane", 30}
    pretty, _ := json.MarshalIndent(p, "", "  ")
    fmt.Println(string(pretty))
}
```

### ✅ Task 48: Delete a key from a JSON-parsed Go `map[string]interface{}` and marshal it back to a string

```go
package main
import (
    "encoding/json"
    "fmt"
)

func main() {
    raw := `{"a":1,"b":2}`
    var obj map[string]interface{}
    json.Unmarshal([]byte(raw), &obj)
    delete(obj, "b")
    res, _ := json.Marshal(obj)
    fmt.Println(string(res))
}
```

### ✅ Task 49: Write a function to write a slice of `Person` structs into a CSV file using `encoding/csv`

```go
package main
import (
    "encoding/csv"
    "os"
)

type Person struct {
    Name string
    Age  string
}

func main() {
    people := []Person{{"Alice", "30"}, {"Bob", "25"}}
    f, _ := os.Create("people.csv")
    defer f.Close()
    w := csv.NewWriter(f)
    defer w.Flush()
    w.Write([]string{"Name", "Age"})
    for _, p := range people {
        w.Write([]string{p.Name, p.Age})
    }
}
```

### ✅ Task 50: Parse HTML from a string using `golang.org/x/net/html` and print all text inside `<h1>` tags

```go
package main
import (
    "fmt"
    "strings"
    "golang.org/x/net/html"
)

func main() {
    doc := `<html><body><h1>Hello</h1><p>World</p><h1>Go!</h1></body></html>`
    r := strings.NewReader(doc)
    z, _ := html.Parse(r)
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "h1" {
            if n.FirstChild != nil {
                fmt.Println(n.FirstChild.Data)
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(z)
}
```

---
