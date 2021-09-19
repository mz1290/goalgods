package queue

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	qp := new(Queue)
	expected := "Hello world"

	qp.Enqueue(expected)
	actual := qp.Peek()

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestDequeue(t *testing.T) {
	qp := new(Queue)
	expected := "Hello World"

	qp.Enqueue(expected)
	actual := qp.Dequeue()

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}

	if qp.Dequeue() != nil {
		t.Errorf("Expected queue to be empty (nil)")
	}
}
