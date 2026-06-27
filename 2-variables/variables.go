package main

import "fmt"

func main() {
	var name string = "yopi"
	name2 := "damn shit" // short hand
	var name3 = "degf"   // type inference

	const c = 45
	const (
		abc = 5856
		def = 5785
	)
	fmt.Println(name, name2, name3, c, abc)
}
