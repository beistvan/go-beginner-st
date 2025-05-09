package main

import (
	"fmt"
	"strings"
)

func main() {

	var input string
	fmt.Scanln(&input)

	result := strings.ReplaceAll(input, "e", "i")

	fmt.Println(result)
}
