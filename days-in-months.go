package main

import "fmt"

func main() {
    var month int
    fmt.Scan(&month)

    daysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

    daysInMonth[1] = 29 // no leap year calculation

    fmt.Println(daysInMonth[month-1])
}
