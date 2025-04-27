package main

import "fmt"

func main() {
    var n, h int
    fmt.Scan(&n, &h)
    
    totalWidth := 0
    for i := 0; i < n; i++ {
        var a int
        fmt.Scan(&a)
        if a > h {
            totalWidth += 2
        } else {
            totalWidth += 1
        }
    }
    
    fmt.Println(totalWidth)
}
