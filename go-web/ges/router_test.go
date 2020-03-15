package ges

import (
	"fmt"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/test", nil)
	r.addRoute("GET", "/test/:testParam", nil)
	r.addRoute("GET", "/test/*testPath", nil)
	r.addRoute("POST", "/", nil)
	r.addRoute("POST", "/test", nil)
	r.addRoute("POST", "/test/:testParam", nil)
	r.addRoute("POST", "test/*testPath", nil)
	return r
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n := r.getRoute("GET", "/test/testParam")
	if n == nil {
		t.Fatal("nil should not be returned")
	}

	if n.pattern != "/test/:testParam" {
		t.Fatal("should match /test/:testParam")
	}

	fmt.Printf("match path: %s", n.pattern)
}
