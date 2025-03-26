# Files

- linear access, not random access
  - mechanical delay
- basic operations
  - `Open` - get handle for access
  - `Read` - read bytes into []byte
  - `Write` - write bytes into []byte
  - `Close` - release handle
  - `Seek` - move read/write head

## ioutil File Read

- "io/ioutil" package has basic functions for reading and writing data

```golang
dat, err:= ioutil.ReadFile("test.txt") //ReadFile reads a whole file at once (?)
```

- `dat` is []byte filled with contents of entire file
- explicit open/close are not needed
- large files cause a problem

## ioutil File Write

```golang
dat, err:= ioutil.WriteFile("test.txt")
```

- writes []byte to file
- creates a file
- unix-style permission bytes
