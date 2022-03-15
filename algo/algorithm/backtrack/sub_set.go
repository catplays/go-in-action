package backtrack

/**
https://leetcode-cn.com/problems/count-number-of-maximum-bitwise-or-subsets/
给你一个整数数组 nums ，请你找出 nums 子集 按位或 可能得到的 最大值 ，并返回按位或能得到最大值的 不同非空子集的数目 。
 */
func countMaxOrSubsets(nums []int) int {
	max := 0
	for _, val := range nums {
		max |= val
	}
	ans :=0
	var choose func (index int,res int, nums[]int)
	choose = func(index int, res int, nums[]int) {
		if index>= len(nums) {
			return
		}
		choose(index+1,res, nums)
		// 选择
		temp := res|nums[index]
		if temp == max {
			ans++
		}
		choose(index+1,temp, nums)
	}
	choose(0,0, nums)
	return ans
}

