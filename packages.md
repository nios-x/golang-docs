Go Packages
===========

What is a Package?
------------------
A package is a collection of related Go source files that provide reusable
functions, types, variables, and constants.

Every Go file starts with a package declaration.

Example:
package main

Types of Packages
-----------------
1. package main
   - Executable program.
   - Must contain a main() function.

2. Library packages
   - Reusable code.
   - Imported into other packages.

Example:
package mathutil

Project Structure
-----------------

myapp/
│
├── go.mod
├── main.go
└── mathutil/
    └── add.go

mathutil/add.go
---------------
package mathutil

func Add(a, b int) int {
    return a + b
}

main.go
-------
package main

import (
    "fmt"
    "myapp/mathutil"
)

func main() {
    fmt.Println(mathutil.Add(10, 20))
}

Exported vs Unexported
----------------------

Starts with uppercase letter:
    Accessible outside the package.

Example:
func Add() {}

Starts with lowercase letter:
    Accessible only inside the same package.

Example:
func subtract() {}

Importing Packages
------------------

Standard library:
import "fmt"
import "os"
import "strings"

Local package:
import "myapp/mathutil"

Multiple imports:
import (
    "fmt"
    "os"
)

Common Standard Packages
------------------------

fmt         -> Printing and formatting
os          -> File and OS operations
io          -> Basic I/O
bufio       -> Buffered I/O
strings     -> String utilities
strconv     -> String conversions
time        -> Date and time
math        -> Mathematical functions
math/rand   -> Random numbers
encoding/json -> JSON encoding/decoding
net/http    -> HTTP server and client
sync        -> Mutexes, WaitGroups
context     -> Cancellation and timeouts

Package Rules
-------------

- One package per directory.
- Package name should match the folder name.
- Do not mix package main and another package in the same folder.
- Export identifiers by starting them with a capital letter.

Useful Commands
---------------

go run .
go build
go test
go list
go doc
go fmt
go vet

Summary
-------

Module  -> Collection of packages.
Package -> Collection of related Go files.
File    -> Individual .go source file.