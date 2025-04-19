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

    if N > 1 {
		lastElement := arr[N-1]
		for i := N - 1; i > 0; i-- {
			arr[i] = arr[i-1]
		}
		arr[0] = lastElement
	}

	for i := 0; i < N; i++ {
		fmt.Print(arr[i], " ")
	}
}
