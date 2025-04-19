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

	for i := 0; i < N-1; i++ {
		if (arr[i] > 0 && arr[i+1] > 0) || (arr[i] < 0 && arr[i+1] < 0) {
			fmt.Println("YES")
			return
		}
	}

	fmt.Println("NO")
}
