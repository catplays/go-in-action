package datastructure

/**
基于go 实现的hashmap
 */
type HashMap struct {
	key string
	val string
	hashCode int
	next *HashMap
}
/**
hashmap 数据存储
 */
var table [16]*HashMap

/**
初始化表
 */
func initTable() {
	for i := range table{
		table[i] = &HashMap{"","", 0, nil}
	}
}

func getInstance() [16]*HashMap {
	if table[0] == nil {
		initTable()
	}
	return table
}

func genHashCode(k string) int {
	if len(k) == 0 {
		return 0
	}
	var hashCode int = 0
	for i := range k{
		hashCode += int(k[i])
	}
	return hashCode
}

func indexOfTable(hashCode int) int  {
	return hashCode % 16
}

func put(key string, val string)  {
	var hashCode int = genHashCode(key)
	var thisNode = HashMap{key,val,hashCode, nil}
	var tableIndex = indexOfTable(hashCode)
	var table [16]*HashMap = getInstance()
	var headNode= table[tableIndex]
	if (*headNode).key == "" {
		*headNode = thisNode
		return
	}
}
