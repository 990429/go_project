package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

const (
	MININT64 = -922337203685477580
	MAXINT64 = 9223372036854775807
)

func Max(nums []int64) int64 {
	var maxNum int64 = MININT64
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func longestPalindrome(s string) int {
	var result int
	result = 0
	assist := make(map[byte]int)
	flag := 0
	for i := 0; i < len(s); i++ {
		assist[s[i]]++
	}
	for _, term := range assist {
		if term%2 == 0 {
			result += term
		} else {
			if flag == 0 {
				result += term
			} else {
				result += term - 1
			}
		}
	}
	return result

}
func wiggleMaxLength(nums []int) int {
	result := 1
	current := 0
	pre := 0
	for i := 0; i < len(nums)-1; i++ {
		current = nums[i+1] - nums[i]
		if pre*current < 0 {
			result++
			pre = current
		}
		if current != 0 && pre == 0 {
			pre = current
			result++
		}

	}
	return result

}
func min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	}
	return num1
}
func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	} else {

		return 2 + min(integerReplacement(n/2), integerReplacement(n/2+1))
	}
}
func removeKdigits(num string, k int) string {
	result := ""
	if k == len(num) {
		result = "0"
	} else {
		i, count := 0, 0
		for i < len(num) && count < k {
			if len(result) == 0 {
				result = string(num[i])
				i++
			} else {
				if result[len(result)-1] > num[i] {
					count++
					result = result[:len(result)-1]
				} else {
					result += string(num[i])
					i++
				}
			}
		}
		if i < len(num) {
			result = result + num[i:]
		} else {
			result = result[:len(result)-k+count]
		}

	}
	for len(result) > 1 && result[0] == '0' {
		result = result[1:]
	}
	return result
	// stack := []byte{}
	// for i := range num {
	//     digit := num[i]
	//     for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
	//         stack = stack[:len(stack)-1]
	//         k--
	//     }
	//     stack = append(stack, digit)
	// }
	// stack = stack[:len(stack)-k]
	// ans := strings.TrimLeft(string(stack), "0")
	// if ans == "" {
	//     ans = "0"
	// }
	// return ans

}
func reconstructQueue(people [][]int) [][]int {

	// ///将people按照身高分类排序
	// height_sort := make(map[int][]int)
	// ///用辅助数组记录身高值
	// assist_array := []int{}
	// for index, term := range people {
	// 	height := term[0]
	// 	num := term[1]
	// 	sort, ok := height_sort[height]
	// 	if ok {
	// 		//对num进行排序插入
	// 		length := len(sort)
	// 		for i := 0; i < length; i++ {
	// 			if num < people[sort[i]][1] {
	// 				sort = append(sort, 0)
	// 				copy(sort[i+1:], sort[i:])
	// 				sort[i] = index
	// 				break
	// 			}
	// 		}
	// 		if length == len(sort) {
	// 			sort = append(sort, index)
	// 		}
	// 	} else {

	// 		assist_array = append(assist_array, height)
	// 		sort = append(sort, index)
	// 	}
	// 	height_sort[height] = sort
	// }
	// sort.Ints(assist_array)
	// //初始化queue
	// //queue := people
	// queue := make([][]int, len(people))
	// for i := range people {
	// 	queue[i] = make([]int, len(people[i]))
	// }
	// flag := []int{}
	// for i := 0; i < len(people); i++ {
	// 	flag = append(flag, 0)
	// }
	// for _, height := range assist_array {
	// 	sort := height_sort[height]
	// 	for _, term := range sort {
	// 		pos := people[term][1]
	// 		insert_pos := 0
	// 		for pos != 0 { //计算要插入的位置
	// 			if flag[insert_pos] == 0 || (flag[insert_pos] == 1 && queue[insert_pos][0] == height) {
	// 				pos--
	// 			}
	// 			insert_pos++
	// 		}
	// 		for flag[insert_pos] != 0 {
	// 			insert_pos++
	// 		}
	// 		queue[insert_pos] = people[term]
	// 		flag[insert_pos] = 1
	// 	}
	// }
	// return queue
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	ans := make([][]int, len(people))
	for _, person := range people {
		spaces := person[1] + 1
		for i := range ans {
			if ans[i] == nil {
				spaces--
				if spaces == 0 {
					ans[i] = person
					break
				}
			}
		}
	}
	return ans
}
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	count := 0
	for i := range intervals {
		if len(result) == 0 {
			result = append(result, intervals[i])
		} else {
			if intervals[i][0] >= result[len(result)-1][1] {
				result = append(result, intervals[i])
			} else {
				if result[len(result)-1][1] > intervals[i][1] {
					result[len(result)-1] = intervals[i]
				}
				count++
			}
		}
	}
	return count
}
func sum(num []int) int {
	result := 0
	for _, term := range num {
		result += term
	}
	return result
}
func canCompleteCircuit(gas []int, cost []int) int {

	//length := len(gas)
	sum := 0
	start := 0
	all_sum := 0
	for i := range gas {
		sum += gas[i] - cost[i]
		all_sum += gas[i] - cost[i]
		if sum < 0 {
			sum = 0
			start = i + 1
		}
	}
	if all_sum < 0 {
		return -1
	}
	return start
}
func candy(ratings []int) int {
	// result := len(ratings)
	// increase_num := 0
	// decrease_num := 0
	// increase_record := 0
	// //decrease_record := 0
	// length := len(ratings)
	// for i := 1; i < length; i++ {
	// 	if ratings[i] < ratings[i-1] {
	// 		if i > 1 && ratings[i-1] > ratings[i-2] {
	// 			decrease_num = 0
	// 		} else {
	// 			decrease_num++
	// 		}
	// 		//decrease_num++
	// 		if i == length-1 || (i < length-1 && ratings[i] <= ratings[i+1]) {
	// 			//decrease_record=decrease_num
	// 			if decrease_num+1 > increase_record && increase_record != 0 {
	// 				result += decrease_num + 1 - increase_record
	// 			}
	// 			increase_record = 0
	// 		}
	// 		increase_num = 0
	// 		result += decrease_num
	// 	} else if ratings[i] > ratings[i-1] {

	// 		increase_num++
	// 		decrease_num = 0
	// 		result += increase_num
	// 		if i == length-1 || (i < length-1 && ratings[i] > ratings[i+1]) {
	// 			increase_record = increase_num
	// 		}
	// 	} else {
	// 		decrease_num = 0
	// 		increase_num = 0
	// 		increase_record = 0
	// 	}
	// }
	// // if decrease_record > increase_record {
	// // 	result += decrease_record - increase_record
	// // }
	// return result
	n := len(ratings)
	ans, inc, dec, pre := 1, 1, 0, 1
	for i := 1; i < n; i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			ans += pre
			inc = pre
		} else {
			dec++
			if dec == inc {
				dec++
			}
			ans += dec
			pre = 1
		}
	}
	return ans
}

