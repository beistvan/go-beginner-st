package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	matrix := make([][]int, n)

	maxSum := -1
	maxIndex := -1

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		sum := 0
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix[i][j])
			sum += matrix[i][j]
		}
		if sum > maxSum {
			maxSum = sum
			maxIndex = i
		}
	}

	fmt.Println(maxSum)
	fmt.Println(maxIndex)
}
