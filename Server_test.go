package TinyGin

import (
	"fmt"
	"testing"
)

func TestNewServer(t *testing.T) {
	s:=NewServer()
	s.AddGet("/Hello", func(ctx *HttpContext) {
		ctx.SendString(200,fmt.Sprintf("Hello%s",ctx.Req.URL))
	})
	s.Run(":9999")
}

