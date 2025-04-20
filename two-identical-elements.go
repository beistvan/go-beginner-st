package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    numbers := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&numbers[i])
    }

    seen := make(map[int]bool)

    for _, num := range numbers {
        if seen[num] {
            fmt.Println("YES")
            return
        }
        seen[num] = true
    }

    fmt.Println("NO")
}
