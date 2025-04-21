package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	A := make([][]int, n)
	for i := 0; i < n; i++ {
		A[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				A[i][j] = 1
			} else {
				A[i][j] = A[i-1][j] + A[i][j-1]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(A[i][j])
			if j != m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
