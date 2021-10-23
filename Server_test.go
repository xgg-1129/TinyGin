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
type student struct {
	Name string
	Age  int8
}
func Test(t *testing.T){
	r := NewServer()
	r.AddHtmlTemplate("D:\\Environment\\ProjectGo\\src\\TinyGin\\Templates\\template1.html","tem1")
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}

	r.AddGet("/students", func(c *HttpContext) {
		c.SendHtml(http.StatusOK, "tem1", H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	r.Run(":9999")
}
func TestPanic(t *testing.T) {
	r := NewServer()
	v1 := r.Group("/v1")
	v1.RegisterMiddles(PanicRecover())
	v1.AddGet("/hello", func(ctx *HttpContext) {
		ctx.Redirect(302,"https://www.bing.com/?mkt=zh-CN")
	})
	r.Run(":9999")
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