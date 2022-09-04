package gee

import (
	"log"
	"net/http"
)

type route struct {
	handlers map[string]HandlerFunc
}

func newRoute() *route {
	return &route{handlers: make(map[string]HandlerFunc)}
}

func (r *route) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handlerFunc
}

func (r *route) handler(ctx *Context) {
	key := ctx.Method + "-" + ctx.Path

	if handlerFunc, ok := r.handlers[key]; ok {
		handlerFunc(ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 NOT FOUND:%s\n", ctx.Path)
	}

}
