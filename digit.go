package main

import (
	"fmt"
	"unicode"
)

func main() {
	var cc rune
	fmt.Scanf("%c", &cc)

	if unicode.IsDigit(cc) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
