package main

import (
	"fmt"
)

func main() {

	// Example on how to use slices
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	fmt.Println(append(slice1, slice2...))

	// Real code begin!!
	factors := []int{}
}

// This should be a recursive method for finding
// the factors of a number. The endpoints of the function
// should be the prime factors of the parent number
func factors(parent int) {

}
