package main

import (
	"fmt"
)

func main() {
	solved := false
	num := 19
	for !solved {
		err := false
		num++
		for x := 2; x <= 20 && !err; x++ {
			if num%x != 0 {
				err = true
			}
		}
		solved = !err
	}

	fmt.Println("Answer:", num)
}
