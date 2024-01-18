package main

import (
	"fmt"
)

func main() {
	var line chan string = make(chan string)
	var quit chan int = make(chan int)
	var input string
	//defer close(line)
	//defer close(quit)
	fmt.Scan(&input)
	go func() {
		for {
			if input == "exit" {
				quit <- 0
			}
			calc(line, quit, input)
		}
	}()
	fmt.Println(<-line)
}

func calc(text chan string, quit chan int, s string) {
	for {
		select {
		case text <- s:
			str := "Echo " + <-text
			fmt.Println(str)
		case <-quit:
			fmt.Println("Quit")
		}
	}
}
