package main

import ("bufio"; "fmt"; "os")

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	var str string
    scaner.Scan()
    sp := scaner.Text()
	for i := 0; i < 3; i++ {
		scaner.Scan()
        if (i == 2) {sp = ""}
        str += scaner.Text() + sp
	}
	fmt.Print(str)
}
