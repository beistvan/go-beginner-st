package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    numbers := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&numbers[i])
    }

    count := make(map[int]int)

    for _, num := range numbers {
        count[num]++
    }

    var result []int
    for _, num := range numbers {
        if count[num] == 1 {
            result = append(result, num)
        }
    }

    for i, num := range result {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(num)
    }
    fmt.Println()
}
