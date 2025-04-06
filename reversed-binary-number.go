package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)

	binary := strconv.FormatInt(int64(N), 2)
	reversed := ""

	for i := len(binary) - 1; i >= 0; i-- {
		reversed += string(binary[i])
	}

	fmt.Println(reversed)
}
