package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	minVal := 10000001
	minIndex := 0

	for i := 0; i < n; i++ {
		if arr[i] < minVal {
			minVal = arr[i]
			minIndex = i
		}
	}

	fmt.Println(minIndex)
}
