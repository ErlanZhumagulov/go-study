package main

import "fmt"

func RemoveAtIndex(s []int, index int, el int) []int {

	if index < 0 || index >= len(s) {
		return s
	}

	var result []int

	for i := 0; i < index; i++ {
		result = append(result, s[i])
	}

	result = append(result, el)

	for i := index; i < len(s); i++ {
		result = append(result, s[i])
	}

	return result
}

func main() {
	fmt.Println(RemoveAtIndex([]int{1, 2, 3, 4}, 2, 2))

}
