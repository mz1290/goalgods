package main

// https://www.youtube.com/watch?v=5o-kdjv7FD0
func numways(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return numways(n-1) + numways(n-2)
}
