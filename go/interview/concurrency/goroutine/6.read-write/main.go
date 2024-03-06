/*
需求：测试协程和管道的使用
（1）用协程writeData() 向管道 intChan 写入50个数
（2）用协程readData() 从管道 intChan 中读取这50个数
*/
package main

import (
	"fmt"
)

func writeData(intChan chan int) {
	for i := 1; i < 51; i++ {
		intChan <- i
	}
	close(intChan)
	fmt.Println("读协程 writeData 执行完毕！")
}

// 读取协程
func readData(intChan chan int, exitChan chan bool) {
	for num := range intChan {
		fmt.Println("num=", num)
	}

	exitChan <- true
}

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	// 向管道中写入数据
	go writeData(intChan)

	// 从管道读取数据
	go readData(intChan, exitChan)

	// 退出主线程，如果没有这一步不会有任何输出，因为它将writeData和readData跑起来之后
	// 就会退出main，退出main之后，这两个协程也就退出了，不会有任何输出
	judge := <-exitChan
	if judge {
		fmt.Println("可以退出主程序main了")
	}
}
