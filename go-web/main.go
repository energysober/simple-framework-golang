package main

import (
	"fmt"
	"github.com/simple-framework-golang/go-web/ges"
	"net/http"
)

func main() {
	r := ges.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hi, ges!!")
	})
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hi,ges, url path is %s", req.URL.Path)
	})
	r.Run(":8899")
}
