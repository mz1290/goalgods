package common

import "testing"

func TestPush(t *testing.T) {
	test := []interface{}{1, 2, 3, 4, 5}
	test = Push(test, 6)
	expected := 6
	actual := test[5]

	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}

func TestPop(t *testing.T) {
	test := []interface{}{1, 2, 3, 4, 5}
	actual, _ := Pop(test)
	expected := 5

	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}

func TestPushFront(t *testing.T) {
	test := []interface{}{1, 2, 3, 4, 5}
	test = PushFront(test, 6)
	expected := 6
	actual := test[0]

	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}

func TestPopFront(t *testing.T) {
	test := []interface{}{1, 2, 3, 4, 5}
	actual, _ := PopFront(test)
	expected := 1

	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}
