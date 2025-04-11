package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)

    var num, max int
    fmt.Scan(&num)
    max = num

    for i := 1; i < n; i++ {
        fmt.Scan(&num)
        if num > max {
            max = num
        }
    }

    fmt.Println(max)
}
