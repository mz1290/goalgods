package main

import (
	"testing"
)

func TestInfixToPostfix(t *testing.T) {
	t.Run("test1", func(subtest *testing.T) {
		expected := 3
		actual := postfixEval("7 8 + 3 2 + /")

		if expected != actual {
			subtest.Errorf("Expected %d but got %d", expected, actual)
		}
	})

	t.Run("test2", func(subtest *testing.T) {
		expected := 9
		actual := postfixEval("17 10 + 3 * 9 /")

		if expected != actual {
			subtest.Errorf("Expected %d but got %d", expected, actual)
		}
	})
}
