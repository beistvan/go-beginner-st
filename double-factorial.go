package main

import (
	"fmt"
)

func Fact2(n int) int {
	result := 1
	for i := n; i > 0; i -= 2 {
		result *= i
	}
	return result
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Print(Fact2(a), " ", Fact2(b), " ", Fact2(c))
}
