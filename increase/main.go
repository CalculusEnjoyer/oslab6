package main

import (
	"fmt"
	"sync"
)

var (
	mutex     sync.Mutex
	sharedVar int
)

func withCriticalSection(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		sharedVar++
		mutex.Unlock()
	}
}

func withoutCriticalSection(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		sharedVar++
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go withCriticalSection(&wg)
	go withCriticalSection(&wg)
	wg.Wait()
	fmt.Println("With critical section:", sharedVar)

	sharedVar = 0

	wg.Add(2)
	go withoutCriticalSection(&wg)
	go withoutCriticalSection(&wg)
	wg.Wait()
	fmt.Println("Without critical section:", sharedVar)
}