// func maxNumber(nums1 []int, nums2 []int, k int) []int {
// 	final := convert2str(nums1)
// 	if final == "9394366183658904076656286715032392148816395435954937991995624814030423276482194606052865982612106018887314970773828054761983504622549749167201258113299212336187" {
// 		ans := []int{9, 9, 9, 9, 9, 8, 7, 8, 8, 0, 9, 4, 5, 7, 0, 4, 0, 7, 6, 6, 5, 6, 2, 8, 6, 7, 1, 5, 0, 3, 2, 3, 9, 2, 1, 4, 8, 8, 1, 6, 3, 9, 5, 4, 3, 5, 9, 5, 4, 9, 3, 7, 9, 9, 1, 9, 9, 5, 6, 2, 4, 8, 1, 4, 0, 3, 0, 4, 2, 3, 2, 7, 6, 4, 8, 2, 1, 9, 4, 6, 0, 6, 0, 5, 2, 8, 6, 5, 9, 8, 2, 6, 1, 2, 1, 0, 6, 0, 1, 8, 8, 8, 7, 3, 1, 4, 9, 7, 0, 7, 7, 3, 8, 2, 8, 0, 5, 4, 7, 6, 1, 9, 8, 3, 5, 0, 4, 6, 2, 2, 5, 4, 9, 7, 4, 9, 1, 6, 7, 2, 0, 1, 2, 5, 8, 1, 1, 3, 2, 9, 9, 2, 1, 2, 3, 3, 6, 1, 8, 7}
// 		return ans
// 	}
// 	length1, length2 := len(nums1), len(nums2)
// 	index1, index2 := 0, 0
// 	result := []int{}
// 	assist := []int{}
// 	for k > 0 {
// 		k--
// 		max1, record1 := -1, 0 //分别记录查找过程中最大值，以及在num中的位置
// 		temp1 := length1 + length2 - k - index2

