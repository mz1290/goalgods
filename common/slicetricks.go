// https://go.dev/blog/slices-intro
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

func InsertLaziest(slice []interface{}, idx int, val interface{}) []interface{} {
	return append(slice[:idx], append([]interface{}{val}, slice[idx:]...)...)
}

func InsertLazy(slice []interface{}, idx int, val interface{}) []interface{} {
	slice = append(slice, nil)
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = val
	return slice
}

func Insert(slice []interface{}, idx int, vector ...interface{}) []interface{} {
	if n := len(slice) + len(vector); n <= cap(slice) {
		slice2 := slice[:n]
		copy(slice2[idx+len(vector):], slice[idx:])
		copy(slice2[idx:], vector)
		return slice2
	}
	slice2 := make([]interface{}, len(slice)+len(vector))
	copy(slice2, slice[:idx])
	copy(slice2[idx:], vector)
	copy(slice2[idx+len(vector):], slice[idx:])
	return slice2
}

func Copy(slice []interface{}) []interface{} {
	newSlice := make([]interface{}, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func Delete(slice []interface{}, idx int) []interface{} {
	// Shift elements up to perform delete
	copy(slice[idx:], slice[idx+1:])

	// Set last copied, initialized, item to zero value for gc
	slice[len(slice)-1] = nil

	// return new slice excluding zeroed value
	return slice[:len(slice)-1]
}

// Reverse contents of slice in-place
func Reverse(slice []interface{}) {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
}

// Filter returns a new slice holding only
// the elements of s that satisfy fn()
func FilterNew(s []interface{}, fn func(interface{}) bool) []interface{} {
	var p []interface{}
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

func FilterInPlace(slice []interface{}, fn func(interface{}) bool) []interface{} {
	n := 0
	for _, x := range slice {
		if fn(x) {
			slice[n] = x
			n++
		}
	}
	return slice[:n]
}
