package TinyGin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpContext struct {
	W http.ResponseWriter
	Req *http.Request

	method string
	path string
	statusCode int
}

func GenerateHttpContext(w http.ResponseWriter, req *http.Request) *HttpContext {
	context:=&HttpContext{
		W:          w,
		Req:        req,
		method:    req.Method,
		path: req.URL.Path,
	}
	fmt.Println(context.path)
	return context
}

func (c *HttpContext) setHeader(key,value string)  {
	c.W.Header().Set(key,value)
}
func (c *HttpContext) sendHeader(code int)  {
	c.statusCode=code
	c.W.WriteHeader(code)
}
func (c *HttpContext) PostForm(key string) string{
	return c.Req.PostForm.Get(key)
}
func (c *HttpContext) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}
func (c *HttpContext) SendData(code int,data []byte)  {
	c.sendHeader(code)
	c.W.Write(data)
}
func (c *HttpContext) SendString(code int,str string)  {
	c.setHeader("Content-Type", "text/plain")
	c.sendHeader(code)
	c.W.Write([]byte(str))
}
type H map[string]interface{}
func (c *HttpContext) SendJson(code int,data H) {
	c.setHeader("Content-Type", "application/json")
	//json数据序列化
	encoder := json.NewEncoder(c.W)
	if err := encoder.Encode(data);err!=nil{
		c.sendHeader(500)
		c.W.Write([]byte("json Decoding error"))
	}
}
func (c *HttpContext) SendHtml(code int,Html string)  {
	c.setHeader("Content-Type", "text/html")
	c.sendHeader(code)
	c.W.Write([]byte(Html))
}

