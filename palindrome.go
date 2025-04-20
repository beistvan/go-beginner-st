package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	isPalindrome := true
	for i := 0; i < n/2; i++ {
		if arr[i] != arr[n-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
