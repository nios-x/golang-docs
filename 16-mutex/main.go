package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func task(x *int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	*x++
	mu.Unlock()
}

func main() {
	x := 0
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go task(&x, &wg)
	}
	wg.Wait()
	fmt.Println(x)
	time.Sleep(time.Second * 2)
}
