package list_node

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildListNode(list []int) *ListNode {
	head := &ListNode{}
	pre := head
	for _,val := range list {
		node := &ListNode{
			Val: val,
		}
		pre.Next = node
		pre = node
	}
	return head.Next
}