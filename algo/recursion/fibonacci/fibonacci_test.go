package main

import "testing"

func TestFibonacci(t *testing.T) {
	expected := 5
	actual := fibonacci(5)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
