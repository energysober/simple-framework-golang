package main

import (
	"github.com/simple-framework-golang/go-web/ges"
	"net/http"
)

func main() {
	r := ges.New()
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *ges.Context) {
			c.String(http.StatusOK, "hi, ges!!")
		})
		v1.GET("/hello/:name", func(c *ges.Context) {
			c.String(http.StatusOK, "hi, I am %s", c.Param["name"])
		})
		v1.GET("/file/*filePath", func(c *ges.Context) {
			c.JSON(http.StatusOK, ges.H{
				"user":     "ges",
				"action":   "say hi",
				"filePath": c.Param["filePath"],
			})
		})
		v1.GET("/html", func(c *ges.Context) {
			c.HTML(http.StatusOK, "<h1>hi! ges, this is a html</h1>")
		})
	}


	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *ges.Context) {
			c.String(http.StatusOK, "hi, I am %s", c.Param["name"])
		})
	}
	r.Run(":8899")
}
