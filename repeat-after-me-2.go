package main

import ("bufio"; "fmt";	"os")

func main() {
	reader := bufio.NewReader(os.Stdin)
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	line3, _ := reader.ReadString('\n')

	fmt.Print(line3 + "\n" + line2 + line1)
}
