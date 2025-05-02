package main

import (
	"fmt"
	"unicode"
)

func main() {
	var c rune
	fmt.Scanf("%c", &c)

	if unicode.IsLower(c) {
		fmt.Printf("%c", unicode.ToUpper(c))
	} else if unicode.IsUpper(c) {
		fmt.Printf("%c", unicode.ToLower(c))
	}
}
