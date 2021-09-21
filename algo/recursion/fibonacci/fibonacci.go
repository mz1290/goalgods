package main

// Returns the nth fibonacci number based on input
func fibonacci(num int) int {

	if num <= 1 {
		return num
	}

	return fibonacci(num-1) + fibonacci(num-2)
}
