package main

import (
	"fmt"
)

func main() {
	var result chan int = make(chan int)
	var input string

	go func() {
		for {
			fmt.Scanf("%s", &input)
			if input == "exit" {
				close(result)
				return
			}
			go Calculate(input, result)
		}
	}()
	for res := range result {
		fmt.Printf("Result %s is %d\n", input, res)
	}
}
