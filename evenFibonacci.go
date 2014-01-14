package main

import (
	"fmt"
)

func main() {
	// Find the sum of the Fibonnaci numbers whose value does not exceed 4mil
	total := fib(2, 1, 2)
	fmt.Printf("Total = %d\n", total)
}

func fib(total int, prev1 int, prev2 int) int {
	next := prev1 + prev2
	if next <= 4000000 {
		if next%2 == 0 {
			total += next
		}
		return fib(total, prev2, next)
	} else {
		return total
	}
}
