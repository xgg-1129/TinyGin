package TinyGin

import (
	"testing"
)

func newTest() *Route{
	r:=NewRoute()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}
func TestNewRoute(t *testing.T) {
	r:=newTest()

	n := r.GetRoute("GET", "/hello/xgg")

	if n == nil{
		t.Fatal("nil shouldn't be returned")
	}
	if n.path != "/hello/:name"{
		t.Fatal("match is fault")
	}
}

