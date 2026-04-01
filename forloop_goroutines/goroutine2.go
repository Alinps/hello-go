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

	for i := 0; i <= 5; i++ {

		wg.Add(1)
		go task(&wg)
	}
	wg.Wait()
}
