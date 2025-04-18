package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for b % 7 != 0 {
		b--
	}

	if b >= a {
		fmt.Println(b)
	} else {
		fmt.Println("NO")
	}
}
