package main

import (
    "fmt"
)

func main() {
    var s string
    fmt.Scanln(&s)

    letterCount := make(map[rune]int)

    for _, ch := range s {
        letterCount[ch]++
        if letterCount[ch] == 2 {
            fmt.Println(string(ch))
            break
        }
    }
}
