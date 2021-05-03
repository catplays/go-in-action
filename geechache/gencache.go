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
	peers PeerPicker
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

func (g *Group) Get(key string)(ByteView, error)  {

	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}

	if v,ok := g.mainCache.get(key); ok {
		log.Printf("cache hit")
		return v,nil
	}
	return g.load(key)
}

func (g *Group) load(key string) (ByteView, error) {
	if g.peers != nil {
		if peerGetter, ok := g.peers.PickPeer(key); ok {
			if value, err := g.getFromPeer(key, peerGetter); err == nil {
				return value,nil
			}
			log.Println("[GeeCache] Failed to get from peer",)
		}
	}
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (ByteView, error)  {
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

func (g *Group) populate(key string, value ByteView) {
	g.mainCache.add(key, value)
}

func (g *Group) RegisterPeers(peer PeerPicker) {
	if g.peers != nil {
		panic("PeerPicker called more than one time")
	}
	g.peers = peer
}

func (g *Group) getFromPeer(key string, getter PeerGetter) (ByteView, error) {
	bytes, err := getter.Get(g.name, key)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: bytes}, nil
}

