package main

// Используй sync.WaitGroup, чтобы контролировать запуск
//10 горутин, каждая из которых обрабатывает "запрос" (в виде числа)
//с задержкой 100–500 мс (используй time.Sleep).

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("Обработка запроса, пришло число " + strconv.Itoa(i))
			time.Sleep(time.Duration(rand.Intn(401)+100) * time.Millisecond)
			fmt.Println("Обработка запроса выполнена для числа " + strconv.Itoa(i))
		}(i)
	}
	wg.Wait()

}
