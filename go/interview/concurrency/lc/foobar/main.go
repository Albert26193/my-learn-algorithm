package main

import (
	"fmt"
	"sync"
)

const n = 5

type FooBar struct {
	fooChan chan struct{}
	barChan chan struct{}
}

func (fb *FooBar) foo(wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		<-fb.fooChan
		fmt.Println("foo", i)
		fb.barChan <- struct{}{}
	}

	wg.Done()
}

func (fb *FooBar) bar(wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		<-fb.barChan
		fmt.Println("bar")
		fb.fooChan <- struct{}{}
	}

	wg.Done()
}

func main() {
	fb := &FooBar{
		make(chan struct{}, 1),
		make(chan struct{}, 1),
	}

	wg := sync.WaitGroup{}

	wg.Add(2)

	go fb.foo(&wg)
	go fb.bar(&wg)

	fb.fooChan <- struct{}{}
	defer close(fb.barChan)
	defer close(fb.fooChan)
	wg.Wait()
}
