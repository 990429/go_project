package main

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	res := 1
	for i := range dp {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1

			}
		}
		res = max(res, dp[i])
	}
	return res
}
func lengthOfLIS2(nums []int) int {
	dp := make([]int, len(nums)+1)
	length := 1
	dp[length] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > dp[length] {
			length++
			dp[length] = nums[i]
		} else { //在dp中二分查找
			l, r, pos := 1, length, 0
			mid := 0
			for l <= r {
				mid = (l + r) / 2
				if dp[mid] < nums[i] {
					pos = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			dp[pos+1] = nums[i]

		}
	}
	return length
}
func lcs(word1, word2 string) string { //求出两个字符串的最长公共子序列
	dp := make([][]int, len(word1))
	for index := range dp {
		dp[index] = make([]int, len(word2))
	}
	for index := 0; index < len(word1); index++ {
		if word1[index] == word2[0] {
			dp[index][0] = 1
		}
	}
	for index := 0; index < len(word2); index++ {
		if word2[index] == word1[0] {
			dp[0][index] = 1
		}
	}
	for i := 1; i < len(word1); i++ {
		for j := 1; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	res := make([]byte, dp[len(word1)-1][len(word2)-1])
	for i, j, index := len(word1)-1, len(word2)-1, len(res)-1; index >= 0; {
		if word1[i] == word2[j] {
			res[index] = word1[i]
			i--
			j--
			index--
		} else {
			if dp[i-1][j] > dp[i][j-1] {
				i--
			} else {
				j--
			}
		}
	}
	return string(res)
}
func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	dp := make([][]int, len1+1)
	if len1*len2 == 0 {
		return len1 + len2
	}
	for index := range dp {
		dp[index] = make([]int, len2+1)
	}
	for index := 0; index <= len(word1); index++ {
		dp[index][0] = index
	}
	for index := 0; index <= len(word2); index++ {
		dp[0][index] = index
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i-1][j-1]-1, min(dp[i-1][j], dp[i][j-1])) + 1
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}

	return dp[len1][len2]
}
