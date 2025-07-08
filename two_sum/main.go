package main

import (
	"fmt"
	"sort"
)

type pair struct {
	val int
	idx int
}

func twoSum(arr []int, k int) (int, int) {

	pairs := make([]pair, len(arr))

	for i, v := range arr {
		pairs[i] = pair{
			val: v,
			idx: i,
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].val < pairs[j].val
	})

	st := 0
	en := len(arr) - 1

	for st < en {
		sum := pairs[st].val + pairs[en].val

		if sum == k {
			return pairs[st].idx, pairs[en].idx
		} else if sum > k {
			en--
		} else {
			st++
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
