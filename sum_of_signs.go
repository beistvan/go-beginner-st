package main

import (
    "fmt"
)

func sign(x int) int {
    if x < 0 {
        return -1
    } else if x == 0 {
        return 0
    } else {
        return 1
    }
}

func main() {
    var a, b int
    fmt.Scan(&a, &b)

    z := sign(a) + sign(b)
    fmt.Println(z)
}
