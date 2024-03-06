package main

import (
	"fmt"
	"sync"
)

func printNum(start int, end int, wg *sync.WaitGroup) {
	for i := start; i <= end; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}

	// Create 10 goroutines
	for i := 0; i < 10; i++ {
		start := i*1000 + 1
		end := start + 999
		wg.Add(1)
		go printNum(start, end, wg)
		wg.Wait()
	}
}
