# Constants

- expression whose value is known at compile time
- type is inferred from righthand side (bool, string, number)

```golang
const x = 1.3
const (
  y = 4
  z = "Hi"
)
```

## iota

- generate a set of related but distinct constant
- often represents a property which has several distinct possible values
  - days of the week
  - months of the year
- constants must be different but actual value is not important
- like an enumerated type in other languages

```golang
type Grades int
const (
  A Grades = iota
  B
  C
  D
  F
)
```

- each constant is assigned to a unique integer value
- starts at 1 and increments
