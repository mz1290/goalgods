package main

import "testing"

func TestHotPotato(t *testing.T) {
	names := []string{"Bill", "David", "Susan", "Jane", "Kent", "Brad"}
	expected := "Susan"
	actual := hotpotato(names, 7)

	if expected != actual {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}
