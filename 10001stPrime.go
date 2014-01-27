package main

import (
	"fmt"
)

func main() {
	num := 1
	count := 0
	for count < 10001 {
		num++
		if isPrime(num) {
			fmt.Println(num)
			count++
		}
	}
	fmt.Println("10001th Prime:", num)
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
