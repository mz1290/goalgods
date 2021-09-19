package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mz1290/goalgods/common"
	"github.com/mz1290/goalgods/ds/stack"
)

func postfixEval(expr string) int {
	// Stack for push operands
	operandStackp := new(stack.Stack)

	// Split expression into string slice
	tokens := strings.Split(expr, " ")

	for _, token := range tokens {
		if common.StringIsNumber(token) {
			res, err := strconv.Atoi(token)
			if err != nil {
				fmt.Println("Failed to convert str to number")
				return 0
			}
			operandStackp.Push(res)
		} else { // Operator
			op1 := operandStackp.Pop().(int)
			op2 := operandStackp.Pop().(int)
			switch token {
			case "+":
				operandStackp.Push(op1 + op2)
			case "-":
				operandStackp.Push(op1 - op2)
			case "*":
				operandStackp.Push(op1 * op2)
			case "/":
				operandStackp.Push(op2 / op1)
			default:
				fmt.Printf("Invalid operator %q\n", token)
				return 0
			}
		}
	}

	return operandStackp.Pop().(int)
}
