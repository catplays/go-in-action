package math

func longestWord(words []string) string {
	data := make(map[string]struct{}, 0)
	res := ""
	for _, word := range words {
		data[word] = struct{}{}
	}
	for _, word := range words {
		flag := false
		for i := 1; i < len(word); i++ {
			_, ok := data[word[0:i]]
			if !ok {
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		if len(res) < len(word) {
			res = word
		} else if len(res) == len(word) {
			for i := 0; i < len(res); i++ {
				if res[i] < word[i] {
					break
				} else if res[i] > word[i] {
					res = word
				}
			}

		}
	}
	return res
}

func reverseVowels(s string) string {
	strs := "AEIOUaeiou"
	data := make(map[rune]struct{}, 0)
	for _, str := range []rune(strs) {
		data[str] = struct{}{}
	}
	n := len(s)
	left, right := int32(0), int32(n-1)
	arr := []rune(s)
	for left < right {
		_, ok := data[arr[left]]
		if !ok {
			left++
			continue
		}
		for left < right {
			_, ok = data[arr[right]]
			if ok {
				break
			}
			right--
		}
		arr[left], arr[right] = arr[right], arr[left]
		right--
		left++
	}
	return string(arr)
}

func canConstruct(ransomNote string, magazine string) bool {

	data := make(map[uint8]int, 0)
	length := len(magazine)
	for i := 0; i < length; i++ {
		key := magazine[i] - 'a'
		count, ok := data[key]
		if !ok {
			count = 0
		}
		data[key] = count + 1
	}

	for i := 0; i < len(ransomNote); i++ {
		key := ransomNote[i] - 'a'
		count, ok := data[key]
		if !ok {
			return false
		}
		count--
		if count <= 0 {
			delete(data, key)
			continue
		}
		data[key] = count
	}
	return true
}

func isSubsequence(s string, t string) bool {
	idx := 0
	n, nt := len(s), len(t)
	for i := 0; i < n; {
		if idx >= nt || nt-idx < n-i {
			return false
		}
		if s[i] != t[idx] {
			idx++
			continue
		}
		idx++
		i++
	}
	return true
}


func longestPalindrome(s string) int {
	data := make(map[int32]int,0)
	for _,val := range s {
		v, ok := data[val-'a']
		if ok {
			data[val-'a'] = v+1
		} else {
			data[val-'a'] = 1
		}
	}
	flag := false
	count := 0
	for _,val := range data {
		if val %2 ==0 {
			count+=val
		} else {
			count+=val-1
			flag = true
		}
	}
	if flag {
		count++
	}
	return count
}