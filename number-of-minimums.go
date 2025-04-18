package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var num int
	min := int(1e6 + 1)
	count := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if num < min {
			min = num
			count = 1
		} else if num == min {
			count++
		}
	}

	fmt.Println(count)
}
