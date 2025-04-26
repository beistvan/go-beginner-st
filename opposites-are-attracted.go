package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	numToIndex := make(map[int]int)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		numToIndex[arr[i]] = i
	}

	for i := 0; i < n; i++ {
		opposite := -arr[i]
		if idx, ok := numToIndex[opposite]; ok && idx != i {
			if i < idx {
				fmt.Println(i, idx)
			} else {
				fmt.Println(idx, i)
			}
			return
		}
	}
}
