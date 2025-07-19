package main

import "fmt"

func mejorityElement(arr []int) int {
	n := len(arr)
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if arr[i] == arr[j] {
				cnt++
			}
		}
		if cnt > n/2 {
			return arr[i]
		}
	}
	return -1
}
func main() {
	nums := []int{1, 2, 2, 1, 1}
	mejority := mejorityElement(nums)
	fmt.Println(mejority)
}
