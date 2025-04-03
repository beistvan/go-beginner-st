package main

import "fmt"

func main() {

    var num int
    fmt.Scan(&num)

    firstDigit := num / 100000
    secondDigit := (num / 10000) % 10
    thirdDigit := (num / 1000) % 10
    fourthDigit := (num / 100) % 10
    fifthDigit := (num / 10) % 10
    sixthDigit := num % 10

    sumFirst := firstDigit + secondDigit + thirdDigit
    sumLast := fourthDigit + fifthDigit + sixthDigit

    if sumFirst == sumLast {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
