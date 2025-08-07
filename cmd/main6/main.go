package main

import (
	"fmt"
)

func MapIsEquals(m1, m2 map[string]int) bool {

	if len(m1) != len(m2) {
		return false
	} else {
		for key, _ := range m1 {

			value, ex := m2[key]
			fmt.Println(value, ex)

			if !ex || (m1[key] != m2[key]) {
				return false
			}
		}
		return true
	}

}

func main() {

	resultBoolean := MapIsEquals((map[string]int{"1": 1, "2": 0}), (map[string]int{"1": 1, "3": 2}))

	fmt.Println("result: ", resultBoolean)

	resultBoolean2 := MapIsEquals((map[string]int{"1": 1, "2": 2}), (map[string]int{"1": 1, "2": 2}))

	fmt.Println("result2: ", resultBoolean2)

	fmt.Println("hello")

}
