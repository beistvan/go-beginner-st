package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i+j == n-1 {
				matrix[i][j] = 1
			} else if i+j > n-1 {
				matrix[i][j] = 2
			} else {
				matrix[i][j] = 0
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j])
			if j != n-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
