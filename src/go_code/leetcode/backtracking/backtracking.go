package main

import (
	"fmt"
	"sort"
)

var results []string

func track_letterCombinations(result string, mapping map[byte][]byte, digits string, depth int) {
	if depth == len(digits) {
		results = append(results, result)
	} else {
		for _, term := range mapping[digits[depth]] {
			//result =
			track_letterCombinations(result+string(term), mapping, digits, depth+1)
		}
	}
}
func letterCombinations(digits string) []string {
	//results := []string{}
	results = []string{}
	if len(digits) == 0 {
		return results
	}
	result := ""
	mapping := make(map[byte][]byte)
	var index, value byte
	for index, value = '2', 'a'; index <= '9'; index++ {
		mapping[index] = append(mapping[index], value)
		mapping[index] = append(mapping[index], value+1)
		mapping[index] = append(mapping[index], value+2)
		value += 3
		if index == '7' || index == '9' {
			mapping[index] = append(mapping[index], value)
			value++
		}
	}
	track_letterCombinations(result, mapping, digits, 0)
	return results
}
func is_match(str string, term byte, n int) bool { //判断加入term后字符串是否违法，str本身是合法的
	count1, count2 := 0, 0 //左右括号数量
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			count1++
		} else {
			count2++
		}
	}
	if term == '(' {
		count1++
	} else {
		count2++
	}
	if count1 >= count2 && count1 <= n {
		return true
	} else {
		return false
	}
}
func track_generateParenthesis(result string, num1 int, num2 int, n int) {
	if num1 == n && num2 == n {
		results = append(results, result)
	} else {
		if is_match(result, '(', n) {
			track_generateParenthesis(result+"(", num1+1, num2, n)
		}
		if is_match(result, ')', n) {
			track_generateParenthesis(result+")", num1, num2+1, n)
		}
	}
}
func generateParenthesis(n int) []string {
	results = []string{}
	track_generateParenthesis("", 0, 0, n)
	return results
}

var combinaSum [][]int

func track_combinationSum(sum int, target int, combina []int, candidates []int) {
	if sum == target {
		help := make([]int, len(combina))
		copy(help, combina)
		combinaSum = append(combinaSum, help)
	} else {
		for _, term := range candidates {
			if sum+term <= target && (len(combina) == 0 || combina[len(combina)-1] <= term) {
				track_combinationSum(sum+term, target, append(combina, term), candidates)
			}
		}
	}
}
func combinationSum(candidates []int, target int) [][]int {
	combinaSum = [][]int{}
	combina := []int{}
	sort.Ints(candidates)
	track_combinationSum(0, target, combina, candidates)
	return combinaSum
}
func isValidSudoku(board [][]byte) bool {

	return true
}
func help_shudu(board [][]byte, i int, j int) []byte {
	avail := []byte{}
	help := [9]bool{}
	for index := 0; index < 9; index++ {
		term := int(board[i][index]) - int('1')
		if term < 9 && term >= 0 {
			help[term] = true
		}
	}
	for index := 0; index < 9; index++ {
		term := int(board[index][j]) - int('1')
		if term < 9 && term >= 0 {
			help[term] = true
		}
	}
	for index1 := i / 3 * 3; index1 < i/3*3+3; index1++ {
		for index2 := j / 3 * 3; index2 < j/3*3+3; index2++ {

			term := int(board[index1][index2]) - int('1')
			if term < 9 && term >= 0 {
				help[term] = true
			}
		}
	}
	for term := range help {
		if !help[term] {
			avail = append(avail, byte(term+'1'))
		}
	}
	return avail
}
func recurse_solveSudoku(board [][]byte, i int, j int) {
	if board[i][j] != '.' {
		if i < 9 && j < 8 {
			recurse_solveSudoku(board, i, j+1)
		}
		if i < 8 && j == 8 {
			recurse_solveSudoku(board, i+1, 0)
		}
		// if i == 8 && j == 8 {
		// 	return
		// }
	} else {
		avail := help_shudu(board, i, j)
		for _, term := range avail {
			board[i][j] = term
			if i < 9 && j < 8 {
				recurse_solveSudoku(board, i, j+1)
			}
			if i < 8 && j == 8 {
				recurse_solveSudoku(board, i+1, 0)
			}
			// if i == 8 && j == 8 {
			// 	return
			// }
			//board[i][j] = '.'
		}
	}

}
func solveSudoku(board [][]byte) {
	recurse_solveSudoku(board, 0, 0)
}

