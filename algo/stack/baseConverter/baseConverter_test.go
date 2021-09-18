package main

import (
	"testing"
)

func TestBaseConverterValid(t *testing.T) {
	t.Run("base2", func(subtest *testing.T) {
		expected := "11001"
		actual := baseConverter(25, 2)

		if expected != actual {
			subtest.Errorf("Expected %s but got %s", expected, actual)
		}
	})

	t.Run("base16", func(subtest *testing.T) {
		expected := "19"
		actual := baseConverter(25, 16)

		if expected != actual {
			subtest.Errorf("Expected %s but got %s", expected, actual)
		}
	})
}

func TestBaseConverterInalid(t *testing.T) {
	t.Run("invalid_base", func(subtest *testing.T) {
		expected := ""
		actual := baseConverter(25, 3)

		if expected != actual {
			subtest.Errorf("Expected %s but got %s", expected, actual)
		}
	})
}
