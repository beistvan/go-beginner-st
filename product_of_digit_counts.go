package main

import (
	"fmt"
	"strconv"
)

func countDigits(n int) int {
	return len(strconv.Itoa(n))
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	
	digitsN := countDigits(n)
	digitsM := countDigits(m)
	
	fmt.Println(digitsN * digitsM)
}
