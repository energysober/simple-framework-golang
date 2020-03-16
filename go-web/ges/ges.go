package ges

import (
	"net/http"
)

// HandlerFunc method handler
type HandlerFunc func(c *Context)

type RouteGroup struct {
	prefix      string
	middleWares []HandlerFunc
	parent      *RouteGroup
	engine      *Engine
}

// Engine http engine
type Engine struct {
	*RouteGroup
	router *router
	groups []*RouteGroup
}

// New Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouteGroup = &RouteGroup{engine: engine}
	engine.groups = []*RouteGroup{engine.RouteGroup}
	return engine
}

// Group
func (group *RouteGroup) Group(prefix string) *RouteGroup {
	engine := group.engine
	newGroup := &RouteGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouteGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, handler)
}

// GET group get
func (group *RouteGroup) GET(path string, handler HandlerFunc) {
	group.addRoute("GET", path, handler)
}

// POST group post
func (group *RouteGroup) POST(path string, handler HandlerFunc) {
	group.addRoute("POST", path, handler)
}

// PUT group put
func (group *RouteGroup) PUT(path string, handler HandlerFunc) {
	group.addRoute("PUT", path, handler)
}

// DELETE group delete
func (group *RouteGroup) DELETE(path string, handler HandlerFunc) {
	group.addRoute("DELETE", path, handler)
}

// GET get method
func (e *Engine) GET(path string, handler HandlerFunc) {
	e.router.addRoute("GET", path, handler)
}

// POST post method
func (e *Engine) POST(path string, handler HandlerFunc) {
	e.router.addRoute("POST", path, handler)
}

// PUT put method
func (e *Engine) PUT(path string, handler HandlerFunc) {
	e.router.addRoute("PUT", path, handler)
}

// DELETE delete method
func (e *Engine) DELETE(path string, handler HandlerFunc) {
	e.router.addRoute("DELETE", path, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

// ServeHTTP
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
