package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)

	strN := strconv.Itoa(N)
	sum := 0

	for _, digit := range strN {
		sum += int(digit - '0')
	}

	if N%sum == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
