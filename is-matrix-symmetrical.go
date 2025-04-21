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
			fmt.Scan(&matrix[i][j])
		}
	}

	isSymmetric := true
	for i := 0; i < n && isSymmetric; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != matrix[j][i] {
				isSymmetric = false
				break
			}
		}
	}

	if isSymmetric {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
