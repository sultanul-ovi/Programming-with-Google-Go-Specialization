# Passing Arrays and Slices

## Passing Array Arguments

- array arguments are copied
- arrays can be big, so this can be a problem

```golang
func foo(x [3]int) int{
  return x[0]
}
func main(){
  a := [3]int{1, 2, 3}
  fmt.Print(foo(a))
}
```

## Passing Array Pointers

- possible to pass array pointers to functions

```golang
func foo(x *[3]int) int{
  (*x)[0] = (*x)[0]+1
}
func main(){
  a := [3]int{1, 2, 3}
  foo(&a)
  fmt.Print(a)
}
```

- messy and unnecessary

## Pass Slices Instead

- slices contain a pointer to the array
- passing a slice copies the pointer

```golang
func foo(sli []int) {
  sli[0] = sli[0] + 1
}
func main(){
  a := []int{1, 2, 3}
  foo(a)
  fmt.Print(a) // prints [2 2 3]
}
```
