package common

import (
	"testing"
)

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

func TestInsert(t *testing.T) {
	test := []interface{}{1, 2, 3, 4, 5}
	expected := 42

	t.Run("insert laziest", func(subtest *testing.T) {
		temp := InsertLaziest(test, 2, 42)
		actual := temp[2]

		if expected != actual {
			subtest.Errorf("Expected %d but got %d", expected, actual)
		}
	})

	t.Run("insert lazy", func(subtest *testing.T) {
		temp := InsertLazy(test, 2, 42)
		actual := temp[2]

		if expected != actual {
			subtest.Errorf("Expected %d but got %d", expected, actual)
		}
	})

	t.Run("insert", func(subtest *testing.T) {
		var t1 interface{} = 42
		var t2 interface{} = 8
		var t3 interface{} = 9
		temp := Insert(test, 2, t1, t2, t3)
		actual := temp[2]

		if expected != actual {
			subtest.Errorf("Expected %d but got %d", expected, actual)
		}
	})
}
