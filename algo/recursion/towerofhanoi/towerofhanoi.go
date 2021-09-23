package main

import "fmt"

// https://play.golang.org/p/T7K1VJEAFA
func movetower(n int, src, dst, temp string) {
	if n > 0 {
		movetower(n-1, src, temp, dst)
		movedisk(src, dst)
		movetower(n-1, temp, dst, src)
	}
}

func movedisk(src, dst string) {
	fmt.Printf("moving disk from %s to %s\n", src, dst)
}

func main() {
	movetower(3, "A", "B", "C")
}
