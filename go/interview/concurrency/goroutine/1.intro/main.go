package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello Goroutine!")
}

func main() {
	go hello()

	time.Sleep(time.Second)
	fmt.Println("main goroutine done!")
}
