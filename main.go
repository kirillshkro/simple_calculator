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
	go func() {
		for {
			fmt.Scan(&input)
			if input == "exit" {
				quit <- 0
			}
			go calc(&line, quit, input)
		}
	}()
	for str := range line {
		fmt.Println(str)
	}
}

func calc(text *chan string, quit chan int, s string) {
	select {
	case *text <- s:
		*text <- "Echo " + <-*text
	case <-quit:
		fmt.Println("Quit")
		close(*text)
		close(quit)
		return
	}
}
