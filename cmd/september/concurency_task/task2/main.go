package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("Hello< from " + strconv.Itoa(i))
			time.Sleep(1 * time.Second)
		}()
	}
	wg.Wait()
}
