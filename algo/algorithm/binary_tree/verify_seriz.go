package binary_tree

import (
	"fmt"
	"strings"
)

var (
	valid = true
)

/**
331. 验证二叉树的前序序列化
序列化二叉树的一种方法是使用 前序遍历 。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 #。
*/

func isValidSerialization(preorder string) bool {
	strs := strings.Split(preorder, ",")
	var preOrder func(strs []string)
	index := 0
	preOrder = func(strs []string) {
		if index >= len(strs) || strs[index] == "#" {
			return
		}
		index++
		preOrder(strs)
		index++
		preOrder(strs)
	}

	preOrder(strs)
	return index == len(strs)-1
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {

	count := make(map[string]int, 0)
	res := make([]*TreeNode, 0)
	var findByPreOrder func(node *TreeNode) string
	findByPreOrder = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		key := fmt.Sprintf("%d,%s,%s", node.Val,
			findByPreOrder(node.Left),
			findByPreOrder(node.Right))
		val, ok := count[key]
		if ok {
			count[key] = val + 1
		} else {
			count[key] = 1
		}
		if count[key] == 2 {
			res = append(res, node)
		}
		return key
	}
	findByPreOrder(root)
	return res
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	if root.Left == nil && root.Right == nil {
		return fmt.Sprintf("%d", root.Val)
	}
	if root.Left == nil {
		return fmt.Sprintf("%d()(%s)", root.Val, tree2str(root.Right))
	}
	if root.Right == nil {
		return fmt.Sprintf("%d(%s)", root.Val, tree2str(root.Left))
	}
	return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str(root.Left), tree2str(root.Right))
}