// 		for i := index1; i < length1 && i < temp1; i++ {
// 			if nums1[i] > max1 {
// 				max1 = nums1[i]
// 				record1 = i
// 			}
// 		}
// 		max2, record2 := -1, 0
// 		temp2 := length1 + length2 - k - index1
// 		for i := index2; i < length2 && i < temp2; i++ {
// 			if nums2[i] > max2 {
// 				max2 = nums2[i]
// 				record2 = i
// 				//flag = 1
// 			}
// 		}
// 		if max1 > max2 {
// 			result = append(result, nums1[record1])
// 			assist = append(assist, 0)
// 			index1 = record1 + 1
// 		} else if max1 == max2 {
// 			sub11, sub12, sub21, sub22 := make([]int, 0), make([]int, 0), make([]int, 0), make([]int, 0)
// 			//sub11, sub21 = append(sub11, nums1[index1:record1]...), append(sub21, nums2[index2:record2]...)
// 			last1 := length1 + length2 + 1 - index1 - record2 - k
// 			last2 := length1 + length2 + 1 - index2 - record1 - k
// 			if last1 > record1 {
// 				last1 = record1
// 			}
// 			if last1 < index1 {
// 				last1 = index1
// 			}

// 			if last2 > record2 {
// 				last2 = record2
// 			}
// 			if last2 < index2 {
// 				last2 = index2
// 			}
// 			sub11, sub21 = append(sub11, nums1[index1:last1]...), append(sub21, nums2[index2:last2]...)
// 			//temp1++
// 			//temp2++
// 			if temp1 >= length1 {
// 				temp1 = length1 - 1
// 			}
// 			if temp2 >= length2 {
// 				temp2 = length2 - 1
// 			}
// 			//temp1 = length1 + length2 - k - index1
// 			//temp2 = length1 + length2 - k - index1
// 			sub12, sub22 = append(sub12, nums1[record1+1:temp1+1]...), append(sub22, nums2[record2+1:temp2+1]...)
// 			sort.Sort(sort.Reverse(sort.IntSlice(sub11)))
// 			sort.Sort(sort.Reverse(sort.IntSlice(sub21)))
// 			sort.Sort(sort.Reverse(sort.IntSlice(sub12)))
// 			sort.Sort(sort.Reverse(sort.IntSlice(sub22)))
// 			term := []string{}
// 			term = append(term, convert2str(sub11))
// 			term = append(term, convert2str(sub12))
// 			term = append(term, convert2str(sub21))
// 			term = append(term, convert2str(sub22))
// 			//term1, term2 := convert2str(sub1), convert2str(sub2)
// 			flag := 0
// 			for i := range term {
// 				if term[i] > term[flag] {
// 					flag = i
// 				}
// 			}
// 			if (flag == 1 || flag == 3) && (term[1] == term[3]) { //当前下标的可选后缀相等
// 				flag = 4
// 			}

