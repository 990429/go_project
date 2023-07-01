package main

import (
	"fmt"
	"math/rand"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxSubArray(nums []int) int {
	result := nums[0]
	sum := nums[0]

	for index := 1; index < len(nums); index++ {
		if sum < 0 {
			sum = nums[index]
		} else {
			sum += nums[index]
		}
		result = max(result, sum)
	}
	return result
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func getElement(nums1, nums2 []int, index int) (ans int) {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index1]
		}
		if index2 == len(nums2) {
			return nums1[index-index2]
		}

	}
	return
}
func findMedianSortedArrays(nums1 []int, nums2 []int) (ans float64) {

	return
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(nums []int) *ListNode {
	head := &ListNode{}
	term := head
	for _, value := range nums {
		newnode := new(ListNode)
		newnode.Val = value
		term.Next = newnode
		term = newnode
	}
	return head
}
func quickSort(node []*ListNode) {
	if len(node) == 0 {
		return
	}
	index := rand.Intn(len(node))
	base := node[index]
	i, j := 0, len(node)-1
	for i < j {
		for node[j].Val >= base.Val && i < j {
			j--
		}
		if i >= j {
			break
		}
		node[i], node[j] = node[j], node[i]
		for node[i].Val <= base.Val && i < j {
			i++
		}
		if i >= j {
			break
		}
		node[i], node[j] = node[j], node[i]
	}
	quickSort(node[:i])
	quickSort(node[i+1:])
}
func merge(L1, L2 *ListNode) (ans [2]*ListNode) {
	L := &ListNode{}
	term := L
	term1, term2 := L1, L2

	for term1 != nil && term2 != nil {
		if term1.Val > term2.Val {
			term.Next = term2
			term = term2
			term2 = term2.Next
		} else {
			term.Next = term1
			term = term1
			term1 = term1.Next
		}
	}
	if term1 != nil {
		term.Next = term1
	}
	if term2 != nil {
		term.Next = term2
	}
	for term1 != nil {
		ans[1] = term1
		term1 = term1.Next
	}
	for term2 != nil {
		ans[1] = term2
		term2 = term2.Next
	}
	ans[0] = L.Next
	return
}
func sortList(head *ListNode) *ListNode {
	term := head
	length := 0
	node := []*ListNode{}
	for term != nil {
		length++
		node = append(node, term)
		term = term.Next
	}
	if length <= 1 {
		return head
	}
	result := [2]*ListNode{}
	for sub_len := 1; sub_len < length; sub_len *= 2 {
		for loop := 0; loop < length/sub_len; loop += 2 {
			index1 := loop * sub_len
			index2 := (loop + 1) * sub_len
			L1 := node[index1]
			node[index2-1].Next = nil
			var L2 *ListNode
			if index2 >= length {
				L2 = nil
			} else {
				L2 = node[index2]
				term := L2
				for i := 0; i < sub_len && term != nil; i++ {
					term = term.Next
				}
				term = nil
			}
			result = merge(L1, L2)
		}
	}
	return result[0]
}
func merge_num(num1, num2, nums []int) {
	i, j, k := 0, 0, 0
	for ; i < len(num1) && j < len(num2); k++ {
		if num1[i] < num2[j] {
			nums[k] = num1[i]
			i++
		} else {
			nums[k] = num2[j]
			j++
		}
	}
	for ; i < len(num1); k++ {
		nums[k] = num1[i]
		i++
	}
	for ; j < len(num2); k++ {
		nums[k] = num2[j]
		j++
	}
}

func mergeSort(nums []int) {
	if len(nums) > 1 {
		length := len(nums) / 2
		num1 := make([]int, length)
		num2 := make([]int, len(nums)-length)
		copy(num1, nums[:length])
		copy(num2, nums[length:])
		mergeSort(num1)
		mergeSort(num2)
		merge_num(num1, num2, nums)
	}
}
func wiggleSort(nums []int) {
	mergeSort(nums)
	//length := (len(nums)+1) / 2
	temp := make([]int, len(nums))
	copy(temp, nums)
	if len(nums)%2 == 0 {

	}
	// offset := length
	// if len(nums)%2 != 0 {
	// 	offset++
	// }
	// for index := 0; index < length; index++ {
	// 	nums[index*2] = temp[index]
	// 	nums[index*2+1] = temp[index+offset]
	// }
	// if len(nums)%2 != 0 {
	// 	nums[len(nums)-1] = temp[length]
	// }
}
func main() {
	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
	fmt.Println(nums)
}
