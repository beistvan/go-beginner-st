package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	if N == 1 || arr[0] == arr[N - 1] {
		fmt.Println(1)
		return
	}

	uniqueCount := 1
	for i := 1; i < N; i++ {
		if arr[i] != arr[i - 1] {
			uniqueCount++
		}
	}

	if uniqueCount == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(uniqueCount)
	}
}
