package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input := scanner.Text()

	var digits []string

	for _, ch := range input {
		if unicode.IsDigit(ch) {
			digits = append(digits, string(ch))
		}
	}

	fmt.Println(strings.Join(digits, " "))
}
