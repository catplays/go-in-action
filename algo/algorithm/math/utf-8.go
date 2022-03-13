package math

/**
https://leetcode-cn.com/problems/utf-8-validation/
给定一个表示数据的整数数组data，返回它是否为有效的 UTF-8 编码。

UTF-8 中的一个字符可能的长度为 1 到 4 字节，遵循以下的规则：

对于 1 字节的字符，字节的第一位设为 0 ，后面 7 位为这个符号的 unicode 码。
对于 n 字节的字符 (n > 1)，第一个字节的前 n 位都设为1，第 n+1 位设为 0 ，后面字节的前两位一律设为 10 。剩下的没有提及的二进制位，全部为这个符号的 unicode 码。
 */
func validUtf8(data []int) bool {
	length := len(data)
	for i := 0;i < length; {
		if data[i]>>7 == 0 {
			i++
			continue
		}
		n := 0
		j :=7
		for ;j>=0;j-- {
			val := data[i] & (1<<j)
			if val > 0 {
				n++
				continue
			}
			break
		}
		if n==1 || n>4 {
			return false
		}
		k:=i+1
		for ;k<i+n&& k<length;k++ {
			if data[k]>>6 == 2 {
				continue
			}
			break
		}
		if k != i+n {
			return false
		}
		i += n
	}
	return true
}
