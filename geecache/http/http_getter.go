package http

import (
	"catwang.com/go-in-action/geecache/protobuf"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type HttpGetter struct {
	baseURL string //baseURL 表示将要访问的远程节点的地址，例如 http://example.com/_geecache/
}

func (hg HttpGetter) Get(in *protobuf.Request, out *protobuf.Response)  error {
	url := fmt.Sprintf("%v%v/%v",
			hg.baseURL,
			url.QueryEscape(in.Group),
			url.QueryEscape(in.Key),
		)
	resp , err := http.Get(url)
	if err != nil {
		log.Panicf("http get :%s error:[%+v]", url, err)
		return  err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server resp:%v", resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading resp body:%v", err)
	}
	if err := proto.Unmarshal(bytes,out) ; err != nil{
		return fmt.Errorf("proto unmatshal err:%v", err)
	}
	return  nil
}
