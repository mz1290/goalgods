package deque

type item struct {
	value interface{}
	next  *item
	prev  *item
}

type Deque struct {
	frontSentinel *item
	backSentinel  *item
	size          int
}

// FS -> BS
func NewDeque() *Deque {
	dp := &Deque{
		frontSentinel: new(item),
		backSentinel:  new(item),
		size:          0,
	}

	// Provision the front sentinel
	dp.frontSentinel.next = dp.backSentinel
	dp.frontSentinel.prev = nil

	// Provision the back sentinel
	dp.backSentinel.next = nil
	dp.backSentinel.prev = dp.frontSentinel

	return dp
}

// Insert a new item with value into deque BEFORE the provided item.
func (d *Deque) addItemBefore(node *item, value interface{}) {
	// Provision new item
	newItem := &item{
		value: value,
		next:  node,
		prev:  node.prev,
	}

	// Update border nodes
	node.prev.next = newItem
	node.prev = newItem

	d.size++
}

// General function to remove an item from deque.
func (d *Deque) removeItem(node *item) {
	// Update border items.
	node.prev.next = node.next
	node.next.prev = node.prev

	// Tell Go to gc node
	node = nil

	// Reduce size
	d.size--
}

// Adds a new item to the front of the deque. It needs the item and returns
// nothing.
func (d *Deque) PushFront(value interface{}) {
	d.addItemBefore(d.frontSentinel.next, value)
}

// Adds a new item to the rear of the deque. It needs the item and returns
// nothing.
func (d *Deque) PushBack(value interface{}) {
	d.addItemBefore(d.backSentinel, value)
}

// Removes the front item from the deque. It needs no parameters and returns
// the item. The deque is modified.
func (d *Deque) PopFront() interface{} {
	if d.IsEmpty() {
		return nil
	}

	// Save front value.
	value := d.frontSentinel.next.value

	// Remove front item.
	d.removeItem(d.frontSentinel.next)

	return value
}

// Removes the rear item from the deque. It needs no parameters and returns the
// item. The deque is modified.
func (d *Deque) PopBack() interface{} {
	if d.IsEmpty() {
		return nil
	}

	// Save front value.
	value := d.backSentinel.prev.value

	//Remove back item.
	d.removeItem(d.backSentinel.prev)

	return value
}

// Get the value at the front of deque. Do not remove.
func (d *Deque) Front() interface{} {
	if d.IsEmpty() {
		return nil
	}

	return d.frontSentinel.next.value
}

// Get the value at the back of deque. Do not remove.
func (d *Deque) Back() interface{} {
	if d.IsEmpty() {
		return nil
	}

	return d.backSentinel.prev.value
}

// Check if the Deque is emtpy.
func (d *Deque) IsEmpty() bool {
	return d.size == 0
}

// Get the number of items in the deque.
func (d *Deque) Size() int {
	return d.size
}
