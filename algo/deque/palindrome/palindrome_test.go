package main

import (
	"testing"
)

func TestPalindromeCheck(t *testing.T) {
	t.Run("valid1", func(subtest *testing.T) {
		expected := true
		actual := palindromeCheck("radar")

		if expected != actual {
			subtest.Errorf("Expected %t but got %t", expected, actual)
		}
	})

	t.Run("valid2", func(subtest *testing.T) {
		expected := true
		actual := palindromeCheck("abccba")

		if expected != actual {
			subtest.Errorf("Expected %t but got %t", expected, actual)
		}
	})

	t.Run("invalid1", func(subtest *testing.T) {
		expected := false
		actual := palindromeCheck("abcdef")

		if expected != actual {
			subtest.Errorf("Expected %t but got %t", expected, actual)
		}
	})

	t.Run("invalid2", func(subtest *testing.T) {
		expected := false
		actual := palindromeCheck("abcdfecba")

		if expected != actual {
			subtest.Errorf("Expected %t but got %t", expected, actual)
		}
	})
}
