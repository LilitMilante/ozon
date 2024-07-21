package api

import (
	"fmt"
	"net/http"
)

type Server struct {
	r   *http.ServeMux
	srv *http.Server
	h   *Handler
	mw  *Middleware
}

func NewServer(port int, h *Handler, mw *Middleware) *Server {
	router := http.NewServeMux()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mw.Log(router),
	}

	return &Server{
		r:   router,
		srv: srv,
		h:   h,
		mw:  mw,
	}
}

func (s *Server) Start() error {
	s.r.Handle("POST /products", s.mw.WithAuth(s.h.AddProduct))
	s.r.HandleFunc("GET /products/{product_id}", s.h.ProductByID)

	s.r.HandleFunc("POST /signup", s.h.AddSeller)
	s.r.HandleFunc("POST /login", s.h.Login)

	return s.srv.ListenAndServe()
}
