package main

import (
	"net/http"
)

type Engine struct {
	router *Router
	*RouterGroup // 这里类似于继承，可以直接迟永RouterGroup中的方法
	groups []*RouterGroup // 路由分组

}
// 定义处理路由请求的接口
type HandleFunc func(c *Context)

type H map[string] interface{}

// Constructor
func New() *Engine {
	engine :=  &Engine{
		router: NewRouter(),
	}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = [] *RouterGroup{engine.RouterGroup}
	return engine
}

func (engine *Engine) Run(addr string) (err error) {
	err = http.ListenAndServe(addr, engine)
	return err
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	context := newContext(w, req)
	engine.router.handle(context)
}


func (engine *Engine) Post(path string, handleFunc HandleFunc)  {
	engine.router.addRoute("POST", path, handleFunc)
}
func (engine *Engine) Get(path string, handleFunc HandleFunc)  {
	engine.router.addRoute("GET", path, handleFunc)
}
