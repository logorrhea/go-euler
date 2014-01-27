package main

import (
	"fmt"
)

func main() {
	sum := 0
	squaresum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		squaresum += i * i
	}
	fmt.Println("Sum^2:", sum*sum)
	fmt.Println("^2Sum:", squaresum)
	fmt.Println("Diff: ", (sum*sum)-squaresum)
}
