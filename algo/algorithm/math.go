package algorithm

func CanFinish(numCourses int, prerequisites [][]int) bool {
	var keyMap map[int] []int = make(map[int] []int)
	var n int = len(prerequisites)
	for i :=0; i<n; i++ {
		keyMap[prerequisites[i][0]] = prerequisites[i]
	}
	var count int = 0
	var i int = 0
	// 作为待学习的集合
	var set map[int] int = make(map[int] int)
	var key = prerequisites[i][0]
	for count<n {
		_, ok := set[key]
		// 存在那么首尾相交了
		if ok {
			return false
		}
		set[key] = 0
		var one []int = keyMap[key]
		// 拿到下一个课程
		key = one[1]
		count++
	}
	return true
}

/**
https://leetcode-cn.com/problems/isomorphic-strings/
同构字符串
 */
func IsIsomorphic(s string, t string) bool {
	set1 := map[byte] int{}
	for i:= range s {
		val,exist := set1[s[i]]
		if !exist {
			set1[s[i]] = 1
		} else {
			set1[s[i]] = val+1
		}
	}
	set2 := map[byte] int{}
	for i := range t {
		val,exist := set2[t[i]]
		if !exist {
			set2[t[i]] = 1
		} else {
			set2[t[i]] = val+1
		}
	}
	for i:= range s {
		if set1[s[i]] != set2[t[i]]  {
			return false
		}
	}
	return true
}