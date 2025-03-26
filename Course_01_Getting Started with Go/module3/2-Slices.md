# Slices

- a "window" on an underlying array
- variable size, up to the whole array
- `Pointer` indicates the start of the slice
- `Length` is the numver of elements in the slice
- `Capacity` is maximum number of elements
  - from start of slice to end of array

## Slice example

```golang
arr := [...]string{"a", "b", "c", "d", "e", "f", "g"} // ellipses is optional
s1:=arr[1:3] // 1 points to the index where the slice starts, 3 is the index after the end of the slice, b,c
s2:=arr[2:5] // has c,b,d,e in the slice
```

## Length and Capacity

- `len()` function returns the length of an array/slice
- `cap()` function returns the capacity of an array
  - using this on an slice returns the capacity of the main array the slice is from

```golang
a1:=[3]string{"a", "b", "c", "d"}
s1:=a1[0:1] // has a
fmt.Printf(len(s1), cap(a1))
```

## Accessing Slices

- writing to a slice changes underlying array
- overlapping slices refer to the same array elements
- slice indexing works the same way as an array

## Slice Literals

- can be used to initialize a slice
- creates the underlying array and creates a slice to reference it
- slice oints to the start of the array, length is capacity
