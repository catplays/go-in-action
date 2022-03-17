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