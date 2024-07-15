package api

import (
	"fmt"
	"net/http"
)

type Server struct {
	r      *http.ServeMux
	srv    *http.Server
	h      *Handler
	authMw *Middleware
}

func NewServer(port int, h *Handler, authMw *Middleware) *Server {
	router := http.NewServeMux()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	return &Server{
		r:      router,
		srv:    srv,
		h:      h,
		authMw: authMw,
	}
}

func (s *Server) Start() error {
	s.r.HandleFunc("POST /signup", s.h.AddSeller)
	s.r.HandleFunc("POST /login", s.h.Login)

	return s.srv.ListenAndServe()
}
