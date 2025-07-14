package main

import "fmt"

func Unique(s []string) []string {

	seen := make(map[string]bool)

	var result []string

	for _, value := range s {
		if !seen[value] {
			result = append(result, value)
			seen[value] = true
		}
	}

	return result
}

func main() {
	fmt.Println(Unique([]string{"1", "2", "3", "4", "1"}))

}
