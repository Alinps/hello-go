package main

import (
	"fmt"
	"sync"
)

func task(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task done")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go task(&wg)
	go task(&wg)
	wg.Wait()
}
