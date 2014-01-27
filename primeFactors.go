package main

import (
	"fmt"
)

func main() {
	for {
		var num int
		fmt.Scanf("%d", &num)
		fax := factors(num)
		fmt.Println("Factors:", fax)
		fmt.Println("Largest:", greatest(fax))
	}
}

func factors(parent int) []int {
	i := 1
	found := false

	// Loop 2..parent, check if i evenly divides parent
	for !found {
		i++
		if parent%i == 0 {
			found = true
		}
	}

	// If i == parent, parent is prime, so return that
	// Otherwise find more factors
	if i == parent {
		if i == 2 {
			return make([]int, 0)
		} else {
			return []int{parent}
		}
	} else {
		return append(factors(i), factors(parent/i)...)
	}
}

func greatest(vals []int) int {
	tmp := vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] > tmp {
			tmp = vals[i]
		}
	}
	return tmp
}
