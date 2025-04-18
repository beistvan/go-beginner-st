package main

import (
	"fmt"
	"strconv"
)

func allUniqueDigits(year int) bool {
	yearStr := strconv.Itoa(year)
	digits := make(map[rune]bool)

	for _, digit := range yearStr {
		if digits[digit] {
			return false
		}
		digits[digit] = true
	}

	return true
}

func main() {
	var y int
	fmt.Scan(&y)

	y++
	for !allUniqueDigits(y) {
		y++
	}

	fmt.Println(y)
}
