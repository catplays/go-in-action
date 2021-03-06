package geecache

import (
	"catwang.com/go-in-action/geecache/protobuf"
	"catwang.com/go-in-action/geecache/singleflight"
	"fmt"
	"log"
	"sync"
)

type Getter interface {
	Get (key string) ([]byte, error)
}

// 管理逻辑定义的缓存集群
type Group struct {
	name string
	getter Getter
	mainCache cache
	peers PeerPicker
	loader *singleflight.Group
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
		loader: &singleflight.Group{},
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

func (g *Group) load(key string) (byteView ByteView,err error) {
	val, err := g.loader.Do(key, func() (interface{}, error) {
		if g.peers != nil {
			if peerGetter, ok := g.peers.PickPeer(key); ok {
				if value, err := g.getFromPeer(key, peerGetter); err == nil {
					return value,nil
				}
				log.Println("[GeeCache] Failed to get from peer",)
			}
		}
		return g.getLocally(key)
	})
	if err == nil {
		return val.(ByteView), nil
	}
	return ByteView{}, err
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
	req := &protobuf.Request{
		Group: g.name,
		Key: key,
	}
	resp := &protobuf.Response{}
	err := getter.Get(req, resp)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: resp.Value}, nil
}

