package main

import (
    "fmt"
    "unicode"
)

func isIdentifier(s string) bool {
    if len(s) == 0 {
        return false
    }

    first := rune(s[0])
    if !unicode.IsLetter(first) && first != '_' {
        return false
    }

    for _, ch := range s {
        if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) && ch != '_' {
            return false
        }
    }

    return true
}

func main() {
    var s string
    fmt.Scanln(&s)

    if isIdentifier(s) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