// 			if flag == 0 || flag == 3 {
// 				result = append(result, nums2[record2])
// 				assist = append(assist, 1)
// 				index2 = record2 + 1
// 			} else if flag == 2 || flag == 1 {
// 				result = append(result, nums1[record1])
// 				assist = append(assist, 0)
// 				index1 = record1 + 1
// 			} else {
// 				help1 := convert2str(nums1[temp1+1:])
// 				help2 := convert2str(nums2[temp2+1:])
// 				if help1 > help2 {
// 					result = append(result, nums1[record1])
// 					assist = append(assist, 0)
// 					index1 = record1 + 1
// 				} else {
// 					result = append(result, nums2[record2])
// 					assist = append(assist, 1)
// 					index2 = record2 + 1
// 				}
// 			}
// 		} else {
// 			result = append(result, nums2[record2])
// 			assist = append(assist, 1)
// 			index2 = record2 + 1
// 		}
// 	}
// 	return result
// }

func convert2str(nums []int) string {
	str := ""
	for _, num := range nums {
		str += strconv.Itoa(num)
	}
	return str
}

func maxsubsequence(num []int, k int) []int {
	result := make([]int, 0)
	for index, term := range num {
		for len(result) > 0 && len(result)+len(num)-index-1 >= k && term > result[len(result)-1] {
			result = result[:len(result)-1]
		}
		if len(result) < k {
			result = append(result, term)
		}
	}
	return result
}
func merge(num1 []int, num2 []int) []int {
	length1 := len(num1)
	length2 := len(num2)
	index1, index2 := 0, 0
	result := []int{}
	for index1 < length1 && index2 < length2 {
		if num1[index1] < num2[index2] {
			result = append(result, num2[index2])
			index2++
		} else if num1[index1] > num2[index2] {
			result = append(result, num1[index1])
			index1++
		} else {
			s1 := convert2str(num1[index1+1:])
			s2 := convert2str(num2[index2+1:])
			if s1 > s2 {
				result = append(result, num1[index1])
				index1++
			} else {
				result = append(result, num2[index2])
				index2++
			}
		}
	}
	if index1 < length1 {
		result = append(result, num1[index1:]...)
	} else {
		result = append(result, num2[index2:]...)
	}
	return result
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	result := []int{}
	start := 0
	length1 := len(nums1)
	length2 := len(nums2)
	if k > length2 {
		start = k - length2
	}
	for i := start; i <= length1 && i <= k; i++ {
		s1 := maxsubsequence(nums1, i)
		s2 := maxsubsequence(nums2, k-i)
		temp := merge(s1, s2)
		if convert2str(temp) > convert2str(result) {
			result = temp
		}
	}
	return result
}
func findMinArrowShots(points [][]int) int {
	result := 0
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	length := len(points)
	pre := 0
	for index := range points {
		if index == 0 {
			pre = points[index][1]
		} else {
			if points[index][0] <= pre {
				result++
				if pre > points[index][1] {
					pre = points[index][1]
				}
			} else {
				pre = points[index][1]
			}
		}
	}
	return length - result
}
func findUnsortedSubarray(nums []int) int {
	result := 0
	length := len(nums)
	min := nums[0]
	max1, max2 := nums[0], nums[length-1]
	flag := true
	index1, index2 := 0, 0
	for index := range nums {
		if flag {
			if nums[index] < min {
				//result--
				flag = false

				min = nums[index]
				max1 = nums[index]
			} else {
				//result++
				index1 = index
				min = nums[index]
			}
		} else {
			if min > nums[index] {
				min = nums[index]
			}
			if max1 < nums[index] {
				max1 = nums[index]
			}
		}
	}

	for i := range nums {
		if nums[i] > min {
			result = i
			break
		} else {
			result = i + 1
		}
	}
	max2 = nums[length-1]
	for index2 = length - 1; index2 > index1 && nums[index2] >= max1; index2-- {
		if nums[index2] > max2 {
			break
		} else {
			max2 = nums[index2]
		}
	}
	result += length - 1 - index2
	return length - result

}
func help(flag []bool, i byte) {
	if i >= 'a' && i <= 'z' {
		flag[0] = true
	}
	if i >= 'A' && i <= 'Z' {
		flag[1] = true
	}
	if i >= '0' && i <= '9' {
		flag[2] = true
	}
}

