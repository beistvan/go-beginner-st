package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)

    var num, min int
    fmt.Scan(&num)
    min = num

    for i := 1; i < n; i++ {
        fmt.Scan(&num)
        if num < min {
            min = num
        }
    }

    fmt.Println(min)
}
