package main

import "strings"

type Node struct {
	pattern string // 待匹配的路由，用来绑定对应的处理器，也用来标识是否为叶子节点
	part string // 路由中的一部分
	children []*Node // 子节点
	isWild bool // 是否是模糊匹配的
}

func (node * Node) insert(pattern string, parts []string, height int)  {
	if len(parts) == height {
		node.pattern = pattern
		return
	}
	part := parts[height]
	child := node.matchChild(part)
	if child == nil {
		child = &Node{part: part, isWild: part[0] ==':'|| part[0] == '*'}
		node.children = append(node.children, child)
	}
	child.insert(pattern, parts, height+1)
}



func (node *Node) search(parts []string, height int)  *Node{
	// 是否是叶子节点 或者*
	if len(parts) == height || strings.HasPrefix(node.part, "*"){
		if node.pattern == "" {
			return nil
		}
		return node
	}
	part := parts[height]
	children := node.matchChildren(part)
	for _, child := range children {
		result := child.search(parts,height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

func (node * Node) matchChild(part string) *Node {
	for _, child := range node.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}
func (node *Node) matchChildren(part string) []*Node  {
	nodes := make([]*Node, 0)

	for _, node := range node.children {
		if node.part == part || node.isWild{
			  nodes = append(nodes, node)
		}
	}
	return nodes
}