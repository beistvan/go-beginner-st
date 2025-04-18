package main

import (
	"fmt"
)

func main() {
	var num int
	var max1, max2 int

	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}

		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}
	}

	fmt.Println(max2)
}
