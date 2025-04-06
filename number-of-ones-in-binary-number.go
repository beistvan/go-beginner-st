package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)

	binary := strconv.FormatInt(int64(N), 2)
	count := 0

	for _, digit := range binary {
		if digit == '1' {
			count++
		}
	}

	fmt.Println(count)
}
