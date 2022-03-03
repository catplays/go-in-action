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
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
				if maxSide < dp[i][j] {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Line struct {
	Start int
	End int
}
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	area1 := abs(ax2-ax1) * abs(ay2-ay1)
	area2 := abs(bx2-bx1) * abs(by2-by1)

	ax := &Line{}
	bx := &Line{}
	if ax1 < ax2 {
		ax.Start,ax.End = ax1,ax2
	} else {
		ax.Start,ax.End = ax2,ax1
	}
	if bx1 < bx2 {
		bx.Start, bx.End = bx1,bx2
	} else {
		bx.Start, bx.End = bx2,bx1
	}
	xlen :=0
	if ax.Start < bx.Start {
		xlen = getDupLen(ax,bx)
	} else {
		xlen = getDupLen(bx,ax)
	}


	ay := &Line{}
	by := &Line{}
	if ay1 < ay2 {
		ay.Start,ay.End = ay1,ay2
	} else {
		ay.Start,ay.End = ay2,ax1
	}
	if by1 < by2 {
		by.Start,by.End = by1,by2
	} else {
		by.Start,by.End = by2,by1
	}
	ylen :=0
	if ay.Start < by.Start {
		ylen = getDupLen(ay,by)
	} else {
		ylen = getDupLen(by,ay)
	}
	return area1+area2- xlen * ylen
}

func getDupLen(one, tow *Line) int {
	// 没有重合
	if one.End <= tow.Start {
		return 0
	}
	// 重合一部分
	if one.End > tow.Start&& one.End< tow.End {
		return one.End- tow.Start
	}
	return tow.End-tow.Start
}
