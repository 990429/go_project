package main

import "sort"

/*********************************************/
/*********************************************/
//array
/*********************************************/
/*********************************************/

// func threeSum(nums []int) (res [][]int) {//递归回溯方法超时
// 	ans := []int{}
// 	sort.Ints(nums)
// 	sum := 0
// 	visited := make(map[int]bool)

// 	var dfs func(sub_nums []int, sum int, ans []int, left int)
// 	dfs = func(sub_nums []int, sum int, ans []int, left int) {
// 		if sum == 0 && len(ans) == 3 {
// 			res = append(res, ans)

// 			return
// 		}
// 		for index := left; index < len(sub_nums); index++ {
// 			if sum+sub_nums[index] <= 0 && len(ans) <= 3 { //怎么去重？（判断：如果当前值与上一个值相等，但上一个值没有没使用，说明有重复）
// 				if index > 0 && !visited[index-1] && sub_nums[index-1] == sub_nums[index] {
// 					continue
// 				} else {
// 					visited[index] = true

// 					dfs(sub_nums, sum+sub_nums[index], append(ans, sub_nums[index]), index+1)

// 					visited[index] = false
// 				}
// 			} else {
// 				return
// 			}

// 		}
// 	}
// 	dfs(nums, sum, ans, 0)
// 	//fmt.Println(track)
// 	return
// }
func spiralOrder(matrix [][]int) (ans []int) {
	row, col := len(matrix), len(matrix[0])
	pre, cur := [2]int{0, 0}, [2]int{0, col - 1}
	for i := 0; i < col; i++ {
		ans = append(ans, matrix[0][i])
	}
	for (cur[0] != row/2 || pre[0] != row/2) && (cur[1] != col/2 || pre[1] != col/2) {
		if cur[0] == pre[0] {
			if cur[1] >= pre[1] { //向下
				for i := cur[0] + 1; i < row-cur[0]; i++ {
					ans = append(ans, matrix[i][cur[1]])
				}
				pre = cur
				cur = [2]int{row - cur[0] - 1, cur[1]}
			} else { //向上
				for i := cur[0] - 1; i > row-cur[0]-1; i-- {
					ans = append(ans, matrix[i][cur[1]])
				}
				pre = cur
				cur = [2]int{row - cur[0], cur[1]}
			}
		} else if cur[1] == pre[1] {
			if cur[0] < pre[0] { //向右
				for i := cur[1] + 1; i < col-cur[1]-1; i++ {
					ans = append(ans, matrix[cur[0]][i])
				}
				pre = cur
				cur = [2]int{cur[0], col - cur[1] - 2}
			} else { //向左
				for i := cur[1] - 1; i >= col-cur[1]-1; i-- {
					ans = append(ans, matrix[cur[0]][i])
				}
				pre = cur
				cur = [2]int{cur[0], col - cur[1] - 1}
			}
		}

	}
	return
}
func search(nums []int, target int) int {

	left, right := 0, len(nums)-1
	mid := (left + right) / 2

	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] { //在左边的升序里面
			if nums[left] <= target && nums[mid] >= target {
				right = mid
			} else {
				left = mid + 1
			}
		} else { //在右边的升序里面
			if nums[mid] <= target && nums[right] >= target {
				left = mid
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
func threeSum(nums []int) (res [][]int) { //使用排序和双指针方法
	sort.Ints(nums)
	if nums[0] > 0 {
		return
	}
	//ans := []int{}

	//visited := make([]bool, len(nums))
	left, right := 0, len(nums)-1
	for index1 := 0; index1 < len(nums)-2; index1++ {
		//ans = append(ans, nums[index1])
		if index1 > 0 && nums[index1-1] == nums[index1] {
			continue
		}
		if nums[index1] > 0 {
			break
		}
		left, right = index1+1, len(nums)-1
		for left < right {
			if left-1 != index1 && nums[left-1] == nums[left] {
				left++
				continue
			}
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			sum := nums[index1] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum == 0 {
				res = append(res, []int{nums[index1], nums[left], nums[right]})
				left++
			} else {
				left++
			}
		}
	}
	return
}

func trap(height []int) int { //超时（思路：从高度为1开始逐渐加入雨水）

	res := 0
	max_height := 0
	for i := range height { //找出最高的柱子
		if height[i] > max_height {
			max_height = height[i]
		}
	}
	for h := 1; h <= max_height; h++ {
		first, second := -1, -1
		for i := range height {
			if height[i] >= h {
				if first == -1 {
					first = i
				} else {
					second = i
					for j := first + 1; j < second; j++ {
						res++
					}
					first = second
				}
			}
		}
	}
	return res
}

func trap2(height []int) int { //(找出先降后升的地方)
	sum := 0
	left_index, right_index := 0, len(height)-1
	left_max, right_max := 0, 0
	for left_index < right_index {
		if height[left_index] < height[right_index] {
			left_max = max(left_max, height[left_index])
			sum += left_max - height[left_index]
			left_index++
		} else {
			right_max = max(right_max, height[right_index])
			sum += right_max - height[right_index]
			right_index--
		}
	}
	return sum
}

	func merge(intervals [][]int) (ans [][]int) {
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		})
		if len(intervals) == 0 {
			return
		}
		start, end := intervals[0][0], intervals[0][1]
		for index := 1; index < len(intervals); index++ {
			if end >= intervals[index][0] {
				end = max(end, intervals[index][1])
			} else {
				ans = append(ans, []int{start, end})
				start, end = intervals[index][0], intervals[index][1]
			}
		}
		ans = append(ans, []int{start, end})
		return
	}
