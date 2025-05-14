package main

import (
    "fmt"
)

func factorial(n uint64) uint64 {
    result := uint64(1)
    for i := uint64(2); i <= n; i++ {
        result *= i
    }
    return result
}

func combinations(n, k uint64) uint64 {
    if k > n {
        return 0
    }
    return factorial(n) / (factorial(k) * factorial(n-k))
}

func main() {
    var n, k uint64
    fmt.Scan(&n)
    fmt.Scan(&k)

    fmt.Println(combinations(n, k))
}
