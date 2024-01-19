package main

import (
	"fmt"
)

func main() {
	var line chan string = make(chan string)
	var quit chan int = make(chan int)
	var input string
	defer close(quit)

	go func() {
		for {
			fmt.Scanf("%s", &input)
			if input == "exit" {
				close(line)
				return
			}
			go calc(line, quit, input)
		}
	}()
	for str := range line {
		fmt.Println(str)
	}
}

func calc(line chan string, quit chan int, s string) {
	select {
	case line <- s:
		line <- "Echo " + s
	case <-quit:
		return
	}
}
