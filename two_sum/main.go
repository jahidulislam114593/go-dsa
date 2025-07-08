package main

import (
	"fmt"
)

func twoSum(arr []int, k int) (int, int) {
	// O(n ^ 2)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == k {
				return i, j
			}
		}
	}
	return -1, -1
}

func main() {
	nums := [4]int{2, 7, 11, 15}
	target := 9

	i, j := twoSum(nums[:], target)
	fmt.Println(i, j)
}
