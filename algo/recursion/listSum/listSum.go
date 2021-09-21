package main

/*
 * THE THREE LAWS OF RECURSION
 * A recursive algorithm must have a base case.
 * A recursive algorithm must change its state and move toward the base case.
 * A recursive algorithm must call itself, recursively.
 */

func listSum(nums []int) int {
	// Base case (law #1)
	if len(nums) == 1 {
		return nums[0]
	}

	// Recursive case
	// nums[1:] -> changing state, making the list smaller, moving towards base (law #2)
	// listSum(xx) -> algorithm calls itself, recursively (law #3)
	return nums[0] + listSum(nums[1:])
}
