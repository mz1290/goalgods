package main

import (
	"github.com/mz1290/goalgods/ds/stack"
)

func parChecker(str string) bool {
	sp := new(stack.Stack)

	for _, ch := range str {
		if ch == '(' {
			sp.Push(')')
		} else if ch == ')' {
			if sp.Pop() == ')' {
				continue
			} else {
				return false
			}
		} else {
			continue
		}
	}

	return sp.IsEmpty()
}
