package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	input = strings.ToLower(input)

	indexX := strings.Index(input, "x")
	indexW := strings.Index(input, "w")

	if indexX == -1 && indexW == -1 {
		fmt.Println("-1")
	} else if indexX == -1 {
		fmt.Println("w")
	} else if indexW == -1 {
		fmt.Println("x")
	} else if indexX < indexW {
		fmt.Println("x")
	} else {
		fmt.Println("w")
	}
}
