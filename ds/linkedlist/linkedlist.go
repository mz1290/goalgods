package linkedlist

import "fmt"

// A predefined function that takes two data types and returns true if equal.
// False if not.
type compare func(interface{}, interface{}) bool

type item struct {
	value     interface{}
	next      *item
	prev      *item
	isEqualTo compare
}

type LinkedList struct {
	frontSentinel *item
	backSentinel  *item
	size          int
}

func NewLinkedList() *LinkedList {
	lp := &LinkedList{
		frontSentinel: new(item),
		backSentinel:  new(item),
		size:          0,
	}

	// Provision the front sentinel
	lp.frontSentinel.next = lp.backSentinel
	lp.frontSentinel.prev = nil

	// Provision the back sentinel
	lp.backSentinel.next = nil
	lp.backSentinel.prev = lp.frontSentinel

	return lp
}

// Insert a new item with value into list BEFORE the provided item.
func (l *LinkedList) addItemBefore(node *item, value interface{}, f compare) {
	// Provision new item
	newItem := &item{
		value:     value,
		next:      node,
		prev:      node.prev,
		isEqualTo: f,
	}

	// Update border nodes
	node.prev.next = newItem
	node.prev = newItem

	l.size++
}

// General function to remove an item from list.
func (l *LinkedList) removeItem(node *item) {
	// Update border items.
	node.prev.next = node.next
	node.next.prev = node.prev

	// Tell Go to gc node (Need to research more)
	node = nil

	// Reduce size
	l.size--
}

// Adds a new item to the rear of the list. It needs the item and returns
// nothing.
func (l *LinkedList) pushBack(value interface{}, f compare) {
	l.addItemBefore(l.backSentinel, value, f)
}

// Removes the rear item from the list. It needs no parameters and returns the
// item. The list is modified.
func (l *LinkedList) popBack() interface{} {
	if l.IsEmpty() {
		return nil
	}

	// Save front value.
	value := l.backSentinel.prev.value

	//Remove back item.
	l.removeItem(l.backSentinel.prev)

	return value
}

// Check if the list is emtpy.
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// Get the number of items in the list.
func (l *LinkedList) Size() int {
	return l.size
}

// Add item with given value to linked list.
func (l *LinkedList) Add(value interface{}, f compare) {
	l.pushBack(value, f)
}

// Removes first occurence of the value from the list. Returns an error if the
// item does not exist in the list.
func (l *LinkedList) Remove(value interface{}) error {
	if l.IsEmpty() {
		return fmt.Errorf("list contains no values")
	}

	var cur *item = l.frontSentinel.next
	for cur != l.backSentinel {
		if cur.isEqualTo(cur.value, value) {
			l.removeItem(cur)
			return nil
		}

		cur = cur.next
	}

	return fmt.Errorf("list does not contain %v", value)
}

// Returns true if value exists in list. False if not. O(n).
func (l *LinkedList) Contains(value interface{}) bool {
	if l.IsEmpty() {
		return false
	}

	var cur *item = l.frontSentinel.next
	for cur != l.backSentinel {
		if cur.isEqualTo(value, cur.value) {
			return true
		}

		// Advance to next node
		cur = cur.next
	}

	return false
}

// Returns the index postion of the first value occurence.
func (l *LinkedList) Index(value interface{}) (int, error) {
	if l.IsEmpty() {
		return -1, fmt.Errorf("list is empty")
	}

	var cur *item = l.frontSentinel.next
	idx := 0
	for cur != l.backSentinel {
		if cur.isEqualTo(value, cur.value) {
			return idx, nil
		}

		// Advance to next node
		cur = cur.next
		idx++
	}

	return -1, fmt.Errorf("%v does not exist in list", value)
}

// Insert item at specific position in list.
func (l *LinkedList) Insert(idx int, value interface{}, f compare) error {
	if idx > (l.size - 1) {
		return fmt.Errorf("index (%d) exceeds list size (%d)", idx, l.size)
	}

	// Advance to specified node.
	var node *item = l.frontSentinel.next
	i := 0
	for i < idx {
		node = node.next
		i++
	}

	// Insert node into list at index
	l.addItemBefore(node, value, f)

	return nil
}

// Removes and returns last item in list.
func (l *LinkedList) Pop() interface{} {
	return l.popBack()
}

func (l *LinkedList) PopIdx(idx int) (interface{}, error) {
	if idx > (l.size - 1) {
		return nil, fmt.Errorf("index (%d) exceeds list size (%d)", idx,
			l.size)
	}

	// Advance to specified node.
	var node *item = l.frontSentinel.next
	i := 0
	for i < idx {
		node = node.next
		i++
	}

	// Save node value
	value := node.value

	// Insert node into list at index
	l.removeItem(node)

	return value, nil
}
