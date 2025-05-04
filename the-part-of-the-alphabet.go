package main

import (
	"fmt"
)

func main() {
	var a, b rune
	fmt.Scanf("%c\n%c", &a, &b)

	if a > b {
		a, b = b, a
	}

	for c := a; c <= b; c++ {
		fmt.Printf("%c ", c)
	}
}
