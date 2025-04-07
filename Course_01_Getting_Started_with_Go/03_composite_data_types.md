# Module 3: Composite Data Types

> Composite data types are types that **aggregate other data types together**. They allow you to build **complex data structures** by combining basic types such as integers, strings, and booleans.

Why Use Them?

- Essential for writing **real-world**, complex programs
- Help model intricate concepts using collections of simpler types
- Enable grouping of related data in a single structure

In this module, we will explore:

- **Arrays**: Fixed-size collections of elements of the same type
- **Slices**: Dynamically-sized, more flexible views into arrays
- **Maps**: Key-value stores for rapid data lookup
- **Structs**: Custom data types that group fields of various types

---

## Arrays in Go

- A fixed-length series of elements of a chosen type
- Elements are accessed using **subscript notation** `[]`
- Indices start at **0**
- Elements are initialized to **zero value**
  - `0` for integers, `""` for strings, `false` for booleans, etc.

### Declaring an Array

```go
var x [5]int            // Array of 5 integers
x[0] = 2                // Set first element
fmt.Println(x[1])       // Outputs 0 (zero-initialized)
```

### Array Literals

> An **array literal** is a pre-defined set of values assigned to an array.

```go
var x [5]int = [5]int{1, 2, 3, 4, 5}
```

- Array length must match the number of elements.
- You can use `...` to infer the size from the initializer:

```go
x := [...]int{1, 2, 3, 4} // Compiler infers size as 4
```

### Iterating Over Arrays

Use a `for` loop with `range` to loop over the array:

```go
x := [3]int{1, 2, 3}
for i, v := range x {
    fmt.Printf("i: %d, val: %d\n", i, v)
}
```

- `range` returns two values:
  - `i`: the index
  - `v`: the value at that index

#### Summary on Array

- Arrays are **fixed in size**, and every element must be of the same type.
- Elements are **zero-initialized** by default.
- Use `[...]` to let the compiler infer array size.
- Iterate using `for i, v := range arr` to access both index and value.

---

## Slices in Go

- A **slice** is a **window on an underlying array**
- Has **variable size** (up to full array length)\
- Slices can overlap and reference shared elements.
- Contains:
  - `Pointer`: starting position in array
  - `Length`: number of elements in the slice
  - `Capacity`: max number of elements from start to end of the array

### Slice Example

```go
arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
s1 := arr[1:3] // includes elements "b", "c"
s2 := arr[2:5] // includes elements "c", "d", "e"
```

- `s1` covers indices 1 and 2 (end is exclusive)
- `s2` covers indices 2, 3, 4

### Length and Capacity

- `len(slice)` – number of elements in the slice
- `cap(slice)` – number of elements from slice start to array end

```go
a1 := [4]string{"a", "b", "c", "d"}
s1 := a1[0:1]
fmt.Println(len(s1), cap(s1)) // Output: 1 4
```

### Accessing & Writing to Slices

- Indexing uses the same `[]` syntax as arrays
- Writing to a slice modifies the underlying array
- Multiple slices can reference the same memory

```go
s1[1] = "X" // Updates the array and any other slice sharing that element
```

### Slice Literals

- Declare and initialize a slice directly
- Compiler creates both the underlying array and the slice
- Length = Capacity = number of literal elements

```go
sli := []string{"go", "lang", "rocks"}
fmt.Println(len(sli), cap(sli)) // Output: 3 3
```

Slice literals omit the array length:

```go
// This is a slice, not an array
s := []int{1, 2, 3, 4}
```

#### Summary on slices

- Slices are dynamic views on arrays
- They support flexible sizing and shared memory
- You can slice, grow, and manipulate them efficiently

### Variable Slices

#### Make

- Create a slice and array using `make()`
- **2-argument version:**
  - Specify type and length
  - Initializes to zero, length = capacity
  - Example: `sli = make([]int, 10)`
    - Underlying array has size 10, slice has length and capacity of 10
- **3-argument version:**
  - Specify type, length, and capacity separately
  - Example: `sli = make([]int, 10, 15)`
    - Underlying array has capacity 15
    - Slice initially has length 10
    - Slice can grow up to 15 elements without creating new array

#### Append

- `append()` increases the size of a slice
- Adds one or more elements to the end of a slice
- Inserts into underlying array
- If capacity exceeded, creates a new, larger underlying array
- Example:
  ```golang
  sli := make([]int, 0, 3) // slice has length 0, capacity 3
  sli = append(sli, 100)   // slice grows to length 1
  ```
- You can keep appending elements
  - Go will automatically manage resizing the underlying array if needed

---

## Hash Table

- Contains key/value pairs
  - Examples: SSN/email, GPS coordinates/address
  - Each key must be **unique**
- Hash function computes a slot for each key
  - Processes the key to determine where to store/retrieve the associated value

### Advantages

