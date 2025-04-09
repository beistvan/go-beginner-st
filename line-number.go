package main

import "fmt"

func main() {
    var s, line string
    fmt.Scanln(&s)
    count := 1

    for {
        fmt.Scanln(&line)
        if line == s {
            fmt.Println(count)
            break
        }
        count++
    }
}
