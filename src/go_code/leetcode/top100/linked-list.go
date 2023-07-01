package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//linked-list
//linked-list
func reverseKGroup(head *ListNode, k int) *ListNode {
	last_tail := head
	temp := head
	index := 0
	pre := head
	//第一段
	for ; index < k && temp != nil; index++ {
		pre = temp
		temp = temp.Next
	}
	if index < k {
		return head
	}
	pre.Next = nil
	result := reverseList(head)
	//循坏实现后面的段
	for temp != nil {
		index = 0
		this_tail := temp

		pre = temp
		current := temp
		for ; index < k && current != nil; index++ {
			pre = current
			current = current.Next
		}
		temp = current
		if index < k {
			last_tail.Next = this_tail
			return result
		}
		pre.Next = nil

		last_tail.Next = reverseList(this_tail)
		last_tail = this_tail
	}
	return result
}
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre := head
	next := head.Next
	for next != nil {
		temp := next.Next
		next.Next = pre
		pre = next
		next = temp
	}
	head.Next = nil
	return pre
}

//Definition for singly-linked list.

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	//添加一个空头，便于处理
	H := &ListNode{}
	H.Next = head
	//先找到翻转起始位置
	start := H
	for i := 0; i < left; i++ {
		start = start.Next
	}
	L := start.Next
	//然后从left到right开始翻转
	term := L   //中间节点
	R := L.Next //下一个节点
	pre := L    //记录前一个节点
	for i := left; i < right; i++ {
		term = R
		R = term.Next
		term.Next = pre
		pre = term
	}
	start.Next = term
	L.Next = R
	return H.Next
}
func createList(nums []int) *ListNode {
	head := &ListNode{}
	term := head
	for _, value := range nums {
		newnode := new(ListNode)
		newnode.Val = value
		term.Next = newnode
		term = term.Next
	}
	return head
}
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	head := &ListNode{}
	count := 0
	term := head
	flag := make([]bool, len(lists))
	for index := range lists {
		if lists[index] == nil {
			flag[index] = true
			count++
		}
	}
	for count < len(lists) {
		min := &ListNode{Val: 1e9}
		pos := 0
		for index := range lists {
			if !flag[index] { //还没空
				if min.Val > lists[index].Val {
					min = lists[index]
					pos = index
				}
			}
		}
		term.Next = min
		term = term.Next
		lists[pos] = lists[pos].Next
		if lists[pos] == nil {
			flag[pos] = true
			count++
		}
	}
	return head.Next
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	flag := false
	slow, fast := head, head.Next
	for fast != nil { //快慢指针
		if slow == fast {
			flag = true
			break
		}
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	if !flag { //表示没有环
		return nil
	}
	for start := head; start != slow; start = start.Next {
		count := 0
		for next := start.Next; count < 2; next = next.Next {
			if next == start { //找到了位置
				return start
			}
			if next == slow { //第一次经过slow，两次经过slow，说明进圈了但是没有找到，循环退出
				count++
			}
		}
	}
	return slow
}

func reorderList(head *ListNode) {
	length := 0
	for term := head; term != nil; term = term.Next {
		length++
	}
	if length <= 2 {
		return
	}
	Larray := make([]*ListNode, length/2)
	for term, index := head, 0; index < length; index, term = index+1, term.Next {
		if index%2 == 1 {
			Larray[index/2] = term
		}
	}
	for term, index := head, length/2-1; index >= 0; index = index - 1 {
		temp := term.Next
		term.Next = Larray[index]
		term.Next.Next = temp
		term = temp
	}
	if length%2 == 1 {
		Larray[0].Next = nil
	} else {
		Larray[0].Next.Next = nil
	}
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre, current, forward := head, head, head
	for n > 0 {
		forward = forward.Next
		n--
	}
	for forward != nil {
		pre = current
		current = current.Next
		forward = forward.Next
	}
	if pre == current {
		return head.Next
	}
	pre.Next = current.Next

	return head
}
func sortList(head *ListNode) *ListNode {
	length := 0
	for temp := head; temp != nil; temp = temp.Next {
		length++
	}
	if length <= 1 {
		return head
	}
	H := head
	for i := length; i > 0; i-- {
		temp := H.Next
		if i > 1 && H.Val > H.Next.Val {
			H.Next = temp.Next
			temp.Next = H
			H = temp
		}
		pre := H
		for j := 1; j < i; j++ {
			if temp.Next != nil && temp.Val > temp.Next.Val {
				term := temp.Next
				temp.Next = term.Next
				term.Next = temp
				pre.Next = term
				pre = pre.Next
			}
		}
	}
	return H
}
