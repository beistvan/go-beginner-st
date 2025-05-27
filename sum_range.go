package main

import (
	"fmt"
)

func SumRange(X, Y int) int {
	if X > Y {
		return 0
	}
	n := Y - X + 1
	return n * (X + Y) / 2
}

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)

	sum1 := SumRange(A, B)
	sum2 := SumRange(B, C)
	total := sum1 + sum2

	fmt.Println(total)
}