- **Faster lookup** than lists
  - Constant-time (O(1)) vs. linear time (O(n))
- **Arbitrary keys**
  - Not limited to integers (like slices/arrays)
  - Keys can be strings, complex types, etc.

### Disadvantages

- **Collisions may occur**
  - Two keys can hash to the same slot
  - Requires collision resolution (e.g., chaining, open addressing)

#### Example

- Keys: "Joe", "Jane", "Pat"
- Values: "x", "y", "z"

Hash function determines storage slot for each key:

```
Key   -> Hash Function -> Slot
"Joe" -> hash("Joe")   -> 3
"Jane"-> hash("Jane")  -> 1
"Pat" -> hash("Pat")   -> 5
```

Values stored in the hash table:

```
Slot | Value
-----|------
1    | y
3    | x
5    | z
```

#### Notes on Hash Table

- Hash tables behave like arrays with constant-time access, but allow meaningful, non-integer keys.
- Easier for programmers to work with keys like names or identifiers.
- Collisions are possible but rare due to good hash function design.
- Go handles hashing and collision resolution internally; the user simply interacts with the map data structure.

---

## Maps in Go

> Maps in Go are the built-in implementation of hash tables. They allow you to store key/value pairs with fast lookups and flexible key types. Each key in a map must be unique, and values are accessed using the key.

You can declare a map and then use the `make()` function to allocate it:

```go
var idMap map[string]int // Declares the map variable
idMap = make(map[string]int) // Initializes the map
```

Alternatively, you can define and initialize a map using a map literal:

```go
idMap := map[string]int{"joe": 123}
```

> This creates a map with one key-value pair, mapping the key "joe" to the value `123`.

You can retrieve the value associated with a key using subscript notation:

```go
fmt.Println(idMap["joe"]) // Output: 123
```

> If the key does not exist in the map, the zero value of the value type is returned (e.g., `0` for `int`).

You can add a new key/value pair or update an existing one:

```go
idMap["jane"] = 456
idMap["jane"] = 789 // Updates value for key "jane"
```

You can remove a key/value pair using the `delete()` function:

```go
delete(idMap, "joe")
```

Go supports two-value assignment when accessing map elements. The second value indicates whether the key was present:

```go
id, present := idMap["joe"]
```

- `id` contains the value (or zero value if not present)
- `present` is a boolean (`true` if the key exists)

Use the `len()` function to get the number of key/value pairs:

```go
fmt.Println(len(idMap))
```

You can iterate over all key/value pairs in a map using a `for` loop with the `range` keyword:

```go
for key, val := range idMap {
    fmt.Println(key, val)
}
```

Each iteration assigns one key/value pair to the loop variables.

> Maps are a powerful feature in Go that make it easy to work with arbitrary key/value data. Their constant-time lookup and flexible key types make them essential for tasks like indexing and configuration storage. Key operations include creation using `make()`, accessing and modifying elements with subscript notation, checking existence with two-value assignment, and iterating with `range`.

---

## Structs in Go

> Structs are an **aggregate data type** in Go, also known as composite data types. They are used to group together different data fields (of potentially different types) into a single object. Structs are extremely useful for organizing and managing related data.


#### Why Use Structs? 
Suppose you are creating a representation of a **Person** who has the following information:

- Name
- Address
- Phone Number

#### Option 1: Use separate variables

You could use three separate variables per person:

```golang
name1 := "Joe"
addr1 := "A St."
phone1 := "123"
```

> But this approach quickly becomes unmanageable when dealing with multiple persons. There's no structure ensuring that the three fields belong together.

#### Option 2: Use a Struct

Instead, you can define a **struct** that logically groups the data fields:

```golang
type Person struct {
  name  string
  addr  string
  phone string
}
```

> This way, each instance of `Person` contains its own `name`, `addr`, and `phone` fields, and all are accessed together.



#### Declaring and Using Structs

```golang
var p1 Person  // Declares a variable of type Person
```

> Each **field** inside the struct is accessed using dot notation:

#### Reading and Writing Fields

```golang
p1.name = "Joe"         // Write to the 'name' field
x := p1.addr            // Read from the 'addr' field
```


### Initializing Structs

#### Using `new()`

```golang
p1 := new(Person)
```

- Allocates memory for a `Person` struct.
- Returns a pointer to a `Person` with zero values for all fields (e.g., empty strings for string fields).

#### Using Struct Literal (with values)

```golang
p1 := Person{
  name:  "Joe",
  addr:  "A St.",
  phone: "123",
}
```

- You can provide initial values directly using a struct literal.
- Fields can be filled in any order using the `fieldName: value` format.


#### Summary on Structure

- A struct groups related values together under a single type.
- Dot notation is used to access or modify individual fields.
- Can be initialized empty with `new()` or populated using a struct literal.
- Structs enhance code clarity and organization by binding related data together in a single unit.


---
