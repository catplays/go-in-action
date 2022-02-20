package algorithm


/**
https://leetcode-cn.com/problems/1-bit-and-2-bit-characters/
717. 1比特与2比特字符
 */
func isOneBitCharacter(bits []int) bool {
	length :=  len(bits)
	if length == 0 {
		return false
	}
	if bits[length-1] != 0 {
		return false
	}
	var res = false
	var  oneBitCharacter func(bits []int,i int)
	oneBitCharacter = func(bits []int,i int) {
		if res {
			return
		}
		if i < 0 {
			res = true
			return
		}
		if bits[i] == 0 {
			oneBitCharacter(bits, i-1)
		}
		if i-1>=0 && bits[i-1] == 1 {
			oneBitCharacter(bits, i-2)
		}
	}
	oneBitCharacter(bits, length-2)
	return res
}
