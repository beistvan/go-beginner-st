package main

import (
	"fmt"
)

func main() {
	var s string
	var k int

	fmt.Scanln(&s)
	fmt.Scanln(&k)

	if k > 0 && k <= len(s) {
		fmt.Printf("%c\n", s[k-1])
	} else {
		fmt.Println("NO")
	}
}
