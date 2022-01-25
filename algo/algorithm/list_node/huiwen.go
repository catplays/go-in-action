package list_node

// 判断是否为回文链表
func isPalindrome(head *ListNode) bool {
	arr := make([]*ListNode, 0)
	for head != nil {
		arr = append(arr,head)
		head = head.Next
	}
	length := len(arr)/2
	for i:=0;i<length;i++ {
		if arr[i].Val != arr[len(arr)-i-1].Val {
			return false
		}
	}
	return true
}

