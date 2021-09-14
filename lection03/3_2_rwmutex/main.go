package main

import (
	"fmt"
)

func main() {
	stub := make(chan struct{})

	// mu := &sync.RWMutex{}
	var s = "123"
	go func() {
		for {
			// mu.RLock()
			fmt.Println(s)
			// mu.RUnlock()
		}
	}()

	var tumbler bool
	go func() {
		for {
			// mu.Lock()
			if tumbler {
				s = "123"
			} else {
				s = "ABCDE"
			}
			// mu.Unlock()

			tumbler = !tumbler
		}
	}()

	<-stub
}

// ABCDE
// 12312
// ABCDE
// ABCDE
// ABC
// ABCDE
// ABCDE
// 123
// 12312
// 123
// 123
// 123
// 123
