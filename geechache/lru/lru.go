package lru

import "container/list"

type Cache struct {
	maxBytes int64 // 允许使用的最大内存
	nBytes int64 // 已经使用的最大内存
	ll *list.List // go提供的双向链表
	cache map[string] *list.Element
	OnEvicted func(key string, value Value) // 某条数据被移除时的回调函数
}

type Entry struct {
	key string
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64 , onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		ll: list.New(),
		cache: make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c Cache) Get(key string) (Value,bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*Entry)
		return kv.value,true
	}
	return nil,false
}

func (c Cache) Add(key string, value Value)  {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*Entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&Entry{key: key, value: value})
		c.cache[key] = ele
		c.nBytes += int64(len(key) + value.Len())
	}
	if c.maxBytes !=0 && c.nBytes > c.maxBytes {
		c.removeOldest()
	}
}

func (c Cache) removeOldest()  {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*Entry)
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key) + kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c Cache) Len() int {
	return c.ll.Len()
}