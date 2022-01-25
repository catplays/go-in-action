package list_node

import "math/rand"

/**
给你一个单链表，随机选择链表的一个节点，并返回相应的节点值。每个节点 被选中的概率一样 。

 */
type Solution struct {
	data []int
	length int
}


func Constructor(head *ListNode) Solution {
	so := Solution{

	}
	if head == nil {
		return so
	}
	len := 0
	data := make([]int, 0)
	for head != nil {
		data = append(data, head.Val)
		len++
		head = head.Next
	}
	so.data = data
	so.length = len
	return so
}


func (this *Solution) GetRandom() int {
	index := rand.Intn(this.length)
	return this.data[index]
}
