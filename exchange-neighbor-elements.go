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

    for i := 0; i < N-1; i += 2 {
		arr[i], arr[i+1] = arr[i+1], arr[i]
	}

	for i := 0; i < N; i++ {
		fmt.Print(arr[i], " ")
	}
}
