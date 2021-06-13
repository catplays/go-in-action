package concurrent

import (
	"fmt"
	"sync"
)

/**
并发安全的有序链表
 */

type IntList struct {
	head *IntNode		// head
	length int			// len of list
	mu sync.RWMutex		// lock
}

func newIntNode (value int) *IntNode {
	return &IntNode{
		value: value,
	}
}

func NewIntList() *IntList {
	return &IntList{
		head: newIntNode(0),
	}
}

func (list IntList) Contains(value int) bool  {
	list.mu.RLock()
	defer list.mu.RUnlock()
	node := list.head
	for node != nil && node.value < value {
		node = node.next
	}
	return node!= nil && node.value == value
}

func (list *IntList) Insert (value int) bool {
	list.mu.Lock()
	defer list.mu.Unlock()
	node := list.head
	pre := node
	for node != nil && node.value < value {
		node = node.next
		pre = node
	}
	if node != nil && node.value == value {
		return false
	}
	one := newIntNode(value)
	pre.next = one
	one.next = node
	list.length++
	return true
}

func (list *IntList) Delete (value int) bool {
	list.mu.Lock()
	defer list.mu.Unlock()
	node := list.head
	pre := node
	for node != nil && node.value < value {
		node = node.next
		pre = node
	}
	if node == nil ||  node.value != value {
		return false
	}
	pre.next = node.next
	node.next = nil
	list.length--
	return true
}
func (list *IntList) Len() int {
	return list.length
}

func (list *IntList) Range(f func(value int) bool)  {
	list.mu.RLock()
	defer list.mu.RUnlock()
	node := list.head.next
	for node != nil {
		if !f(node.value) {
			break
		}
		fmt.Println(node.value)
	}
}
