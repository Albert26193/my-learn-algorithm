package main

import (
	"fmt"
	"sync"
	"time"
)

func printNum(i int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for num := range ch {
		fmt.Println(num)
		if i == num {
			return
		}
	}
}

func main() {
	startTime := time.Now()
	ch := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go printNum(10, wg, ch)

	for i := 1; i <= 1000; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()

	duration := time.Since(startTime)
	fmt.Println(duration)
}
