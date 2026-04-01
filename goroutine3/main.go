package main

import (
	"fmt"
	"sync"
	"time"
)

func task(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Println("Start:", id)
	time.Sleep(1 * time.Second)
	fmt.Println("End:", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go task(&wg, i)
	}
	wg.Wait()
}
