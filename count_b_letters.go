package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countB(s string) int {
	count := 0
	for _, ch := range s {
		if ch == 'b' {
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	str1, _ := reader.ReadString('\n')
	str2, _ := reader.ReadString('\n')

	str1 = strings.TrimSpace(str1)
	str2 = strings.TrimSpace(str2)

	total := countB(str1) + countB(str2)

	fmt.Println(total)
}
