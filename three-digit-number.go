package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    if n >= 100 && n <= 999 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
