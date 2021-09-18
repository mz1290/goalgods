package main

import (
	"strings"

	"github.com/mz1290/goalgods/common"
	"github.com/mz1290/goalgods/ds/stack"
)

func infixToPostfix(expr string) string {
	var prec = map[string]int{
		"*": 3,
		"/": 3,
		"+": 2,
		"-": 2,
		"(": 1,
	}
	opStackp := new(stack.Stack)
	var postfixList []string

	// Split expression into string slice
	tokens := strings.Split(expr, " ")

	for _, token := range tokens {
		// If token is digit or letter, append to postfix slice
		if common.StringIsAlpha(token) || common.StringIsNumber(token) {
			postfixList = append(postfixList, token)
		} else if token == "(" {
			opStackp.Push(token)
		} else if token == ")" {
			// Pop top off stack and
			topToken := opStackp.Pop().(string)
			// Continue to pop and append to postfix until closing ')' is popped
			for topToken != "(" {
				postfixList = append(postfixList, topToken)
				topToken = opStackp.Pop().(string)
			}
		} else { // We have an operand
			// While stack is not empty and top of stack priorty >= current operand
			for (!opStackp.IsEmpty()) && (prec[opStackp.Peek().(string)] >= prec[token]) {
				// Pop off and append to postfix format
				postfixList = append(postfixList, opStackp.Pop().(string))
			}
			// Push operand to stack
			opStackp.Push(token)
		}
	}

	// We completed iterating. If stack is not empty, append remaining tokens
	for !opStackp.IsEmpty() {
		postfixList = append(postfixList, string(opStackp.Pop().(string)))
	}

	return strings.Join(postfixList, " ")
}
