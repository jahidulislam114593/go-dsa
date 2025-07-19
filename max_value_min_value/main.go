package main

import "fmt"

func main() {
	nums := []int{2, 5, 123, 55, 3, 26, 443, 3}
	if len(nums) == 0 {
		fmt.Println("Array is empty")
		return
	}

	mx, mn := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > mx {
			mx = nums[i]
		} else if nums[i] < mn {
			mn = nums[i]
		}
	}

	fmt.Println("Max value, ", mx)
	fmt.Println("Min value, ", mn)
}
