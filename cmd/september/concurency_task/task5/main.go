package main

import (
	"fmt"
	"strconv"
	"sync"
)

//Сделай воркер-пул: N воркеров читают из канала задачи
//и обрабатывают их. Главная горутина пишет в канал 100 заданий. Используй
//WaitGroup, чтобы дождаться выполнения.

func main() {

	wg := &sync.WaitGroup{}
	ch := make(chan int)
	n := 10

	wg.Add(n)

	for range n {
		go func() {
			defer wg.Done()
			for {
				v, ok := <-ch
				if !ok {
					break
				}
				fmt.Println("Получено задание " + strconv.Itoa(v))
			}
		}()
	}

	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)

	wg.Wait()

}
