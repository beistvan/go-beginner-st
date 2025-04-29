package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	kStr, _ := reader.ReadString('\n')
	kStr = strings.TrimSpace(kStr)
	k, _ := strconv.Atoi(kStr)

	runes := []rune(s)

	if k >= 1 && k <= len(runes) {
		runes = append(runes[:k-1], runes[k:]...)
	}

	fmt.Println(string(runes))
}
