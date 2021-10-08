package TinyGin

import (
	"log"
	"net/http"
)

type Server struct {
	route *Route
}

func NewServer()*Server{
	s:=&Server{}
	s.route=NewRoute()
	return s
}
func (s *Server) Run(addr string)  {
	log.Fatal(http.ListenAndServe(addr,s))
}
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	context := GenerateHttpContext(w, req)
	s.route.handle(context)
}
func (s *Server) AddGet(path string,fun HandleFun) error{
	return s.route.AddGet(path,fun)
}
func (s *Server) AddPost(path string,fun HandleFun)error {
	return s.route.AddPost(path,fun)
}

