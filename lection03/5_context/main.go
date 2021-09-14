package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ctx done")
				return
			case <-ticker.C:
				fmt.Println(time.Now().Format(time.RFC1123))
			}
		}
	}()

	time.Sleep(time.Second * 10)
	cancel()

	wg.Wait()
}
