package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x)

    if x > -1 && x < 17 {
        fmt.Println("The point belonging to an interval")
    } else {
        fmt.Println("The point is outside the interval")
    }
}
