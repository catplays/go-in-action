package main

import "fmt"


func main() {
	fmt.Println("hello world")
	fmt.Println(IsIsomorphic("asd","sdf"))
}

func IsIsomorphic(s string, t string) bool {
	set1 := map[byte] byte{}
	for i:= range s {
		a := s[i]
		val,exist := set1[a]
		// 如果不存在，赋值
		if !exist {
			// 因为关系是一一绑定的，这里判断t[i] 有没有被绑定过
			val2, tExts := set1[t[i]]
			if tExts && s[i] != val2 {
				return false
			}
			set1[a] = t[i]
		} else if val!= t[i]{
			return false
		}

	}

	return true
}