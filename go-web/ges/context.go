package ges

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Context
type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
	Method string
	Path   string
	Params map[string]string

	handlers []HandlerFunc
	index    int

	engine *Engine
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Method: req.Method,
		Path:   req.URL.Path,
		index:  -1,
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

// SetHeader set http header
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// Status set Write code
func (c *Context) Status(code int) {
	c.Writer.WriteHeader(code)
}

// String format string
func (c *Context) String(code int, format string, values ...interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/plain")
	fmt.Fprintf(c.Writer, format, values...)
}

// JSON format json
func (c *Context) JSON(code int, obj interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

// HTML format HTML
func (c *Context) HTML(code int, name string, data interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/html")
	if err := c.engine.htmlTemplate.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(http.StatusInternalServerError, err.Error())
	}
}

// PostForm get form value by key
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query get url param by key
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}
