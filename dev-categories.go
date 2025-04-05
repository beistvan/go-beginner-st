package main

import "fmt"

func main() {
    var techCount int
    fmt.Scan(&techCount)

    if techCount <= 3 {
        fmt.Println("beginner")
    } else if techCount >= 4 && techCount <= 7 {
        fmt.Println("junior developer")
    } else if techCount >= 8 && techCount <= 15 {
        fmt.Println("medior developer")
    } else {
        fmt.Println("senior developer")
    }
}
