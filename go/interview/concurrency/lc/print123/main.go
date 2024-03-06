package main

import (
	"fmt"
	"sync"
)

type printNum struct {
	firstChan  chan struct{}
	secondChan chan struct{}
	thirdChan  chan struct{}
}

func (pn *printNum) first(wg *sync.WaitGroup) {
	<-pn.firstChan
	fmt.Println("first")
	pn.secondChan <- struct{}{}
	wg.Done()
}

func (pn *printNum) second(wg *sync.WaitGroup) {
	<-pn.secondChan
	fmt.Println("second")
	pn.thirdChan <- struct{}{}
	wg.Done()
}

func (pn *printNum) third(wg *sync.WaitGroup) {
	<-pn.thirdChan
	fmt.Println("third")
	wg.Done()
}

func main() {
	pn := &printNum{
		firstChan:  make(chan struct{}, 1),
		secondChan: make(chan struct{}, 1),
		thirdChan:  make(chan struct{}, 1),
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(3)
		pn.firstChan <- struct{}{}
		go pn.first(&wg)
		go pn.second(&wg)
		go pn.third(&wg)
		wg.Wait()
		fmt.Println("----------")
	}

}
