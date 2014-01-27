package main

import (
	"fmt"
	"strconv"
)

func main() {
	pallies := make([]int, 1)
	for i := 999; i > 900; i-- {
		for j := 999; j > 900; j-- {
			prod := i * j
			if isPalindrome(prod) == true {
				pallies = append(pallies, prod)
			}
		}
	}
	fmt.Println("Palindromes found:")
	fmt.Println(pallies)
	fmt.Println()
	fmt.Println("Greatest:", greatest(pallies))
}

func isPalindrome(num int) bool {
	str := strconv.Itoa(num)
	rev := reverse(str)
	return str == rev
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

func reverse(str string) string {
	tmp := ""
	for i := len(str) - 1; i >= 0; i-- {
		tmp += string(str[i])
	}
	return tmp
}
