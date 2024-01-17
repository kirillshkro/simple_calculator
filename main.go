package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var line string
	var result int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line = scanner.Text()
	go Calculate(line, result)
	fmt.Printf("%s = %d\n", line, result)
}
