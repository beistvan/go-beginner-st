package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str1 := scanner.Text()

	scanner.Scan()
	str2 := scanner.Text()

	if len(str1) > 0 && len(str2) > 0 && str1[0] == str2[len(str2)-1] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
