package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)

	strN := strconv.Itoa(N)
	reversed := ""

	for i := len(strN) - 1; i >= 0; i-- {
		reversed += string(strN[i])
	}

	fmt.Println(reversed)
}
