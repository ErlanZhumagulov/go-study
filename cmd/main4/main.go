package main

import (
	"fmt"
)

func WordCount(s []string) map[string]int {

	chekMap := make(map[string]int)

	for key, value := range s {

		fmt.Println(key, value)
		chekMap[value] = chekMap[value] + 1

	}

	return chekMap
}

func main() {

	s := []string{"a", "b", "c", "a"}
	var res map[string]int = WordCount(s)

	fmt.Println("check:  ", res)

}
