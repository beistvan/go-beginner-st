package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var speed, maxSpeed int
	var isBelow30 bool

	maxSpeed = -1

    for i := 0; i < n; i++ {
		fmt.Scan(&speed)
		if speed < 30 {
			isBelow30 = true
		}
		if speed > maxSpeed {
			maxSpeed = speed
		}
	}

	if isBelow30 {
		fmt.Printf("%d YES\n", maxSpeed)
	} else {
		fmt.Printf("%d NO\n", maxSpeed)
	}
}
