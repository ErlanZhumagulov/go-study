package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for num := range ch {
			fmt.Println(num * 2)
		}

	}()

	for i := 0; i < 10; i++ {
		ch <- i
	}

}
