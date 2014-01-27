package main

import (
	"fmt"
	"math"
)

func main() {
	// max := 20
	max := 2000000
	sum := 2
	i := 3
	for i < max {
		if isPrime(i) {
			sum += i
		}
		i += 2
	}
	fmt.Println("Sum:", sum)
}

func isPrime(num int) bool {
	max := int(math.Ceil(math.Sqrt(float64(num))))
	for i := 3; i <= max; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
