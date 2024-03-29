package math

/**
编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
*/
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

/**
给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-arrays-ii
*/
func intersect(nums1 []int, nums2 []int) []int {
	data1 := make(map[int]int, 0)
	for _, val := range nums1 {
		count, ok := data1[val]
		if !ok {
			count = 0
		}
		data1[val] = count + 1
	}
	data2 := make(map[int]int, 0)
	res := make([]int, 0)
	for _, val := range nums2 {
		count, ok := data2[val]
		if !ok {
			count = 0
		}
		data2[val] = count + 1
		count1, ok := data1[val]
		if !ok {
			continue
		}
		if count1 >= count+1 {
			res = append(res, val)
		}
	}
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func imageSmoother(img [][]int) [][]int {
	n, m := len(img), len(img[0])
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)

		for j := 0; j < m; j++ {
			sum, num := 0, 0
			for _, c := range img[max(i-1, 0):min(i+2, n)] {
				for _, val := range c[max(j-1, 0):min(j+2, m)] {
					sum += val
					num++
				}
			}

			res[i][j] = sum / num
		}
	}
	return res
}

func thirdMax(nums []int) int {
	arr := make([]int, 0)
	for i := 0; i < len(nums); i++ {

		if len(arr) ==0 {
			arr = append(arr, nums[i])
			continue
		}
		// 比较
		for j := 0; j < len(arr); j++ {
			if arr[j] == nums[i] {
				break
			}
			// 来了最大的数，后移
			if nums[i] > arr[j] {
				if len(arr) <3 {
					arr = append(arr,0)
				}
				for k:=len(arr)-1;k>j;k-- {
					arr[k] = arr[k-1]
				}
				arr[j] = nums[i]
				break
			}
		}

	}
	if len(arr) > 2 {
		return arr[2]
	}
	return arr[0]
}
