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
	data := make(map[rune]struct{},0)
	for _, str := range []rune(strs) {
		data[str] = struct{}{}
	}
	n := len(s)
	left, right := int32(0), int32(n-1)
	arr := []rune(s)
	for left < right {
		_, ok := data[arr[left]]
		if  !ok {
			left++
			continue
		}
		for left< right {
			_, ok = data[arr[right]]
			if ok {
				break
			}
			right--
		}
		arr[left],arr[right] = arr[right], arr[left]
		right--
		left++
	}
	return string(arr)
}
