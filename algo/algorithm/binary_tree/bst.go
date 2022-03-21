package binary_tree

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	var dfs func(root *TreeNode)
	data := make(map[int]struct{},0)
	res := false
	dfs = func(root *TreeNode) {
		if root == nil || res {
			return
		}
		val := k - root.Val
		_, ok := data[val]
		if ok {
			res = true
			return
		}
		data[root.Val] = struct{}{}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return res
}
