package datastructure

import "fmt"

/**
请你设计一个用于存储字符串计数的数据结构，并能够返回计数最小和最大的字符串。

实现 AllOne 类：

AllOne() 初始化数据结构的对象。
inc(String key) 字符串 key 的计数增加 1 。如果数据结构中尚不存在 key ，那么插入计数为 1 的 key 。
dec(String key) 字符串 key 的计数减少 1 。如果 key 的计数在减少后为 0 ，那么需要将这个 key 从数据结构中删除。测试用例保证：在减少计数前，key 存在于数据结构中。
getMaxKey() 返回任意一个计数最大的字符串。如果没有元素存在，返回一个空字符串 "" 。
getMinKey() 返回任意一个计数最小的字符串。如果没有元素存在，返回一个空字符串 "" 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/all-oone-data-structure
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

key:hash+双向链表
 */
type AllOne struct {
	key2Count map[string]int  // key的计数
	count2Node map[int]*ListNode // 计数对应的keys
	minCount int
	maxCount int
}


type ListNode struct {
	Count int // 当前计数
	Keys map[string]struct{} // 计数对应的keys
	Pre *ListNode
	Next *ListNode
}

func Constructor() AllOne {
	return AllOne{
		key2Count:  make(map[string]int,0),
		count2Node: make(map[int]*ListNode, 0),
		minCount:   0,
		maxCount:   0,
	}
}


func (this *AllOne) Inc(key string)  {
	val, ok := this.key2Count[key]
	if !ok {
		val = 0
	}

	this.key2Count[key] = val+1
	/**
	- 更新key2count
	- 删除val计数List中的key,判断map是否为空，为空则移除
	- 添加val+1计数中的key
	- 处理最小值，当前最小值是val，则判断val计数是否为空，为空则是val+1
	- 处理最大值，当前最大值是val，是则最大值为val+1
	*/
	node,ok := this.count2Node[val+1]
	if !ok {
		node = &ListNode{
			Count: val+1,
			Keys:  make(map[string]struct{}, 0),
			Pre:   nil,
			Next:  nil,
		}
	}
	node.Keys[key] = struct{}{}
	this.count2Node[val+1] = node

	if val >0  {
		beforeNode,_ := this.count2Node[val]
		// 删除上个频次里的key信息
		delete(beforeNode.Keys, key)
		if len(beforeNode.Keys) == 0 {// 为空了，则更新前后节点
			if beforeNode.Pre != nil {
				beforeNode.Pre.Next = node
			}
			node.Pre = beforeNode.Pre
			// 删除平次节点数据
			delete(this.count2Node, val)
		} else {// val对应的节点不空，更新next指针
			beforeNode.Next = node
			node.Pre = beforeNode
		}
		// 处理最小值,当前最小值是val，则判断val计数是否为空，为空则是val+1
		if this.minCount == val && len(beforeNode.Keys) ==0 {
			this.minCount = val+1
		}
		// 处理最大值，当前最大值是val，是则最大值为val+1
		if this.maxCount == val {
			this.maxCount = val+1
		}
	} else {
		head,ok := this.count2Node[this.minCount]
		if ok {
			node.Next = head
			head.Pre = node
		}

		//	- 处理最小值，1
		//	- 处理最大值，max(max,1)
		this.minCount = 1
		this.maxCount = max(this.maxCount,1)
	}

}

func (this *AllOne) Dec(key string)  {
	val, ok := this.key2Count[key]
	if !ok {
		return
	}
	if val == 1 {
		delete(this.key2Count, key)
		node,_ := this.count2Node[1]
		delete(node.Keys, key)
		if len(node.Keys) == 0 {
			if node.Next == nil {
				this.minCount = 0
				this.maxCount = 0
			} else {
				this.minCount = node.Next.Count
				node.Next.Pre = nil
			}
			delete(this.count2Node,1)
		}

	} else {
		// 1. 更新key2count
		this.key2Count[key] = val-1
		// 2. 更新后的节点信息
		node,ok := this.count2Node[val-1]
		if !ok {
			node = &ListNode{
				Count: val-1,
				Keys:  make(map[string]struct{}, 0),
				Pre:   nil,
				Next:  nil,
			}
			this.minCount = min(this.minCount, val-1)
		}
		node.Keys[key] = struct{}{}
		this.count2Node[val-1] = node

		// 3. 删除val频次中的key 信息
		beforeNode,_ := this.count2Node[val]
		delete(beforeNode.Keys, key)

		if len(beforeNode.Keys) == 0 {// 为空了，则更新前后节点
			if beforeNode.Pre != nil {
				// 如果上一个节点就是当前val-1频次
				if beforeNode.Pre == node {
					beforeNode.Pre.Next = beforeNode.Next
				} else {
					beforeNode.Pre.Next = node
				}
			}
			// 更新后节点的关系
			if beforeNode.Next != nil {
				beforeNode.Next.Pre = node
			}
			// 删除平次节点数据
			delete(this.count2Node, val)

			if this.maxCount == val {
				this.maxCount = val-1
			}
			if this.minCount == val {
				this.maxCount = val-1
			}
		} else {// val对应的节点不空，pre指针
			beforeNode.Pre = node
		}

	}

}


func (this *AllOne) GetMaxKey() string {
	if this.maxCount == 0 {
		return ""
	}
	node,ok := this.count2Node[this.maxCount]
	if !ok {
		fmt.Println("GetMaxKey err")
		return ""
	}
	for key,_:= range node.Keys {
		return key
	}
	return ""
}


func (this *AllOne) GetMinKey() string {
	if this.minCount == 0 {
		return ""
	}

	node,ok := this.count2Node[this.minCount]
	if !ok {
		fmt.Println("GetMinKey err")
		return ""
	}
	for key,_:= range node.Keys {
		return key
	}
	return ""
}

func max(x,y int) int {
	if x>y {
		return x
	}
	return y
}
func min(x,y int) int {
	if x<y {
		return x
	}
	return y
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
