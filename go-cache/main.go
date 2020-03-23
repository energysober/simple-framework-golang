package main

import (
	"fmt"
	"github.com/simple-framework-golang/go-cache/gescache"
	"log"
	"net/http"
)

var db = map[string]string{
	"Windy": "100",
	"Mack":  "99",
	"Stone": "80",
}

func main() {
	gescache.NewGroup("scores", 2<<10, gescache.GetterFunc(
		func(key string) ([]byte, error) {
			if value, ok := db[key]; ok {
				return []byte(value), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := gescache.NewHTTPPool(addr)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
