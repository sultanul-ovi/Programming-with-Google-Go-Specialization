# Variable Slices

## Make

- create a slice and array using `make()`
- 2-argument version:
  - specify type and length/capaciy
  - initializes to zero, length = capacity
    - `sli = make([int],10)`
- 3-argument version: specify length and capacity separately
  - `sli = make([]int, 10, 15) // slice size is 10 but array capacity is 15`

## Append

- size of a slice can be increased by append()
- adds elements to the end of a slice
- inserts into underlying array
- increases size of array if necessary
  - `sli = make([]int, 0, 3)`
- length of sli is 0
  - `sli = append(sli, 100)` - automatically increases the length of the slice
