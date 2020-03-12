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
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Method: req.Method,
		Path:   req.URL.Path,
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

// HTML format HTML
func (c *Context) HTML(code int, html string) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/html")
	c.Writer.Write([]byte(html))
}

// PostForm get form value by key
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query get url param by key
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}
