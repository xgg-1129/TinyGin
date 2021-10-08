package TinyGin

import (
	"fmt"
	"net/http"
)

type Server struct {

}

func NewServer()*Server{
	s:=&Server{}
	return s
}
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	switch req.URL.Path {
	case "/":
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "URL-PATH=%s", req.URL.Path)
	case "/Hello":
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,"Helo world")
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"404 no found")
	}
}
