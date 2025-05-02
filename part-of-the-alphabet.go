package main

import (
	"fmt"
)

func main() {
	var cc rune
	fmt.Scanf("%c", &cc)

	for ch := 'a'; ch <= cc; ch++ {
		if ch != 'a' {
			fmt.Print(" ")
		}
		fmt.Printf("%c", ch)
	}
}