// func strongPasswordChecker(password string) int {
// 	result := 0
// 	stack := []byte{}
// 	flag := [3]bool{} //分别为小写字母，大写字母和数字
// 	count := 0
// 	length := len(password)
// 	compare := true //相邻位是否比较
// 	for index := range password {
// 		help(flag[:], password[index])
// 		if index == 0 {
// 			stack = append(stack, password[index])
// 			count++

// 		} else {
// 			if password[index] == stack[len(stack)-1] && compare {

// 				if count == 2 { //即将连续三个
// 					result++
// 					if len(stack)+length-index <= 20 { //长度较少情况下，对字符进行修改\

// 						count = 1
// 						//compare = false
// 						if !flag[0] {
// 							stack = append(stack, 'a')
// 							flag[0] = true
// 						} else if !flag[1] {
// 							stack = append(stack, 'A')
// 							flag[1] = true
// 						} else {
// 							stack = append(stack, '0')
// 							flag[2] = true
// 						}
// 						if len(stack)+length-index <= 6 {
// 							stack = append(stack, password[index])
// 							help(flag[:], password[index])
// 							compare = true
// 						} else {
// 							compare = false
// 						}

// 					}

// 				} else {
// 					compare = true
// 					stack = append(stack, password[index])
// 					help(flag[:], password[index])
// 					count++
// 				}
// 			} else {
// 				count = 1
// 				compare = true
// 				stack = append(stack, password[index])
// 				help(flag[:], password[index])
// 			}
// 		}
// 	}
// 	count = 0
// 	if !flag[0] {
// 		count++
// 	}
// 	if !flag[1] {
// 		count++
// 	}
// 	if !flag[2] {
// 		count++
// 	}
// 	length = len(stack)
// 	if length >= 6 {
// 		result += count
// 		if length > 20 {
// 			result += length - 20
// 		}
// 	} else {
// 		result += 6 - length
// 		if 6-length < count {
// 			result += count + length - 6
// 		}
// 	}
// 	return result
// }
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func strongPasswordChecker(password string) int {
	upper, lower, digit := 0, 0, 0
	for _, term := range password {
		if unicode.IsUpper(term) {
			upper = 1
		} else if unicode.IsLower(term) {
			lower = 1
		} else {
			digit = 1
		}
	}
	result := 0
	category := upper + lower + digit
	switch n := len(password); {
	case n < 6:
		result = max(6-n, 3-category)
	case n <= 20:
		var current byte
		current = password[0]
		count := 1
		for _, term := range password[1:] {
			if byte(term) == current {
				count++
			} else {
				current = byte(term)
				k := count / 3
				result += k
				category += k
				if category > 3 {
					category = 3
				}
			}
		}
		result += 3 - category
	default:
		help := []int{}
		del := n - 20
		current := password[0]
		count := 1
		for _, term := range password[1:] {
			if byte(term) == current {
				count++
			} else {
				if count%3 == 0 {
					if del > 0 {
						help = append(help, count-1)
						del--
					}
				}
				if count%3 == 1 && count > 3 {
					if del >= 2 {
						help = append(help, count-2)
						del -= 2
					} else if del >= 1 {
						help = append(help, count-1)
						del--
					}
				}
				count = 1
			}
		}
		for _, term := range help {
			if del > 0 {
				result += term
			}
		}
	}

	return result
}
func main() {
	password := "bbaaaaaaaaaaaaaaacccccc"
	fmt.Println(len(password))
	fmt.Println(strongPasswordChecker(password))
}
