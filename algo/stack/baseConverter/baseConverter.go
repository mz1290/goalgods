package main

import (
	"strings"

	"github.com/mz1290/goalgods/ds/stack"
)

func baseConverter(num, base int) string {
	remStackp := new(stack.Stack)
	var strb strings.Builder
	digits := "0123456789ABCDEF"

	if base != 2 && base != 8 && base != 16 {
		return ""
	}

	for num > 0 {
		rem := num % base
		remStackp.Push(rem)

		// int / int should yield int division behavior
		num = num / base
	}

	for !remStackp.IsEmpty() {
		val := digits[remStackp.Pop().(int)]
		strb.WriteString(string(val))
	}

	return strb.String()
}
