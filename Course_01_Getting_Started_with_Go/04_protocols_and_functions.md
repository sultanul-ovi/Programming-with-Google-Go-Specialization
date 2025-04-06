# Module 4: Protocols and Formats

## Introduction

This module introduces standardized **protocols** and **data formats** used for communication between systems. Whether through networking, file exchange, or data storage, knowing how to handle common formats is essential for modern development.

---

## 1. Why Use Protocols and Formats?

- To **exchange data** between systems (e.g., client-server communication, file transfers)
- To ensure **interoperability** between different programming environments
- To work with **common formats** like JSON, XML, HTML, and CSV

---

## 2. Common Data Formats

### JSON (JavaScript Object Notation)

- Lightweight format for data exchange
- Human-readable and language-independent
- Common in APIs and configuration files

```python
import json

# Serialize Python object to JSON
person = {"name": "Alice", "age": 30}
json_str = json.dumps(person)
print(json_str)

# Deserialize JSON to Python object
person_back = json.loads(json_str)
print(person_back)
```

---

### HTML (HyperText Markup Language)

- Format for web documents
- Parsed and manipulated using libraries like `BeautifulSoup`

```python
from bs4 import BeautifulSoup

html = "<html><body><h1>Hello, world!</h1></body></html>"
soup = BeautifulSoup(html, 'html.parser')
print(soup.h1.text)
```

---

### CSV (Comma-Separated Values)

- Plain-text tabular format
- Common for spreadsheets, databases, exports

```python
import csv

# Reading from a CSV file
with open('data.csv', newline='') as f:
    reader = csv.reader(f)
    for row in reader:
        print(row)

# Writing to a CSV file
with open('output.csv', 'w', newline='') as f:
    writer = csv.writer(f)
    writer.writerow(['name', 'age'])
    writer.writerow(['Alice', 30])
```

---

## 3. Common Protocols

### HTTP (HyperText Transfer Protocol)

- Foundation of data communication on the web
- Used with libraries like `requests`

```python
import requests

response = requests.get('https://api.github.com')
print(response.status_code)
print(response.json())
```

---

### FTP (File Transfer Protocol)

- Protocol to transfer files between computers
- Can be used via the `ftplib` module

```python
from ftplib import FTP

ftp = FTP('ftp.example.com')
ftp.login(user='username', passwd='password')
ftp.retrlines('LIST')
ftp.quit()
```

---

## 4. Summary

- Use **standard formats** like JSON, HTML, and CSV for interoperability
- Use **protocols** like HTTP and FTP to communicate across systems
- Python has built-in and third-party libraries to handle most of these

---

## 5. Additional Packages

- `xml.etree.ElementTree` – for XML
- `yaml` – for YAML files (`pip install PyYAML`)
- `socket` – for low-level network programming

---

## End of Module 4

You are now ready to integrate standardized formats and protocols into your applications for powerful communication capabilities.

---

# Request for Comments (RFC)

## Definition

RFCs are standardized documents that define internet protocols and data formats. These are critical to ensure that systems across the internet can communicate effectively.

---

## Examples of RFCs

- **HTML**: Hypertext Markup Language (RFC 1866)
- **URI**: Uniform Resource Identifier (RFC 3986)
- **HTTP**: Hypertext Transfer Protocol (RFC 2616)

These standards ensure that:

- HTML is universally understood by browsers
- URIs are consistently formatted and interpreted
- HTTP defines how web clients and servers exchange messages

---

## Protocol Packages in Golang

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

These packages offer functions to encode and decode data, abstracting away the low-level protocol details.

---

## JSON - JavaScript Object Notation

- **RFC 7159**
- Widely accepted format to represent structured data as attribute-value pairs
- Easily maps to Go structs or maps

### Basic Types Supported

- Boolean
- Integer
- String
- Array
- Object

### Go Struct to JSON Example

#### Go Code:

```go
p1 := Person{
  Name: "Joe",
  Addr: "A St.",
  Phone: "123",
}
```

#### Equivalent JSON:

```json
{
  "name": "Joe",
  "addr": "A St.",
  "phone": "123"
}
```

This format can be transferred between systems. Any program capable of reading JSON can interpret this data, making JSON a universal format for data exchange.

---

## Summary

- RFCs define universal internet communication standards
- Golang supports key protocols via packages like `net/http` and `net`
- JSON is a widely-used data format ideal for structured information exchange

Use these standards and packages to enable seamless data interaction between systems in your applications.

---

# JSON: Properties and Processing in Golang

## JSON Properties

- Fully **Unicode**
- **Human-readable** and understandable
- **Fairly compact** – readable yet small
- Supports **recursive composition**:
  - Arrays of structs
  - Structs containing structs or arrays

---

## JSON Marshalling (Object to JSON)

Generating a JSON representation from a Go object.

### Go Struct:

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

---

## JSON Unmarshalling (JSON to Object)

Converting JSON byte array into a Go object.

### Example:

```go
var p2 Person
err := json.Unmarshal(barr, &p2)
```

- `json.Unmarshal()`:
  - Takes a JSON `[]byte` and a **pointer** to the Go object
  - Populates the Go object fields with values from the JSON
  - Returns `nil` error on success

### Important:

- Field names in struct must match JSON keys
- Types must be compatible with JSON content

---

## Summary

- JSON is Unicode-based, readable, compact, and structurally flexible
- `json.Marshal()` converts Go structs to JSON
- `json.Unmarshal()` converts JSON to Go structs
- Use matching field names and types for successful unmarshalling

---

# File Operations in Golang

## File Access Characteristics

- Files are accessed **linearly**, not randomly
  - Rooted in physical/mechanical limitations (tape-based systems)
  - Even modern file access follows this model

## Basic File Operations

1. **Open** – acquire a handle to access the file
2. **Read** – read file content into a `[]byte`
3. **Write** – write content from a `[]byte` to the file
4. **Close** – release the file handle after use
5. **Seek** – move the read/write pointer to a specific position

---

## Using `ioutil` Package

The `io/ioutil` package provides simplified file access functions.

### Reading a File

```go
dat, err := ioutil.ReadFile("test.txt")
```

- Reads entire file content into `dat` (`[]byte`)
- No need to explicitly open or close the file
- Best suited for **small files**

### Writing to a File

```go
dat := []byte("Hello, World")
ioutil.WriteFile("outfile.txt", dat, 0777)
```

- Writes full content to a file
- Creates the file if it doesn’t exist
- Uses **Unix-style permission bytes** (e.g., `0777` for full access)
- Not suitable for appending or partial writes

---

## Summary

- File access in Go mimics linear access
- Basic operations include Open, Read, Write, Close, and Seek
- `ioutil.ReadFile` and `ioutil.WriteFile` simplify file reading/writing
- Ideal for quick use with **small** to **moderate-sized** files

---

# OS Package: File Access in Golang

## Overview

Use the `os` package for precise, controlled file operations in Go. It provides functions for opening, reading, writing, and closing files at a lower level than `ioutil`.

---

## Basic Functions

- `os.Open()` – opens a file and returns a file descriptor
- `os.Close()` – closes an open file
- `File.Read()` – reads from a file into a `[]byte`
  - Reads up to the length of the byte array
  - Returns the number of bytes read
- `File.Write()` – writes from a `[]byte` into a file
- `File.WriteString()` – writes a string directly to the file

---

## File Reading Example

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

---

## File Writing Example

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

---

## Summary

- The `os` package allows detailed file manipulation
- Supports partial reads/writes and low-level control
- Ideal for large files or advanced I/O operations

---
