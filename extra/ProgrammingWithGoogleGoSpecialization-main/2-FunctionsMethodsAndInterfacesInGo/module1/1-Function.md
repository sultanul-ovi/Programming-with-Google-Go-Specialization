# Function

- a set of instructions with a name

```golang
// all go programs have a main func
func main(){
  PrintHello()
}
func PrintHello(){
  fmt.Printf("Hello, world!\n")
}
```

## Reusability

- you only need to declare a function once
- good for commonly used operations
- graphics editing program might have `ThresholdImage()`
- database program might have `QueryDBase()`

## Abstraction

- details are hidden in the function
- only need to understand operations
- improves understability

```golang
func FindPupil(){
  GrabImage()
  FilterImage()
  FindEllipses()
}
```

- naming is important for clarity
