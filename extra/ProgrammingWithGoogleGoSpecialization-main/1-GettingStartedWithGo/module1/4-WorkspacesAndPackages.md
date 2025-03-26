# Workspace and Packages

## Workspace

- hierarchy of directories
- common organization is good for sharing
  - so people can collaborate easily
- three subdirectories
  - src - contains source code files
  - pkg - contains packages/libraries
  - bin - contains executable
- programmer typically has one workspace for many projects

### workspace information

- directory is hierarchy is recommended, not enforced
  - i.e. you can have an executable in src
- workspace directory defined by `GOPATH` environment variable
  - depends on OS
  - GOPATH is define during installation
    - C:\Users\username\go
  - Go tools assume that code is in `GOPATH`

## Packages

- group of related source file
- each package can be imported by other packages
- enables software reuse
- first line of file names package
- there must be one package called `main`
  - building the main package generates an executable program
  - `main()` is where code execution starts
