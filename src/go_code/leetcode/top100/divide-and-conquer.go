package main

import (
	"container/heap"
	"fmt"
)

//divide-and-conquer
func partition(sub_nums []int) int { //将数组分为两部分，左边的值小于右边的，返回中间位置
	mid, i, j := sub_nums[0], 0, len(sub_nums)-1
	for i < j {
		for mid <= sub_nums[j] && i < j {
			j--
		}
		if mid >= sub_nums[j] {
			sub_nums[i], sub_nums[j] = sub_nums[j], sub_nums[i]
		}
		for mid >= sub_nums[i] && i < j {
			i++
		}
		if mid <= sub_nums[i] {
			sub_nums[i], sub_nums[j] = sub_nums[j], sub_nums[i]
		}
	}
	return i
}

//使用快速排序的思想进行实现
func FindKthLargest(nums []int, k int) int {
	sub_k := len(nums) - k
	var dfs func(sub_nums []int, k int)
	dfs = func(sub_nums []int, k int) {
		pos := partition(sub_nums)
		if pos == k {
			return
		}
		if pos > k {
			dfs(sub_nums[:pos], k)
		}
		if pos < k {
			dfs(sub_nums[pos+1:], k-pos-1)
		}
	}
	dfs(nums, sub_k)
	return nums[sub_k]
}

//使用堆思想进行实现
type Heap_int []int

func (H Heap_int) Len() int {
	return len(H)
}
func (H Heap_int) Swap(i, j int) {
	H[i], H[j] = H[j], H[i]
}
func (H Heap_int) Less(i, j int) bool {
	return H[i] > H[j]
}

// func (H *Heap_int) Push(h interface{}) {
// 	*H = append(*H, h.(int))
// }
func (H Heap_int) Push(h interface{}) {
	H = append(H, h.(int))
	fmt.Println(H)
}

// func (H *Heap_int) Pop() (x interface{}) {
// 	n := len(*H) ///为什么用*？？
// 	x = (*H)[n-1]
// 	*H = (*H)[:n-1]
// 	return
// }
func (H Heap_int) Pop() (x interface{}) {
	n := len(H) ///为什么用*？？
	x = (H)[n-1]
	H = (H)[:n-1]
	return
}
func findKthLargest(nums []int, k int) int {
	// hp := &Heap_int{}
	hp := Heap_int{}

	for i := range nums {
		hp = append(hp, nums[i])
	}
	heap.Init(hp)
	for i := 0; i < k-1; i++ {
		heap.Pop(hp)
	}
	return heap.Pop(hp).(int)
}
func getKthelement(num1, num2 []int, k int) int {
	for {
		len1, len2 := len(num1), len(num2)
		if len1 == 0 {
			return num2[k-1]
		}
		if len2 == 0 {
			return num1[k-1]
		}
		if k == 1 {
			return min(num1[0], num2[0])
		}
		sub_k := k / 2
		index1 := min(len1, sub_k)
		index2 := min(len2, sub_k)
		if num1[index1-1] < num2[index2-1] {
			k -= index1
			num1 = num1[index1:]
		} else {
			k -= index2
			num2 = num2[index2:]
		}
	}
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	k := (len1 + len2) / 2

	if (len1+len2)%2 == 0 {
		return float64(getKthelement(nums1, nums2, k)+getKthelement(nums1, nums2, k+1)) / 2.0
	} else {
		return float64(getKthelement(nums1, nums2, k+1))
	}
}
