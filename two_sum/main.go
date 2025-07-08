package main

import (
	"fmt"
)

func twoSum(arr []int, k int) (int, int) {

	seen := make(map[int]int)

	for i, num := range arr {
		complement := k - num
		if idx, found := seen[complement]; found {
			return idx, i
		}
		seen[num] = i
	}

	return -1, -1

}

func main() {
	nums := [4]int{2, 7, 11, 15}
	target := 9

	i, j := twoSum(nums[:], target)
	fmt.Println(i, j)
}
