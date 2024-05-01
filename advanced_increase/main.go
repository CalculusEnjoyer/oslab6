package main

import (
	"fmt"
	"sync"
)

var (
	sharedVar int
)

func increment(ch1, ch2 chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		<-ch1
		if sharedVar >= 1000 {
			ch2 <- struct{}{}
			return
		}
		sharedVar++
		ch2 <- struct{}{}
	}
}

func main() {
	sharedVar = 0
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)

	var wg sync.WaitGroup
	wg.Add(2)

	go increment(ch1, ch2, &wg)
	go increment(ch2, ch1, &wg)

	ch1 <- struct{}{}

	wg.Wait()

	fmt.Println(sharedVar)
}
