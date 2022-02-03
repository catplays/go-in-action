package backtrack


type WordDictionary struct {
	data map[int]Data
}

type Data map[string]bool

func Constructor() WordDictionary {
	return WordDictionary{
		data: make(map[int]Data,0),
	}
}


func (this *WordDictionary) AddWord(word string)  {

	val,ok := this.data[len(word)]
	if ok {
		val[word] = true
		return
	}
	val = Data{}
	val[word] = true
	this.data[len(word)] = val
}


func (this *WordDictionary) Search(word string) bool {
	length := len(word)
	valMap,ok :=this.data[length]
	if !ok {
		return false
	}
	_, exist := valMap[word]
	if exist {
		return true
	}
	for key,_ := range valMap {
		i:=0
		for ;i<length;i++ {
			if word[i] == '.' {
				continue
			}
			if word[i] != key[i] {
				break
			}
		}
		if i== length {
			return true
		}
	}
	return false
}
