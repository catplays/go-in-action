package math

func isUgly(n int) bool {
	for n>1 {
		if n %2 ==0 {
			n /=2
			continue
		}
		if n %3 == 0 {
			n /=3
			continue
		}
		if n % 5==0 {
			n /= 5
			continue
		}
		break
	}
	return n == 1
}

/**
假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错

链接：https://leetcode-cn.com/problems/first-bad-version

 */
func firstBadVersion(n int) int {
	start, end := 1,n
	mid := (start+end)/2
	for start<= end {
		// 是错误版本，往前找
		mid = (start+end)/2
		if isBadVersion(mid) {
			if !isBadVersion(mid-1) {
				return mid
			}
			end = mid-1
		} else {
			start = mid+1
		}

	}
	return mid
}
func isBadVersion(num int) bool {
	return true
}

func isPowerOfThree(n int) bool {
	for n>1 {
		if n%3 == 0 {
			n /=3
			continue
		}
		return false
	}
	return n==1
}

func isPowerOfFour(n int) bool {
	for n>1 {
		if n%4 == 0 {
			n /=4
			continue
		}
		return false
	}
	return n==1
}

/**
给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。
动态规划： f(x) = f(x-highBit) + 1,highBit表示2的幂数，比如4，那么f(5) = f(5-4)+1
其中5=101，4=100， 1=01；含义是：highBit的最高位是1，加上去掉最高位后剩下1的个数

*/
func countBits(n int) []int {
	highBit :=0
	res := make([]int,0)
	res = append(res, 0)
	for i:=1; i<=n;i++ {

		if i &(i-1) == 0 {
			highBit = i
		}
		res = append(res,res[i-highBit]+1)
	}
	return res
}