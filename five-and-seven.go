package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input := scanner.Text()
		var result []byte

		for i := 0; i < len(input); i++ {
			if input[i] != '5' && input[i] != '7' {
				result = append(result, input[i])
			}
		}

		i := 0
		for i < len(result) && result[i] == '0' {
			i++
		}

		if i == len(result) {
			fmt.Println("0")
		} else {
			fmt.Println(string(result[i:]))
		}
	}
}
