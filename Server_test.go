package TinyGin

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

func Logger() HandleFun {
	return func(c *HttpContext) {
		// Start timer
		t := time.Now()
		// Process request
		c.doAllNext()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.statusCode, c.Req.RequestURI, time.Since(t))
	}
}
func onlyForV2()HandleFun {
	return func(c *HttpContext) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.SendString(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.statusCode, c.Req.RequestURI, time.Since(t))
	}
}

func Test(t *testing.T){
	r := NewServer()

	v2 := r.Group("/v2")
	v2.RegisterMiddles(onlyForV2()) // v2 group middleware
	{
		v2.AddGet("/hello/:name", func(c *HttpContext) {
			// expect /hello/geektutu
			c.SendString(http.StatusOK, fmt.Sprintf("hello , you're at %s\n",  c.path))
		})
	}

	r.Run(":9999")
	return
}
func TestNewRoute(t *testing.T) {

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