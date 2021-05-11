package http

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type HttpGetter struct {
	baseURL string //baseURL 表示将要访问的远程节点的地址，例如 http://example.com/_geecache/
}

func (hg HttpGetter) Get(group, key string) ([] byte, error) {
	url := fmt.Sprintf("%v%v/%v",
			hg.baseURL,
			url.QueryEscape(group),
			url.QueryEscape(key),
		)
	resp , err := http.Get(url)
	if err != nil {
		log.Panicf("http get :%s error:[%+v]", url, err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server resp:%v", resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading resp body:%v", err)
	}
	return bytes, nil
}
