package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    _ = scanner.Scan()
    string1 := scanner.Text()
    _ = scanner.Scan()
    string2 := scanner.Text()
    _ = scanner.Scan()
    string3 := scanner.Text()
    fmt.Println(string1)
    fmt.Println(string2)
    fmt.Println(string3)
}
