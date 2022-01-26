package algorithm

import "math"

//209. 长度最小的子数组
func minSubArrayLen(target int, nums []int) int {
	min := math.MaxInt32
	left, right := 0, 0
	length := len(nums)
	sum := 0
	for right < length {
		sum += nums[right]
		for sum >= target {
			min = minVal(min, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if min == math.MaxInt32 {
		return 0
	}
	return min
}
func minVal(x, y int) int {
	if x > y {
		return y
	}
	return x
}
