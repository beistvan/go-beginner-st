package main

import "fmt"

func main() {
    var age int
    var gender string

    fmt.Scan(&age)
    fmt.Scan(&gender)

    if (age >= 12 && age <= 18) && gender == "m" {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
