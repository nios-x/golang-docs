package main

import (
	"fmt"
	"sync"
)

func task(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task Completed:", i)
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i < 20; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	// time.Sleep(time.Second * 3)
	wg.Wait()
}
