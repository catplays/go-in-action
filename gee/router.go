package main

import (
	"fmt"
	"strings"
)

type Router struct {
	root map[string]*Node
	handlers map[string] HandleFunc
}

// Constructor
func NewRouter() *Router {
	return &Router{
		root: make(map[string]*Node),
		handlers: make(map[string]HandleFunc),
	}
}

func (router *Router) addRoute(method string, pattern string, handleFunc HandleFunc) {
	key := fmt.Sprintf("%s-%s", method,pattern)
	_, ok := router.root[method]
	if !ok {
		router.root[method] = &Node{}
	}
	parts := parsePattern(pattern)
	router.root[method].insert(pattern,parts,0)
	_, exist := router.handlers[key]
	if exist {
		panic(fmt.Sprintf("URL:%s with method:%s has registered.",pattern,method))
	}
	router.handlers[key] = handleFunc
}

// parse pattern into slice, and only one * is allowed
func parsePattern(pattern string) []string {
	strs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range strs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				 break
			}
		}
	}
	return parts

}

func (router *Router) handle(context *Context) {
	// 处理路由
	
	node, params := router.GetRoute(context.Method, context.Path)
	if node != nil {
		context.Params = params
		urlKey :=  fmt.Sprintf("%s-%s", context.Method, node.pattern)
		if handler, ok := router.handlers[urlKey]; ok {
			handler(context)
			return
		}
	}


	fmt.Fprintf(context.Writer, "404 not found.")

}

func (router *Router) GetRoute(method string, path string) (*Node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	
	root, exist := router.root[method]
	if !exist {
		return nil, nil
	}
	node := root.search(searchParts,0)
	if node != nil {
		var parts = parsePattern(node.pattern)
		for index, item := range parts {
			if item[0] == ':' {
				// /p/go/doc匹配到/p/:lang/doc，解析结果为：{lang: "go"}
				params[item[1:]] = searchParts[index]
			}
			if item[0] == '*' && len(item)>1 {
				///static/css/geektutu.css匹配到/static/*filepath，解析结果为{filepath: "css/geektutu.css"}
				params[item[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
	}
	return node, params
}