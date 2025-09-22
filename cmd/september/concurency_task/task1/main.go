package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	res := 0

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			res++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(res)
}
