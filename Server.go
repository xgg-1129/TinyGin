package TinyGin

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Server struct {
	route *Route
	Groups []*Group
	Templates map[string]*template.Template
}

func NewServer()*Server{
	s:=&Server{}
	s.route=NewRoute()
	s.Groups=make([]*Group,0)
	s.Templates=make(map[string]*template.Template)
	return s
}
func (s *Server) Run(addr string)  {
	log.Fatal(http.ListenAndServe(addr,s))
}
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	context := GenerateHttpContext(w, req)
	context.server=s
	for _,group := range s.Groups{
		if strings.HasPrefix(context.path,group.prefix){
			context.handles=append(context.handles,group.middles...)
		}
	}
	s.route.handle(context)
}
func (s *Server) AddGet(path string,fun HandleFun) error{
	return s.route.AddGet(path,fun)
}
func (s *Server) AddPost(path string,fun HandleFun)error {
	return s.route.AddPost(path,fun)
}
func (s *Server) Group(prefix string)*Group{
	g:=&Group{
		prefix:prefix,
		Server: s,
	}
	s.Groups=append(s.Groups,g)
	return g
}
//静态文件服务，标准库提供了静态文件服务，实现重定向功能即可
func (s *Server) RegisterStatic(prefixpath string,filepath string)  {
	handle:=s.createStaticHandler(prefixpath,http.Dir(filepath))
	url:=prefixpath+"/*filename"
	fmt.Println(url)
	s.AddGet(url,handle)
}

func (s *Server) createStaticHandler(relativePath string, fs http.FileSystem) HandleFun{
	fileserver:=http.StripPrefix(relativePath,http.FileServer(fs))
	return func(ctx *HttpContext) {
		fmt.Println(ctx.path)
		fileserver.ServeHTTP(ctx.W,ctx.Req)
	}
}
func (s *Server) AddHtmlTemplate(path string,name string) {
	template := template.Must(template.ParseFiles(path))
	s.Templates[name]=template
}

