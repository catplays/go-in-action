package backtrack

/**
https://leetcode-cn.com/problems/combination-sum-iii/
找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]

216
 */

func combinationSum3_1(k int, n int) [][]int {
	// 每一位数字选或者不选，一共有2^n个组合。用一个9位二进制数mask表示选择的结果，
	var arr []int
	check := func(mask int) bool {
		arr = nil
		sum := 0
		for i:=0;i<9;i++ {
			if 1<<i & mask >0 {//说明i这个位置选了数
				arr = append(arr, i+1)
				sum += i+1
			}
		}
		return sum == n && len(arr)==k
	}
	result := make([][]int,0)

	for i:=0 ;i<1<<9;i++{
		if check(i) {
			result = append(result, arr)
		}
	}
	return result
}

func combinationSum3(k int, n int) [][]int {

	var (
		result = make([][]int,0)
		temp []int
		combination func (curr int,sum int)
	)
	combination = func (curr int/*当前第几个数 */, sum int) {
		// 剪枝
		if len(temp) + 10-curr < k || sum > n {
			return
		}

		if len(temp) == k && sum == n {
			result = append(result, append([]int(nil), temp...))
			return
		}

		// 跳过当前元素
		combination(curr+1, sum)

		// 选择当前元素
		temp = append(temp, curr)
		combination(curr+1, sum+curr)
		temp = temp[:len(temp)-1]
	}
	combination(1,0)
	return result
}


