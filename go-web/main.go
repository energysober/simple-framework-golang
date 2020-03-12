package main

import (
	"github.com/simple-framework-golang/go-web/ges"
	"net/http"
)

func main() {
	r := ges.New()
	r.GET("/", func(c *ges.Context) {
		c.String(http.StatusOK, "hi, ges!!")
	})
	r.GET("/hello", func(c *ges.Context) {
		c.String(http.StatusOK, "hi,ges, url path is %s", c.Path)
	})
	r.GET("/json", func(c *ges.Context) {
		c.JSON(http.StatusOK, ges.H{
			"user":   "ges",
			"action": "say hi",
		})
	})
	r.GET("/html", func(c *ges.Context) {
		c.HTML(http.StatusOK, "<h1>hi! ges, this is a html</h1>")
	})
	r.Run(":8899")
}
