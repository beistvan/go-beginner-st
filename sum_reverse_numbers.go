package main

import (
	"fmt"
	"strconv"
)

func reverseNumber(n int) int {
	str := strconv.Itoa(n)
	reversed := ""

	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	result, _ := strconv.Atoi(reversed)
	return result
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	reversedA := reverseNumber(a)
	reversedB := reverseNumber(b)

	sum := reversedA + reversedB
	fmt.Println(sum)
}
