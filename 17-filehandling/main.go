package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println(stat.Name())
	fmt.Println(stat.IsDir())
	fmt.Println(stat.Size())
	fmt.Println(stat.Mode())
	fmt.Println(stat.ModTime())

	defer f.Close()

	buf := make([]byte, 10)
	d, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data", d, string(buf))
	dir, err := os.Open("./")
	if err != nil {
		panic(err)
	}
	fileInfo, err := dir.ReadDir(-1)
	for _, fi := range fileInfo {
		fmt.Println(fi)
	}
}
