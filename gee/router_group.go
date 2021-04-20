package main

import (
	"log"
	"net/http"
	"path"
)

// 路由分组
type RouterGroup struct {
	prefix string
	engine *Engine
	middlewares []HandleFunc // 中间件支持
}

// 创建分组
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: group.engine,
		middlewares: make([]HandleFunc,0),
	}
	group.engine.groups = append(group.engine.groups, newGroup)
	return newGroup
}

// 将中间件应用到分组上
func (group *RouterGroup) Use(middlewares ...HandleFunc)  {
	group.middlewares = append(group.middlewares, middlewares...)
}



// @title 给分组添加路由
// @param method 路由的请求方法
// @param comp 路由组下的子路由
// @param handleFunc 路由对应的处理方法
func (group *RouterGroup) addRouter(method string, comp string, handleFunc HandleFunc) {
	path := group.prefix + comp
	log.Printf("route %4s - %s", method, path)
	group.engine.router.addRoute(method, path, handleFunc)
}

func (group *RouterGroup) Get(pattern string, handleFunc HandleFunc) {
	group.addRouter("GET", pattern, handleFunc)
}

func (group *RouterGroup) Post(pattern string, handleFunc HandleFunc) {
	group.addRouter("POST", pattern, handleFunc)
}

func (group *RouterGroup) Static(relativePath string, rootPath string) {
	handler := group.createStaticFileHandler(relativePath, http.Dir(rootPath))
	urlPattern := path.Join(relativePath, "/*filePath")
	group.Get(urlPattern, handler)
}

func (group *RouterGroup)  createStaticFileHandler(relativePath string, fd http.Dir) HandleFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fd))
	return func(c *Context) {
		file := c.Param("filePath")
		if _, err := fd.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}