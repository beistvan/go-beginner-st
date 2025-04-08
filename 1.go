package main

import "fmt"
import "math"

func maxDivisor(N int) int {
    maxDiv := 1

    for i := 1; i <= int(math.Sqrt(float64(N))); i++ {
        if N % i == 0 {
            if i != N && i > maxDiv {
                maxDiv = i
            }

            if N / i != N && N / i > maxDiv {
                maxDiv = N / i
            }
        }
    }

    return maxDiv
}

func main() {
    var N int
    fmt.Scan(&N)

    result := maxDivisor(N)

    fmt.Println(result)
}
