package main

import (
	"testing"
)

func TestInfixToPostfix(t *testing.T) {
	t.Run("test1", func(subtest *testing.T) {
		expected := "A B * C D * +"
		actual := infixToPostfix("A * B + C * D")

		if expected != actual {
			subtest.Errorf("Expected %q but got %q", expected, actual)
		}
	})

	t.Run("test2", func(subtest *testing.T) {
		expected := "A B + C * D E - F G + * -"
		actual := infixToPostfix("( A + B ) * C - ( D - E ) * ( F + G )")

		if expected != actual {
			subtest.Errorf("Expected %q but got %q", expected, actual)
		}
	})

	t.Run("test3", func(subtest *testing.T) {
		expected := ""
		actual := infixToPostfix("10 + 3 * 5 / (16 - 4")

		if expected != actual {
			subtest.Errorf("Expected %q but got %q", expected, actual)
		}
	})
}
