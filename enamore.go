package enamore

import (
	"fmt"
	"net/http"
)

// HandlerFunc 定义了enamore使用的请求处理函数类型。
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 实现了ServeHTTP接口的请求引擎
type Engine struct {
	router map[string]HandlerFunc // 路由映射，存储请求方法和路径到处理函数的映射
}

// New 是Engine的构造函数。
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute 为指定的方法和路径添加处理函数。
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern // 生成唯一的键
	engine.router[key] = handler  // 将处理函数映射到键
}

// GET 用于添加GET请求的方法。
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 用于添加POST请求的方法。
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 启动一个HTTP服务器。
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP 处理HTTP请求。
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path // 生成请求的键
	if handler, ok := engine.router[key]; ok {
		handler(w, req) // 调用对应的处理函数
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL) // 未找到路由，返回404
	}
}
