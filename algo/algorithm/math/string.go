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
