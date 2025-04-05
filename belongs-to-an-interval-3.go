package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x)

    if (x > -30 && x <= -2) || (x > 7 && x <= 25) {
        fmt.Println("The point belonging to an interval")
    } else {
        fmt.Println("The point is outside the interval")
    }
}
