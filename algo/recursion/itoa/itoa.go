package main

import "fmt"

func itoa(n, base int) string {
	if n < base {
		return fmt.Sprint(n)
	}

	return itoa(n/base, base) + fmt.Sprint(n%base)
}
