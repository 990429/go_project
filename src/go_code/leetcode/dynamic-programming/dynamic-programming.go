package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strconv"
)

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}
func min3(a, b, c int) int {
	result := a
	if result > b {
		result = b
	}
	if result > c {
		result = c
	}
	return result
}
func NthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		num1, num2, num3 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		term := min3(num1, num2, num3)
		dp[i] = term
		if term == num1 {
			p2++
		}
		if term == num2 {
			p3++
		}
		if term == num3 {
			p5++
		}
	}
	return dp[n-1]
}

var factors = []int{2, 3, 5}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func nthUglyNumber(n int) int {
	h := &hp{sort.IntSlice{1}}
	seen := map[int]struct{}{1: {}}
	for i := 1; ; i++ {
		x := heap.Pop(h).(int)
		if i == n {
			return x
		}
		for _, f := range factors {
			next := x * f
			if _, has := seen[next]; !has {
				heap.Push(h, next)
				seen[next] = struct{}{}
			}
		}
	}
}
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minn = min(minn, dp[i-j*j])
		}
		dp[i] = minn + 1
	}
	return dp[n]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func nthSuperUglyNumber(n int, primes []int) int {
	pr := make([]int, len(primes))   //指向每个质因数的下标
	help := make([]int, len(primes)) //存放备选数字
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		min := dp[pr[0]] * primes[0]
		for j := range help {
			help[j] = dp[pr[j]] * primes[j]
			if help[j] < min {
				min = help[j]
			}
		}
		dp[i] = min
		for j := range help {
			if help[j] == min {
				pr[j]++
			}
		}
	}
	return dp[n-1]
}
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin == i {
				dp[i] = 1
			}
			if coin < i {
				if dp[i-coin] > 0 {
					if dp[i] == -1 {
						dp[i] = dp[i-coin] + 1
					} else {
						dp[i] = min(dp[i], dp[i-coin]+1)
					}
				}
			}
		}
	}
	return dp[amount]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func create_tree(tree []int) *TreeNode {
	var root *TreeNode
	queue := []*TreeNode{root}
	index := 0
	for len(queue) > 0 {
		if tree[index] > 0 {
			queue[0] = &TreeNode{}
			queue[0].Val = tree[index]

			index++
			queue = append(queue, queue[0].Left)
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
	}
	return root
}
func rob(root *TreeNode) int {
	f := make(map[*TreeNode]int)
	g := make(map[*TreeNode]int)
	var dfs func(sub_root *TreeNode)
	dfs = func(sub_root *TreeNode) {
		if sub_root == nil {
			return
		}
		dfs(sub_root.Left)
		dfs(sub_root.Right)
		f[sub_root] = sub_root.Val + g[sub_root.Left] + g[sub_root.Right]
		g[sub_root] = max(f[sub_root.Left], g[sub_root.Left]) + max(f[sub_root.Right], g[sub_root.Right])
	}
	dfs(root)
	return max(f[root], g[root])
}
func integerBreak(n int) int {
	dp := make([]int, n+1)
	if n < 4 {
		return n - 1
	}
	for index := 3; index <= n; index++ {
		dp[index] = max(max(2*dp[index-2], 2*(index-2)), max(3*dp[index-3], 3*(index-3)))
	}
	return dp[n]
}
func diffWaysToCompute(expression string) (ans []int) {
	var dfs func(sub_s string) []int
	dfs = func(sub_s string) []int {
		if value, err := strconv.ParseInt(sub_s, 10, 32); err == nil {
			return []int{int(value)}
		}
		result := []int{}
		for i := range sub_s {
			vaule := sub_s[i]
			if vaule == '+' || vaule == '-' || vaule == '*' {
				left := dfs(sub_s[:i])
				right := dfs(sub_s[i+1:])
				for _, l := range left {
					for _, r := range right {
						switch vaule {
						case '+':
							result = append(result, l+r)
						case '-':
							result = append(result, l-r)
						case '*':
							result = append(result, l*r)
						}
					}
				}
			}
		}
		return result
	}

	return dfs(expression)
}
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	result := 0
	for i := range dp {
		dp[i] = 1
		for j := 0; j <= i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}
	return result
}
func maxProfit(prices []int) int {
	dp := make([][3]int, len(prices))
	dp[0][0] = -prices[0]
	dp[0][1] = 0 //不冷冻
	dp[0][2] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][2])
		dp[i][2] = dp[i-1][0] + prices[i]
	}
	return max(dp[len(prices)-1][1], dp[len(prices)-1][2])
}
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	for i := range dp {
		for _, coin := range coins {
			if coin <= i && dp[i-coin] != -1 {
				if dp[i] >= 0 {
					dp[i] = min(dp[i], dp[i-coin]+1)
				} else {
					dp[i] = dp[i-coin] + 1
				}

			}
		}
	}
	return dp[amount]
}
func wiggleMaxLength(nums []int) int {
	dp := make([][2]int, len(nums))
	dp[0][0] = 1 //最后为上升
	dp[0][1] = 1 //最后为下降

	result := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j <= i; j++ {
			if nums[i] > nums[j] {
				dp[i][0] = max(dp[i][0], dp[j][1]+1)
				result = max(result, dp[i][0])
			}
			if nums[i] < nums[j] {
				dp[i][1] = max(dp[i][1], dp[j][0]+1)
				result = max(result, dp[i][1])
			}
		}
	}
	return result
}
func integerReplacement(n int) int {

	var dfs func(sub_n int) int
	dfs = func(sub_n int) int {
		if sub_n == 1 {

			return 0
		}

		if sub_n%2 == 0 {
			return 1 + dfs(sub_n/2)
		}
		return 2 + min(dfs(sub_n/2), dfs(sub_n/2+1))
	}
	return dfs(n)
}
func findSubstringInWraproundString(p string) int {
	result := make(map[string]int)
	dp := make([][]bool, len(p))
	for i := range dp {
		dp[i] = make([]bool, len(p))
		dp[i][i] = true
		result[string(p[i])] = 1
	}
	for i := range dp {
		for j := i + 1; j < len(p); j++ {
			if (p[j-1] == 'z' && p[j] == 'a') || (p[j]-p[j-1] == 1) {
				dp[i][j] = dp[i][j-1]
				if dp[i][j] {
					result[p[i:j+1]] = 1
				}
			}
		}
	}
	return len(result)
}
func makesquare(matchsticks []int) bool {
	total := 0
	for _, value := range matchsticks {
		total += value
	}
	if total%4 != 0 {
		return false
	}

	dp := make([][4]bool, total/4+1)
	flag := make([]bool, len(matchsticks))
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))
	if matchsticks[0] > total/4 {
		return false
	}
	for i := 0; i < 4; i++ {
		dp[0][i] = true
		for length := 1; length <= total/4; length++ {
			for index, value := range matchsticks {
				if flag[index] {
					continue
				}
				if value <= length {
					dp[length][i] = dp[length-value][i]
				}

			}
		}
		if !dp[total/4][i] {
			return false
		}
		begin := total / 4
		for index := len(matchsticks) - 1; index >= 0; index-- {
			if begin >= matchsticks[index] && dp[begin-matchsticks[index]][i] {
				flag[index] = true
				begin -= matchsticks[index]
			}
		}
	}
	return true
}
func findMaxForm(strs []string, m int, n int) int {

	return 0
}
func main() {
	matchsticks := []int{10, 6, 5, 5, 5, 3, 3, 3, 2, 2, 2, 2}
	fmt.Println(makesquare(matchsticks))

}
