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

func TestCopy(t *testing.T) {

	t.Run("copy ints", func(subtest *testing.T) {
		test := []interface{}{1, 2, 3, 4, 5}
		actual := Copy(test)
		expected := []interface{}{1, 2, 3, 4, 5}

		if len(expected) != len(actual) {
			subtest.Fatalf("Length of expected (%d) not equal to length of actual (%d)",
				expected, actual)
		}

		for i, v := range expected {
			if v != actual[i] {
				subtest.Fatalf("Expected %d but got %d", expected, actual)
			}
		}
	})

	t.Run("copy strings", func(subtest *testing.T) {
		test := []interface{}{"h", "e", "l", "l", "o"}
		actual := Copy(test)
		expected := []interface{}{"h", "e", "l", "l", "o"}

		if len(expected) != len(actual) {
			subtest.Fatalf("Length of expected (%d) not equal to length of actual (%d)",
				expected, actual)
		}

		for i, v := range expected {
			if v != actual[i] {
				subtest.Fatalf("Expected %d but got %d", expected, actual)
			}
		}
	})
}

func TestDelete(t *testing.T) {
	test := []interface{}{1, 2, 3, 4, 5}
	actual := Delete(test, 2)
	expected := []interface{}{1, 2, 4, 5}

	if len(expected) != len(actual) {
		t.Fatalf("Length of expected (%d) not equal to length of actual (%d)",
			expected, actual)
	}

	for i, v := range expected {
		if v != actual[i] {
			t.Fatalf("Expected %d but got %d", expected, actual)
		}
	}
}