// var result_combinationSum2 [][]int

// func Track(cands []int, index_cands int, result []int, target int) {
// 	if target == 0 {
// 		result_combinationSum2 = append(result_combinationSum2, result)
// 	} else {
// 		for i := index_cands; i < len(cands); i++ {
// 			value := cands[i]
// 			if value <= target {
// 				//Track(cands, i+1, append(result, value), target-value)
// 				if i > 0 && cands[i] == cands[i-1] && ((len(result) > 0 && result[len(result)-1] != cands[i]) || len(result) == 0) {
// 					//continue
// 				} else {
// 					Track(cands, i+1, append(result, value), target-value)
// 				}
// 			}
// 		}
// 	}
// }
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func combinationSum2(candidates []int, target int) (ans [][]int) {

	sort.Ints(candidates)
	//help := [][2]int{}
	var help [][2]int
	for _, term := range candidates {
		if help == nil || term != help[len(help)-1][0] {
			help = append(help, [2]int{term, 1})
		} else {
			help[len(help)-1][1]++
		}
	}
	result := []int{}
	var dfs func(pos, res int)
	dfs = func(pos, res int) {
		if res == 0 {
			term := make([]int, len(result))
			copy(term, result)
			ans = append(ans, term)
			return
		}
		if pos == len(help) || res < help[pos][0] {
			return
		}
		dfs(pos+1, res)
		times := min(help[pos][1], res/help[pos][0])
		for i := 1; i <= times; i++ {
			result = append(result, help[pos][0])
			dfs(pos+1, res-i*help[pos][0])
		}
		result = result[:len(result)-times]
	}
	dfs(0, target)
	return
}
func permute(nums []int) (ans [][]int) {
	result := []int{}

	var dfs func(sub_nums []int, result []int)
	dfs = func(sub_nums []int, result []int) {
		if len(sub_nums) == 0 {
			ans = append(ans, append([]int(nil), result...))
			return
		}
		for index, value := range sub_nums {
			//result = append(result, value)
			term1 := append([]int(nil), sub_nums[:index]...) ///append会改变原来slice内容，所以应该新开辟一个空间
			term1 = append(term1, sub_nums[index+1:]...)
			dfs(term1, append(result, value))

			//result = result[:len(result)-1]
		}
	}
	dfs(nums, result)
	return
}
func permuteUnique(nums []int) (ans [][]int) {
	var freq [][2]int
	sort.Ints(nums)
	///统计出来各个数字的频率
	for _, value := range nums {
		if freq == nil || value != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{value, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}
	var help []int
	var dfs func(sub_freq [][2]int, rest_num int, help []int)
	dfs = func(sub_freq [][2]int, rest_num int, help []int) {
		if rest_num == 0 {
			ans = append(ans, append([]int(nil), help...))
			return
		}
		for index, value := range sub_freq {
			if value[1] > 0 {
				term := append([][2]int(nil), sub_freq...)
				term[index][1]--
				dfs(term, rest_num-1, append(help, value[0]))
			}
			// for i := 0; i < value[1]; i++ {
			// 	term := append([][2]int(nil), sub_freq...)
			// 	term[index][1]--
			// 	dfs(term, rest_num-1, append(help, value[0]))
			// }
		}
	}
	dfs(freq, len(nums), help)
	return
}
func solveNQueens(n int) (ans [][]string) {
	flag := make([]bool, n)

	//pre := -1
	result := []string{}
	diag1 := make([]bool, 2*n-1) //正对角线
	diag2 := make([]bool, 2*n-1) //负对角线
	var dfs func(sub_n int)
	dfs = func(sub_n int) {
		if sub_n == n {
			ans = append(ans, append([]string(nil), result...))
			return
		}
		for i := 0; i < n; i++ {
			if !flag[i] && !diag1[i-sub_n+n-1] && !diag2[i+sub_n] { //符合要求的位置
				term := make([]byte, n)
				for index := range term {
					if index == i {
						term[index] = 'Q'
					} else {
						term[index] = '.'
					}
				}
				result = append(result, string(term))
				diag1[i-sub_n+n-1] = true
				diag2[i+sub_n] = true
				flag[i] = true

				dfs(sub_n + 1)
				result = result[:len(result)-1]
				flag[i] = false
				diag1[i-sub_n+n-1] = false
				diag2[i+sub_n] = false
			}
		}
	}
	dfs(0)
	return
}
func totalNQueens(n int) (ans int) {
	flag := make([]bool, n)

	//pre := -1
	result := []string{}
	diag1 := make([]bool, 2*n-1) //正对角线
	diag2 := make([]bool, 2*n-1) //负对角线
	var dfs func(sub_n int)
	dfs = func(sub_n int) {
		if sub_n == n {
			ans += 1
			return
		}
		for i := 0; i < n; i++ {
			if !flag[i] && !diag1[i-sub_n+n-1] && !diag2[i+sub_n] { //符合要求的位置
				term := make([]byte, n)
				for index := range term {
					if index == i {
						term[index] = 'Q'
					} else {
						term[index] = '.'
					}
				}
				result = append(result, string(term))
				diag1[i-sub_n+n-1] = true
				diag2[i+sub_n] = true
				flag[i] = true

				dfs(sub_n + 1)
				result = result[:len(result)-1]
				flag[i] = false
				diag1[i-sub_n+n-1] = false
				diag2[i+sub_n] = false
			}
		}
	}
	dfs(0)
	return
}
func Is_good(str1, str2 string) (flag bool) {
	flag = false
	count := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			count++
		}
	}
	if count == 1 {
		flag = true
	}
	return
}
func findLadders(beginWord string, endWord string, wordList []string) (ans [][]string) {
	flag := false
	begin := []int{}
	for index, str := range wordList { //判断wordlist中是否有end
		if str == endWord {
			flag = true
		}
		if str == beginWord {
			begin = append(begin, index)
		}
	}
	for _, index := range begin {
		wordList = append(wordList[:index], wordList[index+1:]...)
	}
	if !flag {
		return
	}
	//used := make([]bool, len(wordList))
	graph := make(map[string][]string)
	//count := 0
	queue := []string{beginWord}
	short_distance := make(map[string]int)
	for i := 0; i < len(wordList); i++ { //创建图
		if Is_good(beginWord, wordList[i]) {
			graph[beginWord] = append(graph[beginWord], wordList[i])
		}
	}
	for _, start := range wordList {
		for _, end := range wordList {
			if Is_good(start, end) {
				graph[start] = append(graph[start], end)
			}
		}
	}
	used := make(map[string]bool) //记录节点是否使用过
exit:
	for i := 0; i < len(queue); i++ {
		j_length := len(graph[queue[i]])
		for j := 0; j < j_length; j++ {
			value := graph[queue[i]][j]
			if !used[value] && Is_good(queue[i], value) {
				used[value] = true
				queue = append(queue, value)

				if short_distance[value] == 0 || short_distance[value] > short_distance[queue[i]]+1 {
					short_distance[value] = short_distance[queue[i]] + 1
				}
				if value == endWord {
					break exit
				}
			}
		}
	}
	// if len(queue)-1 != len(wordList) {
	// 	return
	// }
	target := short_distance[endWord]
	result := []string{beginWord}
	used = make(map[string]bool)
	var dfs func(key string, distance int, used map[string]bool)
	dfs = func(key string, distance int, used map[string]bool) {
		if distance == target {
			if key == endWord {
				ans = append(ans, append([]string(nil), result...))
			}
			return
		}
		for _, value := range graph[key] {
			if !used[value] {
				used[value] = true
				result = append(result, value)
				dfs(value, distance+1, used)
				used[value] = false
				result = result[:len(result)-1]
			}
		}
	}
	dfs(beginWord, 0, used)
	return
}
func wordBreak(s string, wordDict []string) (ans []string) {
	dict := make(map[string]bool)
	for _, value := range wordDict {
		dict[value] = true
	}
	result := ""
	var dfs func(sub_s string)
	dfs = func(sub_s string) {
		if sub_s == "" {
			ans = append(ans, result[:len(result)-1])
			return
		}
		for i := range sub_s {
			if dict[sub_s[:i+1]] {
				result += sub_s[:i+1]
				result += " "
				dfs(sub_s[i+1:])
				result = result[:len(result)-i-2]
			}
		}
	}
	dfs(s)
	return
}
func FindWords(board [][]byte, words []string) (ans []string) {
	dict := make(map[byte][][2]int)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dict[board[i][j]] = append(dict[board[i][j]], [2]int{i, j})
		}
	}
	used := make([][]bool, len(board)) //监督当前字符是否用过
	added := make(map[string]bool)     //监督string是否已经加进答案，防止重复
	for i := 0; i < len(used); i++ {
		used[i] = make([]bool, len(board[0]))
	}
	for _, word := range words {
		added[word] = false
	}
	max_row, max_col := len(board), len(board[0])
	var dfs func(row int, col int, target string)
	result := ""
	dfs = func(row, col int, target string) {
		if target == "" {
			if !added[result] {
				added[result] = true
				ans = append(ans, result)
			}
			return
		}
		left, right, up, down := col-1, col+1, row-1, row+1
		if left >= 0 && !used[row][left] && board[row][left] == target[0] { //左
			used[row][left] = true
			result += string(target[0])
			dfs(row, left, target[1:])
			result = result[:len(result)-1]
			used[row][left] = false
		}
		if right < max_col && !used[row][right] && board[row][right] == target[0] { //右
			used[row][right] = true
			result += string(target[0])
			dfs(row, right, target[1:])
			result = result[:len(result)-1]
			used[row][right] = false
		}
		if up >= 0 && !used[up][col] && board[up][col] == target[0] { //上
			used[up][col] = true
			result += string(target[0])
			dfs(up, col, target[1:])
			result = result[:len(result)-1]
			used[up][col] = false
		}
		if down < max_row && !used[down][col] && board[down][col] == target[0] { //下
			used[down][col] = true
			result += string(target[0])
			dfs(down, col, target[1:])
			result = result[:len(result)-1]
			used[down][col] = false
		}
	}
	for _, word := range words {
		for _, pos := range dict[word[0]] {
			result += string(word[0])
			used[pos[0]][pos[1]] = true
			dfs(pos[0], pos[1], word[1:])
			result = ""
			used[pos[0]][pos[1]] = false
		}
	}
	return
}

