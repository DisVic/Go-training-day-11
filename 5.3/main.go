package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&numbers[i])
	}
	str := ""
	for j := 0; j < n; j++ {
		counter := 0
		for _, number := range numbers {
			if number == numbers[j] {
				counter++
			}
		}
		if counter > 1 {
			continue
		} else {
			str += strconv.Itoa(numbers[j]) + " "
		}
	}
	fmt.Println(str)
}
