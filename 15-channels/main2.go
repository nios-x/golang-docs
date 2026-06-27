// exit blocking
package main

import "fmt"

func waitfn(chann chan string) {
	defer func() { chann <- "done" }()
	fmt.Print("Task Done")
}
func main() {
	chann := make(chan string)
	waitfn(chann)
	<-chann
}
