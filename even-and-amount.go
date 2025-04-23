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
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("%d ", num)
			count++
		}
	}
	fmt.Println()
	fmt.Println(count)
}
