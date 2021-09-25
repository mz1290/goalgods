// https://github.com/golang/go/wiki/SliceTricks
package common

// Push to the end of a slice
func Push(slice []interface{}, val interface{}) []interface{} {
	return append(slice, val)
}

// Pop item from end of slice
func Pop(slice []interface{}) (interface{}, []interface{}) {
	return slice[len(slice)-1], slice[:len(slice)-1]
}

// Push value onto front of slice. Creates a new slice with val as first
// element and appends slice.
func PushFront(slice []interface{}, val interface{}) []interface{} {
	return append([]interface{}{val}, slice...)
}

// Pop item from front of slice. Returns value at slice[0] and returns new
// slice with previous index 1 as new index 0.
func PopFront(slice []interface{}) (interface{}, []interface{}) {
	return slice[0], slice[1:]
}
