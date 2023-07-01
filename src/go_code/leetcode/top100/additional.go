package main

import (
	"math/rand"
	"time"
)

//additional
func add_partition(sub_nums []int) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(sub_nums))
	sub_nums[0], sub_nums[r] = sub_nums[r], sub_nums[0]
	mid := sub_nums[0]
	i, j := 0, len(sub_nums)-1
	for i < j {
		for i < j && sub_nums[j] >= mid {
			j--
		}
		if i < j && sub_nums[j] < mid {
			sub_nums[i], sub_nums[j] = sub_nums[j], sub_nums[i]
		}
		for i < j && sub_nums[i] <= mid {
			i++
		}
		if i < j && sub_nums[i] > mid {
			sub_nums[i], sub_nums[j] = sub_nums[j], sub_nums[i]
		}
	}
	return i
}
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pos := add_partition(nums)

	sortArray(nums[:pos])
	sortArray(nums[pos+1:])

	return nums
}
