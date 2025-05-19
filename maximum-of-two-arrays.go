package main

import (
	"fmt"
)

func maxInArray(arr []int) int {
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

func main() {
	var n, m int

	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Scan(&m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}

	maxA := maxInArray(a)
	maxB := maxInArray(b)
	fmt.Println(maxA * maxB)
}
