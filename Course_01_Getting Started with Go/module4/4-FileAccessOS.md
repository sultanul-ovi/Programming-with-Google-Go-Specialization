# OS Package File Access

- `os.Open()` opens a file
  - returns a file descriptor (File)
- `os.Close()` closes a file
- `os.Read()` reads from a file into a []byte
  - fills the []byte
  - control the amount read
- `os.Write()` writes a []byte into a file

## os File Reading

- opening and reading

```golang
f, err := os.Open("dt.txt")
barr := make([]byte, 10)
nb, err := f.Read(barr) // fills with the first 10 bytes, and then every next 10
f.Close()
```

- reads and fills `barr`
- `Read` returns # of bytes read
- may be less than []byte length

## os File Create/Write

```golang
f, err := os.Create("outfile.txt")
barr := []byte{1, 2, 3}
nb, err := f.Write(barr)
nb, err := f.WriteString("Hi")
```

- `WriteString()` writes a string
- `Write()` writes a []byte
  - any unicode sequence
