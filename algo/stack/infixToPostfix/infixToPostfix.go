package main

import (
	"strings"
	"unicode"

	"github.com/mz1290/goalgods/ds/stack"
)

func infixToPostfix(expr string) string {
	var prec = map[rune]int{
		'*': 3,
		'/': 3,
		'+': 2,
		'-': 2,
		'(': 1,
	}
	opStackp := new(stack.Stack)
	var postfixList []string

	for _, ch := range expr {
		if ch == ' ' {
			continue
		}

		if unicode.IsDigit(ch) || unicode.IsLetter(ch) {
			postfixList = append(postfixList, string(ch))
		} else if ch == '(' {
			opStackp.Push(ch)
		} else if ch == ')' {
			topToken := opStackp.Pop().(rune)
			for topToken != '(' {
				postfixList = append(postfixList, string(topToken))
				topToken = opStackp.Pop().(rune)
			}
		} else {
			for (!opStackp.IsEmpty()) && (prec[opStackp.Peek().(rune)] >= prec[(ch)]) {
				postfixList = append(postfixList, string(opStackp.Pop().(rune)))
			}
			opStackp.Push(ch)
		}
	}

	for !opStackp.IsEmpty() {
		postfixList = append(postfixList, string(opStackp.Pop().(rune)))
	}

	return strings.Join(postfixList, " ")
}
