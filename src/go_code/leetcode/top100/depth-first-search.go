package main

/***************************************************/
/***************************************************/
///depth-first-search
/***************************************************/
/***************************************************/
// func is_noway(grid [][]byte, i, j int) bool { //判断是否有相邻情况
// 	if i > 0 && grid[i-1][j] == '0' {
// 		return false
// 	}
// 	if i < len(grid)-1 && grid[i+1][j] == '0' {
// 		return false
// 	}
// 	if j > 0 && grid[i][j-1] == '0' {
// 		return false
// 	}
// 	if j < len(grid[0])-1 && grid[i][j+1] == '0' {
// 		return false
// 	}
// 	return true
// }
func numIslands(grid [][]byte) int { //(先将’1‘的位置记录下来，然后依次在每个记录位置进行深度优先遍历，遍历过程中将’1‘变为’0‘)
	index_set := [][2]int{}
	for i := range grid[0] {
		if i == 0 {
			if grid[0][0] == '1' {
				index_set = append(index_set, [2]int{0, 0})
			}
		} else {
			if grid[0][i] == '1' && grid[0][i-1] != '1' {
				index_set = append(index_set, [2]int{0, i})
			}
		}
	}
	for i := 1; i < len(grid); i++ {
		for j := range grid[0] {
			if j == 0 {
				if grid[i-1][j] != '1' && grid[i][j] == '1' {
					index_set = append(index_set, [2]int{i, j})
				}
			} else {
				if grid[i-1][j] != '1' && grid[i][j] == '1' && grid[i][j-1] != '1' {
					index_set = append(index_set, [2]int{i, j})
				}
			}
		}
	}
	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		// if is_noway(grid, i, j) {
		// 	return
		// }
		if i > 0 && grid[i-1][j] == '1' {
			grid[i-1][j] = '0'
			dfs(grid, i-1, j)
		}
		if i < len(grid)-1 && grid[i+1][j] == '1' {
			grid[i+1][j] = '0'
			dfs(grid, i+1, j)
		}
		if j > 0 && grid[i][j-1] == '1' {
			grid[i][j-1] = '0'
			dfs(grid, i, j-1)
		}
		if j < len(grid[0])-1 && grid[i][j+1] == '1' {
			grid[i][j+1] = '0'
			dfs(grid, i, j+1)
		}
	}
	result := 0
	for index := range index_set {
		i, j := index_set[index][0], index_set[index][1]
		if grid[i][j] == '1' {
			result++
			grid[i][j] = '0'
			dfs(grid, i, j)
		}
	}

	return result
}
