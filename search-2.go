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

    for i := 0; i < N; i++ {
		if arr[i] % 3 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
}
