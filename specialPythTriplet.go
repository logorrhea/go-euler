package main

import (
	"fmt"
	"math"
)

func main() {
	b := 4.0
	for b < 500.0 {
		b += 1
		a := 3.0
		for a < b {
			a += 1
			success, c := pyth(a, b)
			if success {
				if a+b+c == 1000 {
					fmt.Println("a:", a)
					fmt.Println("b:", b)
					fmt.Println("c:", c)
					fmt.Println("Product:", int(a*b*c))
					return
				}
			}
		}
	}
}

func pyth(a, b float64) (bool, float64) {
	c := math.Sqrt(a*a + b*b)
	success := math.Mod(c, 1) == 0
	return success, c
}
