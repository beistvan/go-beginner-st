package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func sumDigits(numStr string) int {
    sum := 0
    for _, ch := range numStr {
        digit := int(ch - '0')
        sum += digit
    }
    return sum
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
    scanner.Scan()
    num1 := strings.TrimSpace(scanner.Text())
    
    scanner.Scan()
    num2 := strings.TrimSpace(scanner.Text())
    
    sum1 := sumDigits(num1)
    sum2 := sumDigits(num2)
    
    if sum1 > sum2 {
        fmt.Println(1)
    } else if sum2 > sum1 {
        fmt.Println(2)
    } else {
        fmt.Println(0)
    }
}
