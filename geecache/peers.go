package geecache

import (
	"catwang.com/go-in-action/geecache/protobuf"
)

//根据传入的 key 选择相应节点的获取key的方式
type PeerPicker interface {
	PickPeer(key string) (PeerGetter, bool)
}

//定义获取缓存值的接口  比如 HTTP实现的Get，TODO 加上RPC实现的Get
type PeerGetter interface {
	//Get() 方法用于从对应 group 查找缓存值
	Get(in *protobuf.Request, out *protobuf.Response) error
}