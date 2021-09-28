package array

import "fmt"

func Push(slice []interface{}, val interface{}) ([]interface{}, error) {
	if cap(slice) < len(slice)+1 {
		return slice, fmt.Errorf("push exceeds slice capacity")
	}

	return append(slice, val), nil
}

// Pop item from end of slice
func Pop(slice []interface{}) (interface{}, []interface{}) {
	if slice == nil {
		return nil, nil
	}

	return slice[len(slice)-1], slice[:len(slice)-1]
}

func Insert(slice []interface{}, idx int, val interface{}) ([]interface{}, error) {
	if cap(slice) < len(slice)+1 {
		return slice, fmt.Errorf("insert exceeds slice capacity")
	}

	slice = append(slice, nil)
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = val

	return slice, nil
}

// Inserts item into list at specified index. Pops off top element as result.
func InsertAndPop(slice []int, idx int, val int) []int {
	copy(slice[idx+1:], slice[idx:len(slice)-1])
	slice[idx] = val
	return slice
}

// Slice tricks encourages us to create a new slice and append to it. Althrough
// this could be optimized in compiler, I'm sticking with the traditional, O(n)
// shift.
func PushFront(slice []interface{}, val interface{}) ([]interface{}, error) {
	if cap(slice) < len(slice)+1 {
		return slice, fmt.Errorf("push exceeds slice capacity")
	}

	inserted, _ := Insert(slice, 0, val)
	return inserted, nil
}
