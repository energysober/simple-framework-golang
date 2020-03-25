package main

import (
	"flag"
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

func createGroup() *gescache.Group {
	return gescache.NewGroup("scores", 2<<10, gescache.GetterFunc(
		func(key string) ([]byte, error) {
			if value, ok := db[key]; ok {
				return []byte(value), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

func startCacheServer(addr string, addrs []string, ges *gescache.Group) {
	peers := gescache.NewHTTPPool(addr)
	peers.Set(addrs...)
	ges.RegisterPeers(peers)
	log.Println("cache is running at: ", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

func startAPIServer(apiAddr string, ges *gescache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := ges.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())
		}))
	log.Println("fontend server is running at: ", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))
}

func main() {
	var port int
	var api bool
	flag.IntVar(&port, "port", 8001, "ges cache server port")
	flag.BoolVar(&api, "api", false, "Start api server?")
	flag.Parse()

	apiAddr := "http://localhost:9999"
	addrMap := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}
	ges := createGroup()
	if api {
		go startAPIServer(apiAddr, ges)
	}
	startCacheServer(addrMap[port], addrs, ges)
}
