// Требуется реализовать функцию uniqN
// которая генерирует слайс длины N уникальных рандомных чисел

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(uniqN(10))
}

func uniqN(n int) []int {

	result := make([]int, n)
	checkMap := make(map[int]struct{})

	for i := 0; i < n; {
		addNum := rand.Int() % 10 // Просто для проверки длбавил
		
		_, flag := checkMap[addNum]
		fmt.Println(flag)

		if !flag {
			result[i] = addNum
			checkMap[addNum] = struct{}{}
			i++
		}

	}
	return result
}
