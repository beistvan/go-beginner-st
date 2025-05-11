package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var word string
		fmt.Scan(&word)
		
		if len(word) > 10 {
			abbreviation := fmt.Sprintf("%c%d%c", word[0], len(word)-2, word[len(word)-1])
			fmt.Println(abbreviation)
		} else {
			fmt.Println(word)
		}
	}
}
