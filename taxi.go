package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    var n int
    fmt.Fscan(reader, &n)

    c := [5]int{}
    for i := 0; i < n; i++ {
        var s int
        fmt.Fscan(reader, &s)
        c[s]++
    }

    taxis := 0

    taxis += c[4]

    taxis += c[3]
    if c[1] >= c[3] {
        c[1] -= c[3]
    } else {
        c[1] = 0
    }

    taxis += c[2] / 2
    if c[2]%2 == 1 {
        taxis++
        if c[1] >= 2 {
            c[1] -= 2
        } else {
            c[1] = 0
        }
    }

    if c[1] > 0 {
        taxis += (c[1] + 3) / 4
    }

    fmt.Println(taxis)
}
