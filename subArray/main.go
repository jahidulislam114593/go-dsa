package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	n := len(nums)
	maxSum := nums[0]
	for i := 0; i < n; i++ {
		currentSum := 0
		for j := i; j < n; j++ {
			currentSum += nums[j]
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
	}
	fmt.Println(maxSum)
}
