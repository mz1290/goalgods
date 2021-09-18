package stack

import (
	"fmt"
	"testing"
)

func TestStackImplementation(t *testing.T) {
	s := new(Stack)

	fmt.Println(s.IsEmpty())
	s.Push(4)
	s.Push("dog")
	fmt.Println(s.Peek())
	s.Push(true)
	fmt.Println(s.Size())
	fmt.Println(s.IsEmpty())
	s.Push(8.4)
	fmt.Printf("\nContents:\n")
	s.Print()
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Size())
}
