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

	min := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}

	for i := 0; i < n; i++ {
		arr[i] -= min
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
