package main

import (
	"github.com/mz1290/goalgods/ds/deque"
)

func palindromeCheck(str string) bool {
	dp := deque.NewDeque()

	// Populate deque with string characters
	for _, ch := range str {
		dp.PushBack(ch)
	}

	// Iterate over
	for dp.Size() > 1 {
		frontCh := dp.PopFront()
		backCh := dp.PopBack()

		if frontCh != backCh {
			return false
		}
	}

	return true
}
