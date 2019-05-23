package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
)

type Server struct {
	RequestCount int64
	File         *os.File
	Port         string
}

func (s *Server) Listen() {
	http.ListenAndServe(":"+s.Port, nil)
}

func NewServer(port string, filePath string) *Server {
	server := &Server{RequestCount: 0}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	return server
}
