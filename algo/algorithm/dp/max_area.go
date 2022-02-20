package dp

/**
https://leetcode-cn.com/problems/maximal-square/
221. 最大正方形
key: 动态规划，dp[i][j] 表示以i行，j列的元素为正方形右下角区域里的最大边长
*/
func maximalSquare(matrix [][]byte) int {
	rowLen, colLen := len(matrix), len(matrix[0])
	dp := make([][]int, rowLen)
	maxSide := 0
	for i, row := range matrix {
		dp[i] = make([]int, len(row))
		for j, _ := range row {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}
	for i := 1; i < rowLen; i++ {
		for j := 1; j < colLen; j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j]))+1
				if maxSide< dp[i][j] {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide*maxSide
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
