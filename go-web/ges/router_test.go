package ges

import (
	"fmt"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/test_a/:testParam", nil)
	r.addRoute("GET", "/test_b/*testPath", nil)
	r.addRoute("POST", "/", nil)
	r.addRoute("POST", "/test_a/:testParam", nil)
	r.addRoute("POST", "/test_b/*testPath", nil)
	return r
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, param := r.getRoute("GET", "/test_a/paramTest")
	if n == nil {
		t.Fatal("nil should not be returned")
	}

	if n.pattern != "/test_a/:testParam" {
		t.Fatal("should match /test_a/:testParam")
	}

	if param["testParam"] != "paramTest" {
		t.Fatal("testParam should be equal paramTest")
	}

	n, param = r.getRoute("GET", "/test_b/testPath/test.txt")
	if n == nil {
		t.Fatal("nil should not be returned")
	}

	if n.pattern != "/test_b/*testPath" {
		t.Fatal("should match /test_b/:testPath")
	}

	if param["testPath"] != "testPath/test.txt" {
		t.Fatal("testPath should be equal testPath/test.txt")
	}
	fmt.Printf("match path: %s", n.pattern)
}
