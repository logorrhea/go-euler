package main

import (
	"fmt"
	"math"
)

func main() {
	divisorCount := 0
	i := 0
	for divisorCount < 500 {
		i++
		divisorCount = countDivisors(triangulate(i))
	}
	t := triangulate(i)
	fmt.Println(t)
	fmt.Println(countDivisors(t))
	fmt.Println(i)
}

func countDivisors(num int) int {
	divs := []int{1, num}
	max := int(math.Ceil(float64(num) / 2))
	for i := 2; i <= max; i++ {
		if num%i == 0 {
			divs = append(divs, i)
		}
	}

	for i := 2; i <= max; i++ {
		if num%i == 0 {

		}
	}

	return len(divs)
}

func triangulate(num int) int {
	v := 0
	for i := 1; i <= num; i++ {
		v += i
	}
	return v
}
