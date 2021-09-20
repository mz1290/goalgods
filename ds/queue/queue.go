package queue

import "fmt"

type item struct {
	value interface{} //value as interface type to hold any data type
	next  *item
}

type Queue struct {
	front *item
	rear  *item
	size  int
}

// Adds a new item to the rear of the queue. It needs the item and returns
// nothing.
func (q *Queue) Enqueue(value interface{}) {
	newItem := &item{
		value: value,
		next:  nil,
	}

	// If emtpy, the new item is front and rear.
	if q.IsEmpty() {
		q.front = newItem
		q.rear = newItem
	} else { // Add new item at end of queue
		q.rear.next = newItem
		q.rear = newItem
	}

	q.size++
}

// Removes the front item from the queue. It needs no parameters and returns
// the item. The queue is modified.
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	// Save front value.
	value := q.front.value

	// Move previous node up to new front.
	q.front = q.front.next

	// If new front is nil, then rear is also nil.
	if q.front == nil {
		q.rear = nil
	}

	// Decrement queue size
	q.size--

	return value
}

// Check if the queue is emtpy.
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Get the value at front of the queue. Do not remove.
func (q *Queue) Peek() interface{} {
	return q.front.value
}

// Get the number of items in the queue.
func (q *Queue) Size() int {
	return q.size
}

// Print out the values in queue.
func (q *Queue) Print() {
	current := q.front.next
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}
