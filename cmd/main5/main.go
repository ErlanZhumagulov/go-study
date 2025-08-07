package main

import (
	"fmt"
)

func InvertMap(originMap map[string]string) map[string][]string {

	resultMap := make(map[string][]string)

	for key, value := range originMap {

		fmt.Println(key, value)

		resultMap[value] = append(resultMap[value], key)

		fmt.Println("res", resultMap)
	}

	return nil
}

func main() {

	InvertMap(map[string]string{"1": "a", "2": "b", "3": "c", "4": "a"})

	fmt.Print("hello")
}
