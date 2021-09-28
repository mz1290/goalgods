package main

import "fmt"

func main() {
	test := "Hi My name is Andrei"
	fmt.Println(reverse(test))
}

func reverse(test string) string {
	chars := make([]byte, len(test))

	// Iterate string in reverse, swapping back and front
	for i := len(test)/2 - 1; i >= 0; i-- {
		opposite := len(test) - 1 - i
		chars[i], chars[opposite] = test[opposite], test[i]
	}

	return string(chars)
}
