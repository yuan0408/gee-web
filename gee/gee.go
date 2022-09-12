package gee

import (
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(ctx *Context)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc //support middleware
	parent      *RouterGroup  //support nesting
	engine      *Engine       //all group share a engine instance
}

// Engine implement the interface of ServeHTTP
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup //store all group
}

// New is the constructor of gee.Engine
func New() *Engine {
	engine := &Engine{router: newRoute()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		engine: engine,
		parent: group,
		prefix: group.prefix + prefix,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handlerFunc HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, handlerFunc)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handlerFunc HandlerFunc) {
	group.addRoute("GET", pattern, handlerFunc)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handlerFunc HandlerFunc) {
	group.addRoute("POST", pattern, handlerFunc)
}

// Run defines the method to start a http server
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := newContext(req, w)
	e.router.handler(ctx)
}
