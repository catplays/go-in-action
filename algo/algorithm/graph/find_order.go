package graph

// https://leetcode-cn.com/problems/course-schedule-ii/
// 课程表210
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		// 描述一个节点有多少条连接边
		edges = make([][]int, numCourses)
		// 描述节点的状态，0表示未搜索，1表示搜索中，2表示搜索完成
		visited      = make([]int, numCourses)
		valid   bool = true
		result  []int
		dfs     func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1

		for _,v:= range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}

		visited[u] = 2
		result = append(result, u)

	}

	for _, info := range prerequisites {
		// 边的后一个节点
		i := info[1]
		edges[i] = append(edges[i], info[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return []int{}
	}
	// result中是倒序排的,反转下
	for i:=0 ;i<numCourses/2;i++ {
		result[i],result[numCourses-i-1] = result[numCourses-i-1],result[i]
	}
	return result
}
