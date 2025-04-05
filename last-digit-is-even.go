package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    lastDigit := n % 10

    if lastDigit % 2 == 0 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
