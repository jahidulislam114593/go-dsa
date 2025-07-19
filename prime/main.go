package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	num := 2

	if isPrime(num) {
		fmt.Println(num, "is a Prime number")
	} else {
		fmt.Println(num, "is not a Prime Number")
	}
}
