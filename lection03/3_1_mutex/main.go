package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int
var mu sync.Mutex

func main() {
	runtime.GOMAXPROCS(1)
	var wg = &sync.WaitGroup{}
	wg.Add(2)
	go incCounter(wg) // routine #1
	go incCounter(wg) // routine #2
	wg.Wait()

	fmt.Println("Final counter: ", counter)
}

func incCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		mu.Lock()
		value := counter
		runtime.Gosched()
		value++
		counter = value
		mu.Unlock()
	}
}
