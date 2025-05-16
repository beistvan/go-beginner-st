package main

import (
	"fmt"
	"strconv"
)

func isHappyTicket(ticket int) bool {
	ticketStr := strconv.Itoa(ticket)
	firstSum := int(ticketStr[0]-'0') + int(ticketStr[1]-'0') + int(ticketStr[2]-'0')
	secondSum := int(ticketStr[3]-'0') + int(ticketStr[4]-'0') + int(ticketStr[5]-'0')
	return firstSum == secondSum
}

func main() {
	var ticket1, ticket2 int

	fmt.Scan(&ticket1, &ticket2)
	
	if isHappyTicket(ticket1) && isHappyTicket(ticket2) {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}
}
