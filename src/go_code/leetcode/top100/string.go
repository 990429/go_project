package main

//string
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans := 1
	included := make(map[byte][2]int)
	included[s[0]] = [2]int{1, 0} //位置1表示是否存在，2表示存在位置
	start := 0
	for index := 1; index < len(s); index++ {

		if included[s[index]][0] != 0 {
			ans = max(ans, index-start)
			for i := start; i < included[s[index]][1]; i++ {
				included[s[i]] = [2]int{0, 0}
			}
			start = included[s[index]][1] + 1
		} else {
			ans = max(ans, index-start+1)
		}
		included[s[index]] = [2]int{1, index}
	}
	return ans
}

func isValid(s string) bool {
	stack := []byte{}
	for i := range s {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			} else if s[i] == '}' {
				if stack[len(stack)-1] == '{' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else if s[i] == ']' {
				if stack[len(stack)-1] == '[' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else if s[i] == ')' {
				if stack[len(stack)-1] == '(' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func longestPalindrome(s string) string {
	left, right := 0, 0
	for index := range s {
		l, r := index, index
		for l >= 0 && r <= len(s)-1 {
			if s[l] == s[r] {
				l--
				r++
			} else {
				break
			}
		}
		if r-l-1 > right-left+1 {
			left = l + 1
			right = r - 1
		}
		if index < len(s)-1 && s[index] == s[index+1] {
			l, r = index, index+1
			for l >= 0 && r <= len(s)-1 {
				if s[l] == s[r] {
					l--
					r++
				} else {
					break
				}
			}
			if r-l-1 > right-left+1 {
				left = l + 1
				right = r - 1
			}
		}
	}
	return s[left : right+1]
}
