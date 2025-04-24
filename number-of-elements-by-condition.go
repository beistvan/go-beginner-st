package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	count := 0
	for _, num := range arr {
		if num % 3 == 0 && num % 10 == 7 {
			count++
		}
	}

	for i, num := range arr {
		if num % 3 == 0 && num % 10 == 7 {
			arr[i] = count
		}
	}

	for i, val := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()
}
