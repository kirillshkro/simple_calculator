package main

import (
	"strconv"
)

type Token struct {
	operand  string
	operand2 string
	opKind   OperandKind
}

type OperandKind byte

const (
	plus  OperandKind = '+'
	minus OperandKind = '-'
	mul   OperandKind = '*'
	div   OperandKind = '/'
	eol   OperandKind = '\n'
)

func isOperand(sym byte) bool {
	return sym == byte(plus) || sym == byte(minus) || sym == byte(mul) || sym == byte(div)
}

func trim(s string) string {
	result := ""
	for _, item := range s {
		if item != ' ' {
			result += string(item)
		}
	}
	return result
}

func parse(str string) (Token, error) {
	var pos int = 0
	token := Token{
		operand:  "",
		operand2: "",
		opKind:   eol,
	}
	trimmed := trim(str)
	for ; pos < len(trimmed); pos++ {
		item := trimmed[pos]
		if ('0' <= item) && (item <= '9') {
			token.operand += string(item)
		} else if isOperand(item) {
			token.opKind = OperandKind(item)
			break
		}
	}
	right := pos + 1
	token.operand2 = trimmed[right:]
	return token, nil
}

func Calculate(str string, result chan int) {
	var res chan int = make(chan int)
	if str == "quit" || str == "exit" {
		close(result)
		close(res)
		return
	}
	go func(s string, out chan int) {
		line, _ := parse(s)
		res := 0
		op1, err1 := strconv.Atoi(line.operand)
		op2, err2 := strconv.Atoi(line.operand2)
		switch line.opKind {
		case plus:
			if err1 == nil && err2 == nil {
				res = op1 + op2
			}
		case minus:
			if err1 == nil && err2 == nil {
				res = op1 - op2
			}
		case mul:
			res = op1 * op2
		case div:
			res = op1 / op2
		}
		out <- res
	}(str, res)
	result = res
}
