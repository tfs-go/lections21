package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	buffered := make(chan string, 10)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= 9; i++ {
			fmt.Println("write to channel")
			buffered <- fmt.Sprintf("Hello #%d", i)
		}
		close(buffered)
		fmt.Println("close channel")
	}()

	time.Sleep(time.Second * 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			v, ok := <-buffered
			if !ok {
				fmt.Println("stop reader")
				return
			}

			fmt.Println(v)
		}
	}()

	wg.Wait()
}
