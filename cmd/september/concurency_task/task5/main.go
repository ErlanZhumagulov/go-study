package main

import (
	"fmt"
	"sync"
)

func main() {

	const numWorkers = 3
	const numTasks = 100 // много задач!

	var wg sync.WaitGroup
	tasks := make(chan int, 10)  // небольшой буфер для задач
	results := make(chan int, 3) // небольшой буфер для результатов

	// Запуск воркеров
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for task := range tasks {
				results <- task * task
			}
		}(i)
	}

	// Горутина для отправки задач
	go func() {
		defer close(tasks)
		for i := 0; i < numTasks; i++ {
			tasks <- i
		}
	}()

	// Горутина для получения результатов
	go func() {
		for i := 0; i < numTasks; i++ {
			result := <-results
			fmt.Printf("Result: %d\n", result)
		}
	}()

	wg.Wait()      // все воркеры закончили
	close(results) // закрываем канал результатов

	fmt.Println("v1 completed\n")
}
