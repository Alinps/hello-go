package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumber(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		fmt.Println("number: ", i)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'A'; i <= 'G'; i++ {
		fmt.Println("Letters: ", string(i))
	}
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumber(&wg)
	go printLetters(&wg)
	elapsed := time.Since(start)
	wg.Wait()
	fmt.Println("Time taken: ", elapsed)
}
