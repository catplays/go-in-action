package math


/**
https://leetcode-cn.com/problems/1-bit-and-2-bit-characters/
717. 1比特与2比特字符
 */
func isOneBitCharacter(bits []int) bool {
	i:=0
	for ;i<len(bits)-1; {
		if bits[i] == 0 {
			i++
		} else  {
			i+=2
		}
	}
	return i==len(bits)-1
}

/**
693. 交替位二进制数
 */
func hasAlternatingBits(n int) bool {
	pre := n%2
	n = n >> 1
	for n>0 {
		k := n%2
		if pre^k ==0 {
			return false
		}
		pre = k
		n = n >> 1
	}
	return true
}