package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	result := ""
	for i := 0; i < len(s); {
		if s[i] == '.' {
			result += "0"
			i++
		} else if s[i] == '-' {
			if i+1 < len(s) && s[i+1] == '.' {
				result += "1"
			} else {
				result += "2"
			}
			i += 2
		}
	}

	fmt.Println(result)
}
