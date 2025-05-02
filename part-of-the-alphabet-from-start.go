package main

import (
	"fmt"
)

func main() {
	var cc rune
	fmt.Scanf("%c", &cc)

	for ch := cc; ch <= 'z'; ch++ {
		if ch != cc {
			fmt.Print(" ")
		}
		fmt.Printf("%c", ch)
	}
}
