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

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	urlParts := r.parsePattern(path)
	param := make(map[string]string, 0)
	n := r.roots[method].search(urlParts, 0)
	if n != nil {
		parts := r.parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				param[part[1:]] = urlParts[index]
			} else if part[0] == '*' && len(part) > 1 {
				param[part[1:]] = strings.Join(urlParts[index:], "/")
				break
			}
		}
		return n, param
	}

	return nil, nil
}

func (r *router) handle(c *Context) {
	n, param := r.getRoute(c.Method, c.Path)
	c.Param = param
	key := c.Method + "_" + n.pattern
	if handler, ok := r.handler[key]; ok {
		handler(c)
	} else {
		fmt.Fprintf(c.Writer, "400, not found method: %s, path: %s", c.Method, c.Path)
	}
}
