package main

import (
	"fmt"
)

func main() {
	var surname, firstName, patronymic string

	fmt.Scanln(&surname)
	fmt.Scanln(&firstName)
	fmt.Scanln(&patronymic)

	initials := string(firstName[0]) + "." + string(patronymic[0]) + "."

	fmt.Printf("%s %s\n", surname, initials)
}
