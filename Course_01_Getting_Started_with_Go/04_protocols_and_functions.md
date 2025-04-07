# Module 4: Protocols and Formats

> This module introduces standardized **protocols** and **data formats** used for communication between systems. Whether through networking, file exchange, or data storage, knowing how to handle common formats is essential for modern development.

---

## Request for Comments (RFC)

RFCs are standardized documents that define internet protocols and data formats. These are critical to ensure that systems across the internet can communicate effectively.

### Examples of RFCs

- **HTML**: Hypertext Markup Language (RFC 1866)
- **URI**: Uniform Resource Identifier (RFC 3986)
- **HTTP**: Hypertext Transfer Protocol (RFC 2616)

These standards ensure that:

- HTML is universally understood by browsers
- URIs are consistently formatted and interpreted
- HTTP defines how web clients and servers exchange messages

### Protocol Packages in Golang

Golang provides built-in packages for working with key RFC-defined protocols.

### `net/http`

- Used for web communication
- Example usage:

```go
resp, err := http.Get("http://www.uci.edu")
```

### `net`

- Lower-level TCP/IP and socket communication
- Example usage:

```go
conn, err := net.Dial("tcp", "uci.edu:80")
```

> These packages offer functions to encode and decode data, abstracting away the low-level protocol details.

---

## JSON - JavaScript Object Notation

- **RFC 7159**
- Widely accepted format to represent structured data as attribute-value pairs
- Easily maps to Go structs or maps

### Basic Types Supported

- Boolean, Integer, String, Array, Object

#### Go Code

```go
p1 := Person{
  Name: "Joe",
  Addr: "A St.",
  Phone: "123",
}
```

#### Equivalent JSON

```json
{
  "name": "Joe",
  "addr": "A St.",
  "phone": "123"
}
```

> This format can be transferred between systems. Any program capable of reading JSON can interpret this data, making JSON a universal format for data exchange.

### Notes on RFC and JSON

- RFCs define universal internet communication standards
- Golang supports key protocols via packages like `net/http` and `net`
- JSON is a widely-used data format ideal for structured information exchange

### JSON Properties

- Fully **Unicode**
- **Human-readable** and understandable
- **Fairly compact** – readable yet small
- Supports **recursive composition**:
  - Arrays of structs
  - Structs containing structs or arrays

### JSON Marshalling (Object to JSON)

> Generating a JSON representation from a Go object.

```go
type Person struct {
  Name  string
  Addr  string
  Phone string
}

p1 := Person{
  Name:  "Joe",
  Addr:  "A St.",
  Phone: "123",
}

barr, err := json.Marshal(p1)
```

- `json.Marshal()` takes a Go object and returns:
  - JSON encoded as a `[]byte`
  - `err` is `nil` if conversion is successful

### JSON Unmarshalling (JSON to Object)

> Converting JSON byte array into a Go object.

```go
var p2 Person
err := json.Unmarshal(barr, &p2)
```

- `json.Unmarshal()`:
  - Takes a JSON `[]byte` and a **pointer** to the Go object
  - Populates the Go object fields with values from the JSON
  - Returns `nil` error on success

### Important things to remember

- Field names in struct must match JSON keys
- Types must be compatible with JSON content

#### Notes on JSON

- JSON is Unicode-based, readable, compact, and structurally flexible
- `json.Marshal()` converts Go structs to JSON
- `json.Unmarshal()` converts JSON to Go structs
- Use matching field names and types for successful unmarshalling

---

## File Operations in Golang

- Files are accessed **linearly**, not randomly
  - Rooted in physical/mechanical limitations (tape-based systems)
  - Even modern file access follows this model

### Basic File Operations

1. **Open** – acquire a handle to access the file
2. **Read** – read file content into a `[]byte`
3. **Write** – write content from a `[]byte` to the file
4. **Close** – release the file handle after use
5. **Seek** – move the read/write pointer to a specific position

### Using `ioutil` Package

> The `io/ioutil` package provides simplified file access functions.

#### Reading a File

```go
dat, err := ioutil.ReadFile("test.txt")
```

- Reads entire file content into `dat` (`[]byte`)
- No need to explicitly open or close the file
- Best suited for **small files**

#### Writing to a File

```go
dat := []byte("Hello, World")
ioutil.WriteFile("outfile.txt", dat, 0777)
```

- Writes full content to a file
- Creates the file if it doesn’t exist
- Uses **Unix-style permission bytes** (e.g., `0777` for full access)
- Not suitable for appending or partial writes

#### Summary on File operations

- File access in Go mimics linear access
- Basic operations include Open, Read, Write, Close, and Seek
- `ioutil.ReadFile` and `ioutil.WriteFile` simplify file reading/writing
- Ideal for quick use with **small** to **moderate-sized** files

---

## OS Package: File Access in Golang

> Use the `os` package for precise, controlled file operations in Go. It provides functions for opening, reading, writing, and closing files at a lower level than `ioutil`.

### Basic Functions

- `os.Open()` – opens a file and returns a file descriptor
- `os.Close()` – closes an open file
- `File.Read()` – reads from a file into a `[]byte`
  - Reads up to the length of the byte array
  - Returns the number of bytes read
- `File.Write()` – writes from a `[]byte` into a file
- `File.WriteString()` – writes a string directly to the file

#### File Reading Example

```go
f, err := os.Open("dt.txt")
barr := make([]byte, 10)
nb, err := f.Read(barr)
f.Close()
```

- Opens `dt.txt`
- Creates a 10-byte slice
- Reads up to 10 bytes from file into `barr`
- `nb` contains the actual number of bytes read (may be < 10)
- Closes the file

#### File Writing Example

```go
f, err := os.Create("outfile.txt")
barr := []byte{1, 2, 3}
nb, err := f.Write(barr)
nb, err = f.WriteString("Hi")
```

- Creates `outfile.txt`
- Writes 3 raw bytes into the file
- Then writes the string "Hi" using `WriteString`
- Content must be valid Unicode for `WriteString`

## Summary on files

- The `os` package allows detailed file manipulation
- Supports partial reads/writes and low-level control
- Ideal for large files or advanced I/O operations

---
