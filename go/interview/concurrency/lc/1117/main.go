package main

import (
	"fmt"
	"sync"
)

var hch = make(chan struct{})
var och = make(chan struct{})
var stopch = make(chan struct{})
var wg = &sync.WaitGroup{}

type H2O struct {
	Water string
}

func (h H2O) hydrogen() {
	defer wg.Done()
	for range hch {
		fmt.Print("H")
		stopch <- struct{}{}
	}

}

func (h H2O) oxygen() {
	defer wg.Done()
	for range och {
		fmt.Println("O")
		stopch <- struct{}{}
	}
}

func main() {
	h2o := H2O{Water: "OOHHHHOOHHHHOOHHHHOOHHHH"}
	wg.Add(2)
	go h2o.oxygen()
	go h2o.hydrogen()
	var mp = make(map[rune]int)

	for _, item := range h2o.Water {
		mp[item] += 1

		Hval, ok := mp['H']
		Oval, ok2 := mp['O']

		if ok && ok2 && Hval >= 2 && Oval >= 1 {
			hch <- struct{}{}
			<-stopch
			hch <- struct{}{}
			<-stopch
			och <- struct{}{}
			<-stopch
			mp['H'] -= 2
			mp['O']--
		}
	}
	close(hch)
	close(och)
	wg.Wait()
}
