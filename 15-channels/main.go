package main

import (
	"fmt"
	"math/rand"
	"time"
)

func procNum(messagechan chan int) {
	fmt.Println(<-messagechan) //blocking
}
func main() {
	messagechannel := make(chan int)
	for i := 0; i < 10; i++ {
		go procNum(messagechannel)
	}
	for i := 0; i < 10; i++ {
		messagechannel <- rand.Intn(100)
	}
	time.Sleep(time.Second)
}
