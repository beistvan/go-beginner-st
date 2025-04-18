package main

import (
	"fmt"
)

func main() {
	var a, n int
	fmt.Scan(&a, &n)

	result := 1
	for i := 0; i < n; i++ {
		result *= a
	}

	fmt.Println(result)
}
