package main

import (
	"fmt"
	"net/http"
)

func main() {
	engine := New()
	// 指定静态文件目录
	engine.Static("/assets", "./gee/static")
	// 注册中间件
	engine.Use(Logger(), Recover())
	engine.Get("/", IndexHandler)
	v1Group := engine.Group("/v1")
	v1Group.Get("/", func(c *Context) {
		c.HTML(200, "<h1> hello v1 </h1>")
	})

	v1Group.Get("/hello", func(c *Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	v2Group := engine.Group("/v2")
	{
		v2Group.Get("/hello/:name", func(c *Context) {
			c.String(http.StatusOK, "hello , your name is %s url: %s", c.Param("name"), c.Path)
		})
		v2Group.Post("/login", func(c *Context) {
			c.JSON(http.StatusOK, H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	engine.Get("/panic", func(c *Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	if err := engine.Run(":9999"); err != nil {
		panic("")
	}
}

// middleware echoes r.URL.Path
func IndexHandler(c *Context) {
	fmt.Fprintf(c.Writer, "URL.Path = %q\n", c.Path)
}

// middleware echoes r.URL.Header
func HelloHandler(c *Context) {
	for k, v := range c.Req.Header {
		fmt.Fprintf(c.Writer, "Header[%q] = %q\n", k, v)
	}
}


