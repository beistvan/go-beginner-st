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

	minIndex, maxIndex := 0, 0
	for i := 1; i < n; i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
	}

	fmt.Println(maxIndex - minIndex)
}
