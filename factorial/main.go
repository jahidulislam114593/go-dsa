package main

import "fmt"

func fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(5))
}
