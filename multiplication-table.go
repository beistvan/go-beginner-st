package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Print(i * j, " ")
		}
		fmt.Println()
	}
}
