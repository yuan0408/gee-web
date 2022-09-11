package gee

import (
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(ctx *Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	route *router
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{route: newRoute()}
}

func (e *Engine) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	e.route.addRoute(method, pattern, handlerFunc)
}

// GET defines the method to add GET request
func (e *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	e.addRoute("GET", pattern, handlerFunc)
}

// POST defines the method to add POST request
func (e *Engine) POST(pattern string, handlerFunc HandlerFunc) {
	e.addRoute("POST", pattern, handlerFunc)
}

// Run defines the method to start a http server
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := newContext(req, w)
	e.route.handler(ctx)
}
