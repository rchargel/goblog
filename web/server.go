package web

import (
	"fmt"
	"github.com/rchargel/goblog/conf"
	"net/http"
)

type Server struct {
	config conf.Config
}

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TESTING")
}

func NewServer() *Server {
	var server *Server
	server = new(Server)
	server.config = conf.NewConfig()
	return server
}

func (s *Server) Start() {
	http.HandleFunc("/", Root)
	fmt.Println("Listening on port:", s.config.Port)

	http.ListenAndServe(fmt.Sprintf(":%v", s.config.Port), nil)
}
