package main

import (
	"fmt"
)

func main() {
	var (
		zeroCh  = make(chan int)      // 控制zero协程打印0
		evenCh  = make(chan int)      // 控制even协程打印even
		oddCh   = make(chan int)      // 控制odd协程打印odd
		orderCh = make(chan struct{}) // 控制所有协程按序列执行
		n       = 11
	)

	zero := func() {
		for range zeroCh {
			fmt.Printf(" %d", 0)
			orderCh <- struct{}{}
		}
	}

	even := func() {
		for i := range evenCh {
			fmt.Printf("%d", i)
			orderCh <- struct{}{}
		}
	}

	odd := func() {
		for i := range oddCh {
			fmt.Printf("%d", i)
			orderCh <- struct{}{}
		}
	}

	go zero()
	go even()
	go odd()

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			zeroCh <- 0
			<-orderCh
			evenCh <- i
		} else {
			zeroCh <- 0
			<-orderCh
			oddCh <- i
		}
		<-orderCh
	}

	fmt.Println()
}
