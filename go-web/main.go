package main

import (
	"fmt"
	"github.com/simple-framework-golang/go-web/ges"
	"html/template"
	"net/http"
	"time"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := ges.Default()
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	//r.LoadHTMLGlob("www/templates/*")
	//r.Static("/assets", "./www/static")

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *ges.Context) {
			c.String(http.StatusOK, "hi, ges!!")
		})
		v1.GET("/hello/:name", func(c *ges.Context) {
			c.String(http.StatusOK, "hi, I am %s", c.Param("name"))
		})
		v1.GET("/file/*filePath", func(c *ges.Context) {
			c.JSON(http.StatusOK, ges.H{
				"user":     "ges",
				"action":   "say hi",
				"filePath": c.Param("filePath"),
			})
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *ges.Context) {
			c.String(http.StatusOK, "hi, I am %s", c.Param("name"))
		})
		//v2.GET("/html", func(c *ges.Context) {
		//	c.HTML(http.StatusOK, "/css.tmpl", nil)
		//})
	}
	r.Run(":8899")
}
