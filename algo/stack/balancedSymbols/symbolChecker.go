package main

import (
	"fmt"

	"github.com/mz1290/goalgods/ds/stack"
)

func symbolChecker(str string) bool {
	sp := new(stack.Stack)

	for _, ch := range str {
		switch ch {
		case '(':
			fmt.Printf("(")
			sp.Push(')')
		case ')':
			fmt.Printf(")")
			if sp.Pop() != ')' {
				return false
			}
		case '[':
			fmt.Printf("[")
			sp.Push(']')
		case ']':
			fmt.Printf("]")
			if sp.Pop() != ']' {
				return false
			}
		case '{':
			fmt.Printf("{")
			sp.Push('}')
		case '}':
			fmt.Printf("}")
			if sp.Pop() != '}' {
				return false
			}
		}
	}

	return sp.IsEmpty()
}
