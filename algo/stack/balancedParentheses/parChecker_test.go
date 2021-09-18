package main

import (
	"testing"
)

func TestParCheckerValid(t *testing.T) {
	t.Run("valid1", func(subtest *testing.T) {
		expected := true
		actual := parChecker(("((()))"))

		if expected != actual {
			t.Errorf("Expected true but got false")
		}
	})

	t.Run("valid2", func(subtest *testing.T) {
		expected := true
		actual := parChecker(("((()()))"))

		if expected != actual {
			t.Errorf("Expected true but got false")
		}
	})
}

func TestParCheckerInvalid(t *testing.T) {
	t.Run("invalid1", func(subtest *testing.T) {
		expected := false
		actual := parChecker(("(()"))

		if expected != actual {
			t.Errorf("Expected false but got true")
		}
	})

	t.Run("invalid2", func(subtest *testing.T) {
		expected := false
		actual := parChecker((")("))

		if expected != actual {
			t.Errorf("Expected false but got true")
		}
	})
}
