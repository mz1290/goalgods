package main

import "fmt"

type item struct {
	value interface{} //value as interface type to hold any data type
	next  *item
}

type Stack struct {
	top  *item
	size int
}

// Check if the stack is emtpy.
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Add an item to the stack.
func (s *Stack) Push(value interface{}) {
	s.top = &item{
		value: value,
		next:  s.top,
	}

	s.size++
}

// Remove item from the stack. If empty, return nil.
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	// Store current top in temp variable
	value := s.top.value

	// Update stack to point at next item for new top
	s.top = s.top.next

	// Decrement stack size
	s.size--

	return value
}

// Get the value of the top item in the stack.
func (s *Stack) Peek() interface{} {
	return s.top.value
}

// Get the number of items in the stack.
func (s *Stack) Size() int {
	return s.size
}

// Print out the values in stack.
func (s *Stack) Print() {
	current := s.top
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func main() {
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
