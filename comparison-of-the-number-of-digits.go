package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	lenA := len(fmt.Sprint(a))
	lenB := len(fmt.Sprint(b))

	if lenA > lenB {
		fmt.Println(1)
	} else if lenB > lenA {
		fmt.Println(2)
	} else {
		fmt.Println(0)
	}
}
