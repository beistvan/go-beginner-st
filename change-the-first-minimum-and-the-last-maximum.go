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

	minIndex := 0
	for i := 1; i < n; i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}


	maxIndex := 0
	for i := 1; i < n; i++ {
		if arr[i] >= arr[maxIndex] {
			maxIndex = i
		}
	}

	arr[minIndex], arr[maxIndex] = arr[maxIndex], arr[minIndex]

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
