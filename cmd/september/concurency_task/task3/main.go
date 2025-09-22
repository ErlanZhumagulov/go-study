package main

import (
	"fmt"
	"sync"
	"time"
)

func consumer(ch <-chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	for n := range ch {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Получено: %d, удвоенное: %d\n", n, n*2)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	nums := []int{1, 2, 3, 4, 5, 10, 20}

	wg.Add(1)
	go consumer(ch, &wg)

	for _, n := range nums {
		ch <- n
	}
	close(ch)

	wg.Wait()

}
