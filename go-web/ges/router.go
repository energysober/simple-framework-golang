package ges

import (
	"fmt"
	"strings"
)

type router struct {
	roots   map[string]*node
	handler map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:   make(map[string]*node),
		handler: make(map[string]HandlerFunc),
	}
}

func (r *router) parsePattern(path string) []string {
	vs := strings.Split(path, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item == "" {
			continue
		}
		parts = append(parts, item)
		if item[0] == '*' {
			break
		}
	}
	return parts
}

func (r *router) addRoute(method string, path string, handler HandlerFunc) {
	parts := r.parsePattern(path)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(path, parts, 0)

	key := method + "_" + path
	r.handler[key] = handler
}

func (r *router) getRoute(method string, path string) *node {
	parts := r.parsePattern(path)
	n := r.roots[method].search(parts, 0)
	if n == nil || n.pattern == "" {
		return nil
	}
	return n
}

func (r *router) handle(c *Context) {
	n := r.getRoute(c.Method, c.Path)
	key := c.Method + "_" + n.pattern
	if handler, ok := r.handler[key]; ok {
		handler(c)
	} else {
		fmt.Fprintf(c.Writer, "400, not found method: %s, path: %s", c.Method, c.Path)
	}
}
