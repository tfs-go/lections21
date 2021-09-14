package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Starting...")

	go func() {
		defer wg.Done()
		for char := 'a'; char < 'a'+26; char++ {
			// runtime.Gosched() // 1
			fmt.Printf("%c ", char)
			// time.Sleep(150 * time.Nanosecond) // 2
		}
	}()

	go func() {
		defer wg.Done()
		for char := 'A'; char < 'A'+26; char++ {
			// runtime.Gosched() // 1
			fmt.Printf("%c ", char)
			// time.Sleep(150 * time.Nanosecond) // 2
		}
	}()

	wg.Wait()
	fmt.Println("\nFinished")
}
