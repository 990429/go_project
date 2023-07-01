package main

/***********************************/
//binary-tree
/***********************************/
/***********************************/

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	l_to_r := false //从左到右取
	queue := []*TreeNode{root}
	r := 0

	for len(queue) > 0 {
		r = len(queue) - 1
		res := []int{}
		for i := range queue {
			res = append(res, queue[i].Val)
		}
		for i := r; i >= 0; i-- {

			if l_to_r {
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
			} else {
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
			}
		}
		l_to_r = !l_to_r
		ans = append(ans, res[:])
		queue = queue[r+1:]
	}
	return
}
func create_tree(tree []int) *TreeNode { //使用队列进行创建
	if len(tree) == 0 {
		return nil
	}
	root := &TreeNode{Val: tree[0]}
	queue := []*TreeNode{root}
	var newnode *TreeNode
	for index := 1; index < len(tree); {
		temp := queue[0]
		if tree[index] >= 0 {
			newnode = &TreeNode{Val: tree[index]}
			temp.Left = newnode
			queue = append(queue, newnode)
		}
		index++
		if index < len(tree) && tree[index] >= 0 {
			newnode = &TreeNode{Val: tree[index]}
			temp.Right = newnode
			queue = append(queue, newnode)
		}
		index++
		queue = queue[1:]
	}
	return root
}
func search_tree(root *TreeNode, target int) *TreeNode {
	if root != nil {
		if root.Val == target {
			return root
		}
		left := search_tree(root.Left, target)
		right := search_tree(root.Right, target)
		if left != nil && left.Val == target {
			return left
		} else {
			return right
		}
	}
	return root

}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode { //使用深度优先搜索分别搜索两个节点，记录路径，然后比对路径
	is_found := false
	path_p, path_q, target_path := []*TreeNode{}, []*TreeNode{}, []*TreeNode{}
	var dfs func(sub_root, target *TreeNode, path []*TreeNode)
	dfs = func(sub_root, target *TreeNode, path []*TreeNode) {
		if sub_root == target {
			target_path = make([]*TreeNode, len(path)+1)
			copy(target_path, append(path, sub_root))

			is_found = true
			return
		}
		if is_found {
			return
		}
		if sub_root != nil {
			dfs(sub_root.Left, target, append(path, sub_root))
			dfs(sub_root.Right, target, append(path, sub_root))
		}
	}
	path := []*TreeNode{}
	dfs(root, p, path)
	path_p = make([]*TreeNode, len(target_path))
	copy(path_p, target_path)
	is_found = false
	path = []*TreeNode{}
	dfs(root, q, path)
	path_q = make([]*TreeNode, len(target_path))
	copy(path_q, target_path)
	len_q, len_p := len(path_q), len(path_p)
	result := root
	for index_q := len_q - 1; index_q >= 0; index_q-- {
		flag := false
		for index_p := len_p - 1; index_p >= 0; index_p-- {
			if path_q[index_q] == path_p[index_p] {
				flag = true
				result = path_q[index_q]
				break
			}
		}
		if flag {
			break
		}
	}
	return result
}
func maxPathSum(root *TreeNode) int {
	res := -10000
	var dfs func(sub_root *TreeNode) int
	dfs = func(sub_root *TreeNode) int {
		if sub_root == nil {
			return 0
		}
		leftMax := max(dfs(sub_root.Left), 0)
		rightMax := max(dfs(sub_root.Right), 0)
		new_value := sub_root.Val + leftMax + rightMax
		res = max(res, new_value)
		return sub_root.Val + max(leftMax, rightMax)
	}
	dfs(root)
	return res
}
func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	temp := make(map[int]int)
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if _, ok := temp[level]; !ok {
			temp[level] = node.Val
		}
		if node.Right != nil {
			dfs(node.Right, level+1)
		}
		if node.Left != nil {
			dfs(node.Left, level+1)
		}
	}
	dfs(root, 0)
	for index := 0; index < len(temp); index++ {
		ans = append(ans, temp[index])
	}
	return
}
