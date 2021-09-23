package main

import (
	"testing"
)

func TestNumWays(t *testing.T) {
	expected := 5
	actual := numways(4)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
