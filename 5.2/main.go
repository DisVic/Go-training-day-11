package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}
	counter := 0
	for i := 0; i < len(array)/2; i++ {
		if array[i] == array[n-1-i] {
			counter++
		}
	}
	if counter == n/2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
