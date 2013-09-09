package main

import (
	"fmt"
	"github.com/hydra13142/parser"
)

var express = &parser.Pattern{}

func init() {
	express.Add(func(s []byte) (interface{}, int) {
		switch s[0] {
		case '+', '-', '*', '/', '(', ')':
			return s[0], 1
		}
		return nil, 0
	})
	express.Add(func(s []byte) (interface{}, int) {
		a, i := parser.ParseFloat(s)
		return a, i
	})
}

func Calc(s string) (float64, error) {
	md := make([]float64, 32)
	op := make([]byte, 32)
	i, j := -1, -1

	cls := func(f func(byte) bool) bool {
		for j >= 0 && i >= 1 && f(op[j]) {
			switch op[j] {
			case '+':
				md[i-1] += md[i]
				i--
			case '-':
				md[i-1] -= md[i]
				i--
			case '*':
				md[i-1] *= md[i]
				i--
			case '/':
				md[i-1] /= md[i]
				i--
			}
			j--
		}
		return j < 0 || !f(op[j])
	}
	scanner := express.NewScanner(s, true)
	for {
		a, b, c := scanner.Scan()
		if c != nil {
			break
		}
		if b == 1 {
			switch t := a.(byte); t {
			case '(':
				j++
				op[j] = t
			case ')':
				if !cls(func(c byte) bool { return c != '(' }) {
					return 0, fmt.Errorf("wrong")
				}
				j--
			case '+', '-':
				if !cls(func(c byte) bool { return c != '(' }) {
					return 0, fmt.Errorf("wrong")
				}
				j++
				op[j] = t
			case '*', '/':
				if !cls(func(c byte) bool { return c == '*' || c == '/' }) {
					return 0, fmt.Errorf("wrong")
				}
				j++
				op[j] = t
			}
		} else {
			i++
			md[i] = a.(float64)
		}
	}
	if !cls(func(_ byte) bool { return true }) {
		return 0, fmt.Errorf("wrong")
	}
	return md[0], nil
}

func main() {
	fmt.Println(Calc("1-2+3*4"))
}
