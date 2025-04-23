package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	count := 0
	for i := 0; i < n-1; i++ {
		if nums[i] < nums[i+1] {
			count++
		}
	}

	fmt.Println(count)
}
