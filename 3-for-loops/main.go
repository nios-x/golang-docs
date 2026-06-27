package main

import "fmt"

func main() {
	i := 1
	for i < 5 {
		fmt.Println("damn ", i)
		i++
	}
	// for {} // infinite

	for i := 5; i < 10; i++ {
		fmt.Println("Idx: ", i)
	}
	for i := range 3 {
		fmt.Println(i)
	}

}
