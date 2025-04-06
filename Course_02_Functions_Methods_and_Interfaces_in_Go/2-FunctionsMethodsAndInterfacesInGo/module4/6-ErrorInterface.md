# Error Interface

- many Go programs return error interface objects to indicate errors

```golang
type error interface{
  Error() string
}
```

- correct operation: `error == nil`
- incorrect operation: `Error()` prints error message

## Handling Errors

```golang
f, err: := os.Open("test.txt")
if err != nil{
  fmt.Println(err)
  return
}
```

- check whether the error is nil
- if it is not nil, handle it
- `fmt` package calls the `error()` method to generate string to print
