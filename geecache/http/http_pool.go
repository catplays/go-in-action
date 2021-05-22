package http

import (
	"catwang.com/go-in-action/geecache"
	"catwang.com/go-in-action/geecache/consistenthash"
	"catwang.com/go-in-action/geecache/protobuf"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
	"strings"
	"sync"
)

const (
	defaultBasePath = "/_geecache/"
	defaultReplicas = 50
)
type HttpPool struct {
	self string
	basePath string
	mu sync.Mutex
	peers *consistenthash.Map
	httpGetter map[string]*HttpGetter

}

func NewHttpPool(self string) *HttpPool {
	return &HttpPool{
		self:     self,
		basePath: defaultBasePath,
	}
}
// 处理http请求
func (h *HttpPool) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	if !strings.HasPrefix(req.URL.Path, h.basePath) {
		log.Printf("httpServ unexpect path:%s", req.URL.Path)
		panic("httpServ unexpect path")
	}
	parts := strings.Split(req.URL.Path[len(h.basePath):], "/")
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	groupName := parts[0]
	key := parts[1]
	group := geecache.GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
		return
	}
	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write the value to the response body as a proto message.
	body, err := proto.Marshal(&protobuf.Response{
		Value: view.ByteSlice(),
	})

	if err != nil {
		http.Error(w, "no such key:"+key, http.StatusNotFound)
		return
	}
	log.Printf("search key:%v group:%v in:%v",key, groupName, h.self)
	w.Header().Set("Content-Type","application/octet-stream")
	w.Write(body)
}

// 实例化哈希一致性算法，并添加节点
func (h *HttpPool) Set(addrs ...string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.peers = consistenthash.New(defaultReplicas, nil)
	h.peers.Add(addrs...)
	h.httpGetter = make(map[string]*HttpGetter, len(addrs))
	for _, addr := range addrs {
		h.httpGetter[addr] = &HttpGetter{
			baseURL: addr+h.basePath,
		}
	}
}

// 包装了一致性哈希算法的Get方法，根据key找到一个缓存节点，并返回这个节点的httpClient
func (h *HttpPool) PickPeer(key string) (geecache.PeerGetter, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if peer := h.peers.Get(key); peer != "" && peer!= h.self {
		return h.httpGetter[peer], true
	}
	return nil, false
}