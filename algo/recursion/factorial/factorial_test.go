package main

import "testing"

func TestFactorial(t *testing.T) {
	expected := 5040
	actual := factorial(7)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
