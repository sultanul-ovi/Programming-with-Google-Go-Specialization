## Example questions and solutions on Go variables and initialization.

### Example 1: Declare an integer variable and assign a value

```go
package main
import "fmt"

func main() {
    var age int = 25
    fmt.Println("Age:", age)
}
```

### Example 2: Declare multiple variables of the same type

```go
package main
import "fmt"

func main() {
    var width, height int = 10, 20
    fmt.Println("Width:", width, "Height:", height)
}
```

### Example 3: Declare variables without initialization (zero value)

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

### Example 4: Use short declaration syntax inside a function

```go
package main
import "fmt"

func main() {
    message := "Hello, Go!"
    number := 42
    fmt.Println(message, number)
}
```

### Example 5: Declare and initialize a float variable

```go
package main
import "fmt"

func main() {
    var price float64 = 19.99
    fmt.Println("Price:", price)
}
```

### Example 6: Infer type from initialization value

```go
package main
import "fmt"

func main() {
    var temperature = 36.5  // inferred as float64
    fmt.Printf("Type: %T, Value: %v\n", temperature, temperature)
}
```

### Example 7: Reassign variable after declaration

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

### Example 8: Declare variables of different types in one line

```go
package main
import "fmt"

func main() {
    var id, name = 101, "Alice"
    fmt.Println("ID:", id, "Name:", name)
}
```

### Example 9: Declare constant vs variable

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

### Example 10: Use type alias for clarity

```go
package main
import "fmt"

type Celsius float64

func main() {
    var temp Celsius = 36.6
    fmt.Println("Body temperature:", temp)
}
```
