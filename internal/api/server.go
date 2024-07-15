package api

import (
	"fmt"
	"net/http"
)

type Server struct {
	r      *http.ServeMux
	srv    *http.Server
	h      *Handler
	authMw *AuthMiddleware
}

func NewServer(port int, h *Handler, authMw *AuthMiddleware) *Server {
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
	s.r.Handle("POST /signup", s.authMw.Require(s.h.AddSeller))

	s.r.HandleFunc("/login", s.h.Login)

	return s.srv.ListenAndServe()
}
