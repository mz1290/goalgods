package main

import (
	"testing"
)

func TestSymbolCheckerValid(t *testing.T) {
	t.Run("valid1", func(subtest *testing.T) {
		expected := true
		actual := symbolChecker(("{({([][])}())}"))

		if expected != actual {
			subtest.Errorf("Expected true but got false")
		}
	})
}

func TestCheckerInvalid(t *testing.T) {
	t.Run("invalid1", func(subtest *testing.T) {
		expected := false
		actual := symbolChecker(("[{()])"))

		if expected != actual {
			subtest.Errorf("Expected false but got true")
		}
	})
}
