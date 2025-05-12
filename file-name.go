package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	count := 0
	for i := 0; i <= n-3; {
		if s[i] == 'x' && s[i+1] == 'x' && s[i+2] == 'x' {
			count++
			i++
		} else {
			i++
		}
	}

	fmt.Println(count)
}
