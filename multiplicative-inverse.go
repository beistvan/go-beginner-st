package main

import "fmt"

func main() {
    var x float64
    fmt.Scan(&x)

    if x == 0 {
        fmt.Println("Multiplicative inverse no exist")
    } else {
        fmt.Println(1 / x)
    }
}
