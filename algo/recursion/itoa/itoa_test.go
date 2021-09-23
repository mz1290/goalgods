package main

import (
	"testing"
)

func TestItoaBase10(t *testing.T) {
	expected := "769"
	actual := itoa(769, 10)

	if expected != actual {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestItoaBase2(t *testing.T) {
	expected := "1100000001"
	actual := itoa(769, 2)

	if expected != actual {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestItoaBase16(t *testing.T) {
	expected := "301"
	actual := itoa(769, 16)

	if expected != actual {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}
