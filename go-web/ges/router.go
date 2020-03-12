package ges

import "fmt"

type router struct {
	handler map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handler: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method string, path string, handler HandlerFunc) {
	key := method + "_" + path
	r.handler[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "_" + c.Path
	if handler, ok := r.handler[key]; ok {
		handler(c)
	} else {
		fmt.Fprintf(c.Writer, "400, not found method: %s, path: %s", c.Method, c.Path)
	}
}
