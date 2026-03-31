package main

import (
	"fmt"
	"sync"
)

func task(wg *sync.WaitGroup) {
	defer wg.Done() // mark task complete
	fmt.Println("Task done")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)        // number of goroutines
	go task(&wg)     // run goroutine

	wg.Wait()        // wait until done
}