package stack

import "math"

/**
https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
84. 柱状图中最大的矩形
key: 单调栈 以i为中心，向左向右找连续大于等于height[i]的位置，计算面积
*/
func largestRectangleArea(heights []int) int {
	n:=len(heights)
	var left,right = make([]int,n),make([]int,n)
	var stack []int
	for i, height := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= height {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	res := math.MinInt32
	for i := 0; i < n; i++ {
		res = max(res, (right[i] - left[i] - 1) * heights[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}


/**
https://leetcode-cn.com/problems/maximal-rectangle/
85. 最大矩形
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
key: 将每一层之上看做柱状图，求最大的值
 */
func maximalRectangle(matrix [][]byte) int {
	heights := make([]int, len(matrix[0]))
	res := 0
	for _,row := range matrix {
		for j, col := range row {
			if col == 1 {
				heights[j] ++
			} else {
				heights[j] = 0
			}
		}
		res = max(res, largestRectangleArea(heights))
	}
	return res
}