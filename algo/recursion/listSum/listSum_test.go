package main

import "testing"

func TestListSum(t *testing.T) {
	test := []int{1, 3, 5, 7, 9}
	expected := 25
	actual := listSum(test)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
