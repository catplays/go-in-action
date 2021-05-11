package main

import (
	"catwang.com/go-in-action/geecache"
	http2 "catwang.com/go-in-action/geecache/http"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	db = map[string]string{
		"Tom":  "630",
		"Jack": "589",
		"Sam":  "567",
		"Cat":  "567",
		"Wan":  "567",
	}
)
func main() {

	var port int
	var api bool
	flag.IntVar(&port, "port", 8001, "Geecache server port")
	flag.BoolVar(&api, "api", false, "Start a api server")
	flag.Parse()


	apiAddr := "http://localhost:9999"
	addrMap := map[int]string {
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}
	var addrs []string
	for _, addr := range addrMap {
		addrs = append(addrs, addr)
	}

	gee := createGroup()
	if api {
		go startApiServer(apiAddr, gee)
	}
	startCacheServer(addrMap[port],[]string(addrs), gee)
}

func startCacheServer(addr string, addrs []string, gee *geecache.Group) {
	peers := http2.NewHttpPool(addr)
	peers.Set(addrs...)
	gee.RegisterPeers(peers)
	log.Println("geecache running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

func startApiServer(addr string, gee *geecache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			key := request.URL.Query().Get("key")
			byteView, err := gee.Get(key)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			writer.Header().Set("Content-Type", "application/octet-stream")
			writer.Write(byteView.ByteSlice())
		}))
	log.Println("fe server is running at ", addr)
	log.Fatal(http.ListenAndServe(addr[7:],nil))
}

func createGroup() *geecache.Group {
	return geecache.NewGroup("scores", 2<<10, geecache.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))
}