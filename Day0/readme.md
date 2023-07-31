# Installing and Hello world

## Installing

Here is the official docs for the step by step guide to install and set the  path accourding the machine you use 

[Download and install - The Go Programming Language](https://go.dev/doc/install)

### Go Modules

Go 1.11 introduced [Modules](https://github.com/golang/go/wiki/Modules), enabling an alternative workflow. This new approach will gradually [become the default](https://blog.golang.org/modules2019) mode, deprecating the use of `GOPATH`.

Modules aim to solve problems related to dependency 
management, version selection and reproducible builds; they also enable 
users to run Go code outside of `GOPATH`.

Using Modules is pretty straightforward. Select any directory outside `GOPATH` as the root of your project, and create a new module with the `go mod init` command.

A `go.mod` file will be generated, containing 
the module path, a Go version, and its dependency requirements, which 
are the other modules needed for a successful build.

If no `<modulepath>` is specified, `go mod init` will try to guess the module path from the directory structure, but it can also be overrided, by supplying an argument.

```
mkdir my-project
cd my-project
go mod init <modulepath>
```

A `go.mod` file could look like this:

```
module cmd

go 1.14

require (
        github.com/google/pprof v0.0.0-20190515194954-54271f7e092f
        golang.org/x/arch v0.0.0-20190815191158-8a70ba74b3a1
        golang.org/x/tools v0.0.0-20190611154301-25a4f137592f
)

```

The built-in documentation provides an overview of all available `go mod` commands.

`go help mod
go help mod init`

## Hello world

key concepts 

- `main` : it is the main function which runs on running the program
- `func` : keyword to define a function

**Point to remember** : It is good to separate your "domain" code from the outside world (side-effects).

```go
package main

import "fmt"

func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
}
```

## Writing tests

Writing a test is just like writing a function, with a few rules

- It needs to be in a file with a name like `xxx_test.go`
- The test function must start with the word `Test`
- The test function takes one argument only `t *testing.T`

arsenal