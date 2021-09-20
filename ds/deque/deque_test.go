package deque

import (
	"fmt"
	"testing"
)

func TestDequeAddFront(t *testing.T) {
	dp := NewDeque()
	dp.PushFront("World")
	dp.PushFront("Hello")
	fmt.Println(dp.PopFront())
	fmt.Println(dp.PopFront())
}

func TestDequeAddRear(t *testing.T) {
	dp := NewDeque()
	dp.PushBack("World")
	dp.PushBack("Hello")
	fmt.Println(dp.PopBack())
	fmt.Println(dp.PopBack())
}
