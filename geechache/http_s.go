package geechache

import (
	"log"
	"net/http"
	"strings"
)

const defaultBasePath = "/_geecache/"
type HttpPool struct {
	self string
	basePath string
}

func NewHttpPool(self string) *HttpPool {
	return &HttpPool{
		self: self,
		basePath:defaultBasePath,
	}
}

func (h *HttpPool) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	if !strings.HasPrefix(req.URL.Path, h.basePath) {
		log.Printf("httpServ unexpect path:%s", req.URL.Path)
		panic("httpServ unexpect path")
	}
	parts := strings.Split(req.URL.Path[len(h.basePath):], "/")
	if len(parts) !=2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	groupName := parts[0]
	key := parts[1]
	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
		return
	}
	view, err := group.Get(key)
	if err != nil {
		http.Error(w, "no such key:"+key, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type","application/octet-stream")
	w.Write(view.ByteSlice())
}