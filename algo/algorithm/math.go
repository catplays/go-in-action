package algorithm

import (
	"fmt"
	"strconv"
)

func Abs (val int) int {
	if val < 0 {
		return -val
	}
	return val
}

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

func summaryRanges(nums []int) []string {
	result := make([]string, 0)
	if len(nums) == 0 {
		return result
	}
	start := nums[0]
	pre := start
	for i:=1 ;i<len(nums);i++ {
		if nums[i] == pre+1 {
			pre = nums[i]
			continue
		}
		if start == pre {
			result = append(result, strconv.Itoa(start))
		} else {
			result = append(result, fmt.Sprintf("%d->%d",start, pre))
		}
		start = nums[i]
		pre = start
	}
	if  start == pre {
		result = append(result, strconv.Itoa(start))
	} else {
		result = append(result, fmt.Sprintf("%d->%d",start, pre))
	}
	return result
}

//201. 数字范围按位与
func rangeBitwiseAnd(left int, right int) int {
	// 找出left和right的最长公共前缀[k]，left到right的直接必定要经过[k][0...0],求&之后就是公共前缀
	k := 1<<30
	ans := 0
	for k>0 && left&k==right&k {
		ans |= left&k
		k >>=1
	}
	return ans
}