package TinyGin

import (
	"fmt"
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

func TestGroup(t *testing.T) {
	s:=NewServer()
	v1 := s.Group("/v1")
	{
		v1.AddGet("/hello", func(ctx *HttpContext) {
				ctx.SendString(200,fmt.Sprintf("you are at the %sGroup",v1.prefix))
		})
	}
	v2:=s.Group("v2")
	{
		v2.AddGet("/hello", func(ctx *HttpContext) {
			ctx.SendString(200,fmt.Sprintf("you are at the %sGroup",v2.prefix))
		})
	}
	s.Run(":9999")
}