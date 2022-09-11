package gee

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

//roots key eg,roots['GET'] roots['POST']
//handlers key eg,handlers['GET-/p/:lang/doc'] handlers['POST-/p/book']

func newRoute() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

//Only one '*' is allowed
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, part := range vs {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}

	return parts
}

func (r *router) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)

	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	parts := parsePattern(pattern)
	r.roots[method].insert(pattern, parts, 0)

	key := method + "-" + pattern
	r.handlers[key] = handlerFunc
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)

	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *router) handler(ctx *Context) {
	route, params := r.getRoute(ctx.Method, ctx.Path)
	if route != nil {
		ctx.Params = params
		key := ctx.Method + "-" + route.pattern
		r.handlers[key](ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
	}
}
