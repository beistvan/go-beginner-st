package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)

    var num int
    fmt.Scan(&num)
    min, max := num, num

    for i := 1; i < n; i++ {
        fmt.Scan(&num)
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
    }

    fmt.Println(max - min)
}
