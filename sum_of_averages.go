package main

import (
    "fmt"
)

func average(n int) float64 {
    return float64(n + 1) / 2.0
}

func main() {
    var n, m int
    fmt.Scan(&n, &m)

    result := average(n) + average(m)
    fmt.Println(result)
}
