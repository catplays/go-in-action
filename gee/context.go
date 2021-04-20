package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Writer http.ResponseWriter
	Req *http.Request
	Path string
	Method string
	Params map[string] string // 存储路由参数到param中
	StatusCode int
	handlers []HandleFunc // 中间件
	index int // 执行到了第几个中间件
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req: r,
		Path: r.URL.Path,
		Method: r.Method,
		Params: make(map[string]string),
		index: -1,
	}
}

func (context *Context) Next()  {
	context.index++
	n := len(context.handlers)
	for ; context.index< n; context.index++ {
		context.handlers[context.index](context)
	}
}

func (context *Context) Param(key string) string  {
	value,_ :=  context.Params[key]
	return value
}

// get post form value
func (context *Context) PostForm(key string) string  {
	return context.Req.FormValue(key)
}

func (context *Context) Query(key string) string {
	return context.Req.URL.Query().Get(key)
}

func (context *Context) Status(code int )  {
	context.StatusCode = code
	context.Writer.WriteHeader(code)
}

func (context *Context) setHeader(key,value string)  {
	context.Writer.Header().Set(key,value)
}

func (context *Context) String (code int, format string, values ...interface{})  {
	context.setHeader("Content-Type","text/plain")
	context.Status(code)
	context.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (context *Context) JSON (code int ,obj interface{})  {
	context.setHeader("Content-Type", "application/json")
	context.Status(code)
	encoder := json.NewEncoder(context.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(context.Writer, err.Error(), 500)
	}
}


func (context *Context) Data(code int, data []byte)  {
	context.Status(code)
	context.Writer.Write(data)
}

func (context *Context) HTML(code int, html string) {
	context.setHeader("Content-Type", "text/html")
	context.Status(code)
	context.Writer.Write([]byte(html))
}

func (context *Context) Fail(code int, err string) {
	context.index = len(context.handlers)
	context.JSON(code, H{"message": err})
}
