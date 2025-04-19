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

    count := 0

	for i := 1; i < N; i++ {
		if arr[i] > arr[i-1] {
			count++
		}
	}

    fmt.Println(count)
}
