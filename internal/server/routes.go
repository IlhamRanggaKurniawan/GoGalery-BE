package server

import (
	"fmt"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler{
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "path /")
	})

	return mux
}