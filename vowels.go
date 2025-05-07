package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	vowels := "aeiou"
	count := 0

	for _, char := range input {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	fmt.Println(count)
}
