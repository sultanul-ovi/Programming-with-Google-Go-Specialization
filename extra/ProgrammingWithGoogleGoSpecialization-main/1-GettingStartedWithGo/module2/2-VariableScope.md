# Variable Scope

- the places in code where a variable can be accessed

```golang
// version 1
var x = 4
func f(){
  fmt.Printf("%d", x)
}

func g(){
  fmt.Printf("%d", x)
}

```

vs.

```golang
// version 2
func f(){
  var x = 4
  fmt.Printf("%d", x)
}

func g(){
  fmt.Printf("%d", x)
}
```

- both versions are basically only printing x
- version 1
  - both `f()` and `g()` prints 4
  - var x here seems to be global global
- version 2
  - only `f()` prints 4 but since `g()` doesn't have `x` declaration, the program will throw an error
  - var x here only lives inside `f()`

## Blocks

- a sequence of declarations and statements within matching brackets, `{}`
  - including function definitions
- hierarcy of implicit blocks also
- universe block - all go source
- package block - all source in a package
- file block - all source in a file
- `if`, `for`, `switch` - all code inside the statement
- clause in `switch` or `select` - individual clauses each get a block

## Lexical Scoping

- go is lexically scoped using blocks
