package main

import "fmt"

func RemoveAtIndex(s []int, index int) []int {

	if index < 0 || index >= len(s) {
		return s
	}

	var result []int

	for i := 0; i < index; i++ {
		result = append(result, s[i])
	}

	for i := index + 1; i < len(s); i++ {
		result = append(result, s[i])
	}

	return result
}

func main() {
	fmt.Println(RemoveAtIndex([]int{1, 2, 3, 4}, 2))
	fmt.Println(RemoveAtIndex([]int{5, 6, 7}, 5))
}
