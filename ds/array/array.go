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

func PopFront(slice []interface{}) (interface{}, []interface{}) {
	if slice == nil {
		return nil, nil
	}

	return slice[0], slice[1:]
}

func Remove(slice []interface{}, idx int) []interface{} {
	if slice == nil {
		return nil
	}

	// Shift elements up to perform delete
	copy(slice[idx:], slice[idx+1:])

	// Set last copied, initialized, item to zero value for gc
	slice[len(slice)-1] = nil

	// return new slice excluding zeroed value
	return slice[:len(slice)-1]
}

// If value exists in sorted slice, returns the index. If not, returns -1.
func BinarySearch(slice []int, val int) int {
	low := 0
	high := len(slice) - 1

	for low <= high {
		mid := (low + high) / 2

		if slice[mid] == val {
			return mid
		} else if slice[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func getMinAndMax(slice []int) (int, int) {
	min := slice[0]
	max := slice[0]

	for _, value := range slice {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func getMinAndMaxIdx(slice []int) (int, int) {
	min := slice[0]
	max := slice[0]
	minIdx := 0
	maxIdx := 0

	for i, value := range slice {
		if value < min {
			min = value
			minIdx = i
		}

		if value > max {
			max = value
			maxIdx = i
		}
	}
	return minIdx, maxIdx
}
