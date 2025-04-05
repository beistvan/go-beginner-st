package main

import "fmt"

func main() {
    var d1, d2, d3 int
    fmt.Scan(&d1, &d2, &d3)

    route1 := d1 + d3 + d2
    route2 := 2 * d1 + 2 * d2
    route3 := 2 * d1 + 2 * d3
    route4 := 2 * d2 + 2 * d3

    minDistance := route1
    if route2 < minDistance {
        minDistance = route2
    }
    if route3 < minDistance {
        minDistance = route3
    }
    if route4 < minDistance {
        minDistance = route4
    }

    fmt.Println(minDistance)
}
