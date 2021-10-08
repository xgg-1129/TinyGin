package TinyGin

import (
	"log"
	"net/http"

	"testing"
)

func TestNewServer(t *testing.T) {
	s:=NewServer()
	log.Fatal(http.ListenAndServe(":9999",s))
}

