package geechache

import (
	"fmt"
	"log"
	"sync"
)

type Getter interface {
	Get (key string) ([]byte, error)
}

type Group struct {
	name string
	getter Getter
	mainCache cache
}
var (
	mux sync.RWMutex
	groups = make(map[string]*Group)
)


type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}



func NewGroup(name string, byteCache int64,getter Getter ) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mux.Lock()
	defer mux.Unlock()
	group := &Group{
		name: name,
		getter: getter,
		mainCache: cache{
			cacheBytes:byteCache,
		},
	}
	groups[name] = group
	return group
}

func GetGroup(name string) *Group  {
	mux.Lock()
	defer mux.Unlock()
	return groups[name]
}

func (g Group) Get(key string)(ByteView, error)  {

	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}

	if v,ok := g.mainCache.get(key); ok {
		log.Printf("cache hit")
		return v,nil
	}
	return g.load(key)
}

func (g Group) load(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView{
		copyByte(bytes),
	}
	g.populate(key,value)
	return value, nil
}

func (g Group) populate(key string, value ByteView) {
	g.mainCache.add(key, value)
}
