package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    line, _ := reader.ReadString('\n')
    line = strings.TrimSpace(line)

    words := strings.Fields(line)
    normalized := strings.Join(words, " ")

    fmt.Println(normalized)
}
