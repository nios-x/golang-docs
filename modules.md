package main

import (
	"os"
)

func main() {
	content := `Go Modules
==========

What is a Go Module?
--------------------
A Go module is a collection of Go packages managed together.

Initialize:
go mod init github.com/soumya/myapp

Useful Commands:
go mod tidy
go mod download
go get
go list -m all
`

	err := os.WriteFile("modules.txt", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}