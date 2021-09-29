package array

import (
	"testing"
)

func TestPush(t *testing.T) {
	t.Run("push true", func(subtest *testing.T) {
		test := make([]interface{}, 5, 10)
		copy(test, []interface{}{1, 2, 3, 4, 5})
		expected := []interface{}{1, 2, 3, 4, 5, 6}

		test, err := Push(test, 6)
		if err != nil {
			subtest.Fatal(err)
		}

		if len(expected) != len(test) {
			subtest.Fatalf("Length of expected (%d) not equal to length of actual (%d)",
				expected, test)
		}

		for i, v := range expected {
			if v != test[i] {
				subtest.Fatalf("Expected %d but got %d", expected, test)
			}
		}
	})

	t.Run("push false", func(subtest *testing.T) {
		test := make([]interface{}, 5)
		_, err := Push(test, 6)

		if err == nil {
			subtest.Fatal("Expected error: Push exceeds slice capacity")
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("pop true", func(subtest *testing.T) {
		test := make([]interface{}, 5, 10)
		copy(test, []interface{}{1, 2, 3, 4, 5})
		expected := []interface{}{1, 2, 3, 4}

		_, test = Pop(test)

		if len(expected) != len(test) {
			subtest.Fatalf("Length of expected (%d) not equal to length of actual (%d)",
				expected, test)
		}

		for i, v := range expected {
			if v != test[i] {
				subtest.Fatalf("Expected %d but got %d", expected, test)
			}
		}
	})

	t.Run("pop empty", func(subtest *testing.T) {
		var test []interface{}
		val, test := Pop(test)

		if val != nil {
			subtest.Fatalf("Expected nil element")
		}

		if test != nil {
			subtest.Fatalf("Expected nil slice")
		}
	})
}

func TestPushFront(t *testing.T) {
	t.Run("pushfront true", func(subtest *testing.T) {
		test := make([]interface{}, 5, 10)
		copy(test, []interface{}{1, 2, 3, 4, 5})
		expected := []interface{}{0, 1, 2, 3, 4, 5}

		test, err := PushFront(test, 0)
		if err != nil {
			subtest.Fatal(err)
		}

		if len(expected) != len(test) {
			subtest.Fatalf("Length of expected (%d) not equal to length of actual (%d)",
				expected, test)
		}

		for i, v := range expected {
			if v != test[i] {
				subtest.Fatalf("Expected %d but got %d", expected, test)
			}
		}
	})

	t.Run("pushfront false", func(subtest *testing.T) {
		test := make([]interface{}, 5)
		_, err := PushFront(test, 0)

		if err == nil {
			subtest.Fatal("Expected error: Push exceeds slice capacity")
		}
	})
}

func TestPopFront(t *testing.T) {
	t.Run("pop true", func(subtest *testing.T) {
		test := make([]interface{}, 5, 10)
		copy(test, []interface{}{1, 2, 3, 4, 5})
		expected := []interface{}{2, 3, 4, 5}
		_, test = PopFront(test)

		if len(expected) != len(test) {
			subtest.Fatalf("Length of expected (%d) not equal to length of actual (%d)",
				expected, test)
		}

		for i, v := range expected {
			if v != test[i] {
				subtest.Fatalf("Expected %d but got %d", expected, test)
			}
		}
	})

	t.Run("pop empty", func(subtest *testing.T) {
		var test []interface{}
		val, test := PopFront(test)

		if val != nil {
			subtest.Fatalf("Expected nil element")
		}

		if test != nil {
			subtest.Fatalf("Expected nil slice")
		}
	})
}

func TestInsert(t *testing.T) {
	t.Run("insert true", func(subtest *testing.T) {
		test := make([]interface{}, 5, 10)
		copy(test, []interface{}{1, 2, 3, 4, 5})
		expected := []interface{}{1, 2, 3, -7, 4, 5}

		test, err := Insert(test, 3, -7)
		if err != nil {
			subtest.Fatal(err)
		}

		if len(expected) != len(test) {
			subtest.Fatalf("Length of expected (%d) not equal to length of actual (%d)",
				expected, test)
		}

		for i, v := range expected {
			if v != test[i] {
				subtest.Fatalf("Expected %d but got %d", expected, test)
			}
		}
	})

	t.Run("insert false", func(subtest *testing.T) {
		test := make([]interface{}, 5)
		_, err := Insert(test, 3, -7)

		if err == nil {
			subtest.Fatal("Expected error: Insert exceeds slice capacity")
		}
	})
}

func TestRemove(t *testing.T) {
	t.Run("pop true", func(subtest *testing.T) {
		test := make([]interface{}, 5, 10)
		copy(test, []interface{}{1, 2, 3, 4, 5})
		expected := []interface{}{1, 2, 4, 5}

		test = Remove(test, 2)

		if len(expected) != len(test) {
			subtest.Fatalf("Length of expected (%d) not equal to length of actual (%d)",
				expected, test)
		}

		for i, v := range expected {
			if v != test[i] {
				subtest.Fatalf("Expected %d but got %d", expected, test)
			}
		}
	})

	t.Run("pop empty", func(subtest *testing.T) {
		var test []interface{}
		val, test := PopFront(test)

		if val != nil {
			subtest.Fatalf("Expected nil element")
		}

		if test != nil {
			subtest.Fatalf("Expected nil slice")
		}
	})
}
