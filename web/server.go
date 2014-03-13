package web

import (
	"fmt"
	"net/http"
)

type Server struct {
	config *Config
}

func Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/javascript; charset=utf-8")
	fmt.Fprintf(w, `{"test":"good"}`)
}

func NewServer() *Server {
	var server *Server
	server = new(Server)
	server.config = NewConfig()
	return server
}

func (s *Server) Start() {
	http.HandleFunc("/", Root)
	fmt.Println("Listening on port:", s.config.port)

	http.ListenAndServe(fmt.Sprintf(":%v", s.config.port), nil)
}