type Trie struct {
	children [26]*Trie
	word     string
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.word = word
}

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findWords(board [][]byte, words []string) []string {
	t := &Trie{}
	for _, word := range words {
		t.Insert(word)
	}

	m, n := len(board), len(board[0])
	seen := map[string]bool{}

	var dfs func(node *Trie, x, y int)
	dfs = func(node *Trie, x, y int) {
		ch := board[x][y]
		node = node.children[ch-'a']
		if node == nil {
			return
		}

		if node.word != "" {
			seen[node.word] = true
		}

		board[x][y] = '#'
		for _, d := range dirs {
			nx, ny := x+d.x, y+d.y
			if 0 <= nx && nx < m && 0 <= ny && ny < n && board[nx][ny] != '#' {
				dfs(node, nx, ny)
			}
		}
		board[x][y] = ch
	}
	for i, row := range board {
		for j := range row {
			dfs(t, i, j)
		}
	}

	ans := make([]string, 0, len(seen))
	for s := range seen {
		ans = append(ans, s)
	}
	return ans
}

type tree struct {
}

func addOperators(num string, target int) (ans []string) {

	return
}
func main() {

	// board := [][]byte{
	// 	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	// }
	// solveSudoku(board)
	// fmt.Printf("%q", board)

	// beginWord := "cet"
	// endWord := "ism"
	// wordList := []string{"kid", "tag", "pup", "ail", "tun", "woo", "erg", "luz", "brr", "gay", "sip", "kay", "per", "val", "mes", "ohs", "now", "boa", "cet", "pal", "bar", "die", "war", "hay", "eco", "pub", "lob", "rue", "fry", "lit", "rex", "jan", "cot", "bid", "ali", "pay", "col", "gum", "ger", "row", "won", "dan", "rum", "fad", "tut", "sag", "yip", "sui", "ark", "has", "zip", "fez", "own", "ump", "dis", "ads", "max", "jaw", "out", "btu", "ana", "gap", "cry", "led", "abe", "box", "ore", "pig", "fie", "toy", "fat", "cal", "lie", "noh", "sew", "ono", "tam", "flu", "mgm", "ply", "awe", "pry", "tit", "tie", "yet", "too", "tax", "jim", "san", "pan", "map", "ski", "ova", "wed", "non", "wac", "nut", "why", "bye", "lye", "oct", "old", "fin", "feb", "chi", "sap", "owl", "log", "tod", "dot", "bow", "fob", "for", "joe", "ivy", "fan", "age", "fax", "hip", "jib", "mel", "hus", "sob", "ifs", "tab", "ara", "dab", "jag", "jar", "arm", "lot", "tom", "sax", "tex", "yum", "pei", "wen", "wry", "ire", "irk", "far", "mew", "wit", "doe", "gas", "rte", "ian", "pot", "ask", "wag", "hag", "amy", "nag", "ron", "soy", "gin", "don", "tug", "fay", "vic", "boo", "nam", "ave", "buy", "sop", "but", "orb", "fen", "paw", "his", "sub", "bob", "yea", "oft", "inn", "rod", "yam", "pew", "web", "hod", "hun", "gyp", "wei", "wis", "rob", "gad", "pie", "mon", "dog", "bib", "rub", "ere", "dig", "era", "cat", "fox", "bee", "mod", "day", "apr", "vie", "nev", "jam", "pam", "new", "aye", "ani", "and", "ibm", "yap", "can", "pyx", "tar", "kin", "fog", "hum", "pip", "cup", "dye", "lyx", "jog", "nun", "par", "wan", "fey", "bus", "oak", "bad", "ats", "set", "qom", "vat", "eat", "pus", "rev", "axe", "ion", "six", "ila", "lao", "mom", "mas", "pro", "few", "opt", "poe", "art", "ash", "oar", "cap", "lop", "may", "shy", "rid", "bat", "sum", "rim", "fee", "bmw", "sky", "maj", "hue", "thy", "ava", "rap", "den", "fla", "auk", "cox", "ibo", "hey", "saw", "vim", "sec", "ltd", "you", "its", "tat", "dew", "eva", "tog", "ram", "let", "see", "zit", "maw", "nix", "ate", "gig", "rep", "owe", "ind", "hog", "eve", "sam", "zoo", "any", "dow", "cod", "bed", "vet", "ham", "sis", "hex", "via", "fir", "nod", "mao", "aug", "mum", "hoe", "bah", "hal", "keg", "hew", "zed", "tow", "gog", "ass", "dem", "who", "bet", "gos", "son", "ear", "spy", "kit", "boy", "due", "sen", "oaf", "mix", "hep", "fur", "ada", "bin", "nil", "mia", "ewe", "hit", "fix", "sad", "rib", "eye", "hop", "haw", "wax", "mid", "tad", "ken", "wad", "rye", "pap", "bog", "gut", "ito", "woe", "our", "ado", "sin", "mad", "ray", "hon", "roy", "dip", "hen", "iva", "lug", "asp", "hui", "yak", "bay", "poi", "yep", "bun", "try", "lad", "elm", "nat", "wyo", "gym", "dug", "toe", "dee", "wig", "sly", "rip", "geo", "cog", "pas", "zen", "odd", "nan", "lay", "pod", "fit", "hem", "joy", "bum", "rio", "yon", "dec", "leg", "put", "sue", "dim", "pet", "yaw", "nub", "bit", "bur", "sid", "sun", "oil", "red", "doc", "moe", "caw", "eel", "dix", "cub", "end", "gem", "off", "yew", "hug", "pop", "tub", "sgt", "lid", "pun", "ton", "sol", "din", "yup", "jab", "pea", "bug", "gag", "mil", "jig", "hub", "low", "did", "tin", "get", "gte", "sox", "lei", "mig", "fig", "lon", "use", "ban", "flo", "nov", "jut", "bag", "mir", "sty", "lap", "two", "ins", "con", "ant", "net", "tux", "ode", "stu", "mug", "cad", "nap", "gun", "fop", "tot", "sow", "sal", "sic", "ted", "wot", "del", "imp", "cob", "way", "ann", "tan", "mci", "job", "wet", "ism", "err", "him", "all", "pad", "hah", "hie", "aim"}
	board := [][]byte{{'a', 'a'}}
	words := []string{"aaa"}
	//wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Println(findWords(board, words))

}
