package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)
//一致性hash
//节点管理
type Map struct {
	hash Hash 		// hash算法
	replicas int	//	虚拟节点的数量
	keys []int 		// sorted, 哈希环
	hashMap	map[int]string	// 通过hash值所在的虚拟节点，找真实节点名

}

type Hash func([]byte) uint32

// 添加节点
func New(replicas int, fn Hash) *Map {
	m := &Map{
		hash: fn,
		replicas: replicas,
		keys: make([]int, 0),
		hashMap: make(map[int]string),
	}
	if fn == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// 添加真实节点的名称到hash上
func (m *Map) Add(keys ...string)  {
	for _, key := range keys {
		for i:=0; i< m.replicas; i++ {
			hash := m.hash([]byte(strconv.Itoa(i)+ key))
			m.keys = append(m.keys, int(hash))
			m.hashMap[int(hash)] = key
		}
	}
	sort.Ints(m.keys)
}

// 根据key计算出的hash值，寻找离他最近的真实节点
func (m Map) Get(key string) string {
	hash := m.hash([]byte(key))
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i]>= int(hash)
	})
	return m.hashMap[m.keys[idx % len(m.keys)]]
}