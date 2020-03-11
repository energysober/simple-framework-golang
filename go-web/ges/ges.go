package ges

import (
	"fmt"
	"net/http"
)

// HandlerFunc method handler
type HandlerFunc func(w http.ResponseWriter, req *http.Request)

// Engine http engine
type Engine struct {
	router map[string]HandlerFunc
}

// New Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (e *Engine) addRoute(method string, path string, handler HandlerFunc) {
	key := method + "_" + path
	e.router[key] = handler
}

// GET get method
func (e *Engine) GET(path string, handler HandlerFunc) {
	e.addRoute("GET", path, handler)
}

// POST post method
func (e *Engine) POST(path string, handler HandlerFunc) {
	e.addRoute("POST", path, handler)
}

// PUT put method
func (e *Engine) PUT(path string, handler HandlerFunc) {
	e.addRoute("PUT", path, handler)
}

// DELETE delete method
func (e *Engine) DELETE(path string, handler HandlerFunc) {
	e.addRoute("DELETE", path, handler)
}

func (e *Engine) Run(addr string) (err error) {
	err = http.ListenAndServe(addr, e)
	return
}

// ServeHTTP
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "_" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "400, can not find method: %s, path: %s ", req.Method, req.URL.Path)
	}
}
