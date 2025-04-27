package main

import "fmt"

func main() {
    var n, m int
    fmt.Scan(&n, &m)
    
    var pixel string
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            fmt.Scan(&pixel)
            if pixel == "C" || pixel == "M" || pixel == "Y" {
                fmt.Println("#Color")
                return
            }
        }
    }
    fmt.Println("#Black&White")
}
