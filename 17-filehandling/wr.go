package main

import (
	"fmt"
	"os"
)

func main() {
	// Create file
	file, err := os.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("File created:", file.Name())

	// Write some data
	_, err = file.WriteString("Hello from Go!\n")
	if err != nil {
		panic(err)
	}

	fmt.Println("Data written.")

	// Close before deleting (especially important on Windows)
	file.Close()

	// Delete file
	err = os.Remove("newfile.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File deleted.")
}

//
