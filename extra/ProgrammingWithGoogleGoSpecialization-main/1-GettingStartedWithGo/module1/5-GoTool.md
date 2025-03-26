# Go Tools

## Import

- `import` keyword is used to access other packages
- Go standard library includes many packages
  - i.e. `"fmt"`
- searches directories specified by `GOROOT` and `GOPATH`

## More on Go Tool

- a tool to manage Go source code

### Commands

- `go build` compiles the program
  - arguments can be a list of packages or a list of .go files
  - creates an executable for the main package, same name as the first .go file
  - .exe suffix for executables in Windows
- `go doc` prints documentation for a pakacge
- `go fmt` formats source code files
- `go get` downloads packages and installs them
- `go list` lists all installed packages
- `go run` compiles .go files and runs the executable
- `go test` runs tests using files ending in "_test.go"
