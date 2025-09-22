// Какие проблемы в этом коде?

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	str := makeString()
	fmt.Println(time.Since(start))
	fmt.Println(str)
}

func makeString() string {
	//str := ""
	var buider strings.Builder

	for i := 0; i < 100_000; i++ {
		buider.WriteString(strconv.Itoa(i))
		//str += fmt.Sprintf("%d", i)
	}
	//return str
	return buider.String()
}
