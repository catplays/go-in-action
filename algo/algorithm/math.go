package algorithm

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func CanFinish(numCourses int, prerequisites [][]int) bool {
	var keyMap map[int][]int = make(map[int][]int)
	var n int = len(prerequisites)
	for i := 0; i < n; i++ {
		keyMap[prerequisites[i][0]] = prerequisites[i]
	}
	var count int = 0
	var i int = 0
	// 作为待学习的集合
	var set map[int]int = make(map[int]int)
	var key = prerequisites[i][0]
	for count < n {
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
	set1 := map[byte]int{}
	for i := range s {
		val, exist := set1[s[i]]
		if !exist {
			set1[s[i]] = 1
		} else {
			set1[s[i]] = val + 1
		}
	}
	set2 := map[byte]int{}
	for i := range t {
		val, exist := set2[t[i]]
		if !exist {
			set2[t[i]] = 1
		} else {
			set2[t[i]] = val + 1
		}
	}
	for i := range s {
		if set1[s[i]] != set2[t[i]] {
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
	for i := 1; i < len(nums); i++ {
		if nums[i] == pre+1 {
			pre = nums[i]
			continue
		}
		if start == pre {
			result = append(result, strconv.Itoa(start))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", start, pre))
		}
		start = nums[i]
		pre = start
	}
	if start == pre {
		result = append(result, strconv.Itoa(start))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", start, pre))
	}
	return result
}

//201. 数字范围按位与
func rangeBitwiseAnd(left int, right int) int {
	// 找出left和right的最长公共前缀[k]，left到right的直接必定要经过[k][0...0],求&之后就是公共前缀
	k := 1 << 30
	ans := 0
	for k > 0 && left&k == right&k {
		ans |= left & k
		k >>= 1
	}
	return ans
}

// 540. 有序数组中的单一元素
func singleNonDuplicate(nums []int) int {
	start, end := 0, len(nums)-1
	return singleNon(start, end, nums)
}

func singleNon(start int, end int, nums []int) int {
	if start == end {
		return nums[start]
	}
	if start+1 == end {
		return nums[start] ^ nums[end]
	}
	mid := (start + end) / 2
	return singleNon(start, mid, nums) ^ singleNon(mid+1, end, nums)
}

func luckyNumbers(matrix [][]int) []int {
	var length int = len(matrix)
	if length == 0 {
		return []int{}
	}
	var colLen int = len(matrix[0])

	var result []int
	for i := 0; i < length; i++ {
		min := math.MaxInt32
		idx := 0
		// 第i行的最小值
		for j := 0; j < colLen; j++ {
			if matrix[i][j] < min {
				min = matrix[i][j]
				idx = j
			}
		}

		// 找出idx列的最大值
		max := math.MinInt32
		for k := 0; k < length; k++ {
			if matrix[k][idx] > max {
				max = matrix[k][idx]
			}
		}
		if min == max {
			result = append(result, min)
		}
	}
	return result
}

/**
给你一个整数数组 nums 和两个整数k 和 t 。请你判断是否存在 两个不同下标 i 和 j，使得abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。
如果存在则返回 true，不存在返回 false。
链接：https://leetcode-cn.com/problems/contains-duplicate-iii
*/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	length := len(nums)
	for i := 0; i < k && i<length; i++ {
		for j := i + 1; j < k&& j<length; j++ {
			if Abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	start, end := 0, k-1
	next := end + 1
	for ; next < length; next++ {
		for i:=start; i<=end&& i<length;i++ {
			if Abs(nums[i]-nums[next]) <= t {
				return true
			}
		}
		start++
		end++
	}
	return false
}

/**
https://leetcode-cn.com/problems/rank-transform-of-an-array/
1331. 数组序号转换
 */
func arrayRankTransform(arr []int) []int {
	dataMap := make(map[int]int, 0)
	uniqueArr := make([]int, 0)

	for _,val := range arr {
		_, ok := dataMap[val]
		if !ok {
			dataMap[val] = 1
			uniqueArr = append(uniqueArr, val)
		}
	}
	sort.Sort(sort.IntSlice(uniqueArr))
	for i,val := range uniqueArr {
		dataMap[val] = i+1
	}
	res := make([]int,0)
	for _, val := range arr {
		index,_ := dataMap[val]
		res = append(res, index)
	}
	return res
}