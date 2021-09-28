/*
	mergeSortedArrays([0,3,4,31], [4,6,30])
	[0,3,4,4,6,30,31]
*/
package main

import "fmt"

func main() {
	a1 := []int{0, 3, 4, 31}
	a2 := []int{4, 6, 30}
	fmt.Println(mergeSortedArray(a1, a2))
}

func mergeSortedArray(a1, a2 []int) []int {
	if a1 == nil {
		return a1
	}

	if a2 == nil {
		return a2
	}

	p1 := 0
	p2 := 0
	p3 := 0
	merged := make([]int, len(a1)+len(a2))
	for {
		var val int
		if a1[p1] < a2[p2] {
			val = a1[p1]
			p1++
		} else if a1[p1] > a2[p2] {
			val = a2[p2]
			p2++
		} else {
			val = a1[p1]
			p1++
		}

		merged[p3] = val
		p3++

		if p1 == len(a1) || p2 == len(a2) {
			break
		}
	}

	if p1 != len(a1) {
		for i := p1; i < len(a1); i++ {
			merged[p3] = a1[i]
			p3++
		}
	} else if p2 != len(a2) {
		for i := p2; i < len(a2); i++ {
			merged[p3] = a2[i]
			p3++
		}
	}

	return merged
}
