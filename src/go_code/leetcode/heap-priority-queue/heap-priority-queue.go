package main

import (
	"container/heap"
	"fmt"
)

type hp struct { //目前实现小顶堆
	H [][2]int
}

func (h *hp) Push(v interface{}) {
	h.H = append(h.H, v.([2]int))
}
func (h *hp) Pop() interface{} {
	a := h.H
	v := a[len(h.H)-1]
	h.H = a[:len(h.H)-1]
	return v
}
func (h *hp) Less(i, j int) bool {

	return h.H[i][0]*h.H[j][1] < h.H[i][1]*h.H[j][0]
}

func (h *hp) Swap(i, j int) {
	h.H[i], h.H[j] = h.H[j], h.H[i]
}
func (h *hp) Len() int {
	return len(h.H)
}

// func (h *heap) push(item int) { //向堆中加入数据
// 	h.H = append(h.H, item)
// 	//往上调整使其符合堆的性质
// 	index := len(h.H)
// 	for index > 1 && h.H[index-1] < h.H[index/2-1] { //当前节点比它的父节点小，需要调整
// 		h.H[index-1], h.H[index/2-1] = h.H[index/2-1], h.H[index-1]
// 		index = index / 2
// 	}
// }
// func (h *heap) pop() (ans int) { //从堆中弹出一个数据
// 	if len(h.H) == 0 {
// 		fmt.Println("堆已空，错误！")
// 		return -1
// 	}
// 	ans = h.H[0]
// 	h.H[0] = h.H[len(h.H)-1]
// 	h.H = h.H[:len(h.H)-1]
// 	//向下沉从而使其符合堆的性质
// 	index := 1
// 	for index <= len(h.H)/2 {

// 		temp_index := index * 2
// 		if temp_index < len(h.H) && h.H[temp_index-1] > h.H[temp_index] { //挑出子节点中较小的那个
// 			temp_index++
// 		}
// 		if h.H[index-1] > h.H[temp_index-1] {
// 			h.H[index-1], h.H[temp_index-1] = h.H[temp_index-1], h.H[index-1] //进行交换
// 			index = temp_index
// 		} else { //已经符合最小堆性质
// 			break
// 		}

// 	}
// 	return
// }
func kthSmallestPrimeFraction(arr []int, k int) (ans []int) {
	h := &hp{}
	h.H = make([][2]int, len(arr)*(len(arr)-1)/2)
	index := 0
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			//heap.Push(h, [2]int{arr[j], arr[i]})
			h.H[index] = [2]int{arr[j], arr[i]}
			index++
		}
	}
	heap.Init(h)
	for i := 1; i < k; i++ {
		heap.Pop(h)
	}
	res := heap.Pop(h)
	return []int{res.([2]int)[0], res.([2]int)[1]}
}

func kthSmallestPrimeFraction2(arr []int, k int) (ans []int) {
	left, right := 0.0, 1.0
	up_item, down_item := 0, 1
	max := []int{arr[0], arr[len(arr)-1]}
	for left < right {
		up_item, down_item = 0, 1
		count := 0
		mid := (left + right) / 2
		max = []int{arr[0], arr[len(arr)-1]}
		for down_item = 1; down_item < len(arr); down_item++ {
			for ; float64(arr[up_item])/float64(arr[down_item]) <= mid; up_item++ {
				if max[0]*arr[down_item] < max[1]*arr[up_item] {
					max = []int{arr[up_item], arr[down_item]}
				}
			}
			count += up_item
		}
		if count == k {
			return max
		} else if count > k {
			right = mid
		} else {
			left = mid
		}
	}

	return max
}
func main() {
	arr := []int{1, 2, 3, 5}
	k := 3
	fmt.Println(kthSmallestPrimeFraction2(arr, k))

}
