package main

import (
	"fmt"
)

func main() {
	// 1195. 交替打印字符串

	var (
		ch1      = make(chan struct{}) // 控制fizzbuzz协程打印fizzbuzz,
		ch2      = make(chan struct{}) // 控制fizz协程打印fizz,
		ch3      = make(chan struct{}) // 控制buzz协程打印buzz,
		orderCh  = make(chan struct{}) // 控制所有协程按序列执行
		numberCh = make(chan int)      // 控制number协程打印fizzbuzz
		n        = 100
	)

	fizzbuzz := func() {
		for {
			<-ch1
			fmt.Println("fizzbuzz ")
			orderCh <- struct{}{}
		}
	}

	fizz := func() {
		for {
			<-ch2
			fmt.Println("fizz")
			orderCh <- struct{}{}
		}
	}

	buzz := func() {
		for range ch3 {
			fmt.Println("buzz")
			orderCh <- struct{}{}
		}
	}

	number := func() {
		for {
			value := <-numberCh
			fmt.Println(value)
			orderCh <- struct{}{}
		}
	}

	go fizzbuzz()
	go fizz()
	go buzz()
	go number()

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ch1 <- struct{}{}
		} else if i%3 == 0 {
			ch2 <- struct{}{}
		} else if i%5 == 0 {
			ch3 <- struct{}{}
		} else {
			numberCh <- i
		}

		<-orderCh
	}

	fmt.Println()
	close(ch1)
	close(ch2)
	close(ch3)
	close(numberCh)
	close(orderCh)
}
