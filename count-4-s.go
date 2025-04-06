package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)

	strN := strconv.Itoa(N)

	count := 0

	for _, digit := range strN {
		if digit == '4' {
			count++
		}
	}

	fmt.Println(count)
}
