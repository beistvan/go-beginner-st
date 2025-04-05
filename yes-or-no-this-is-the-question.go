package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    if n % 2 != 0 {
        fmt.Println("YES")
    } else {
        if n >= 2 && n <= 5 {
            fmt.Println("NO")
        } else if n >= 6 && n <= 20 {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}
