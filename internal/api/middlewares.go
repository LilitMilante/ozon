package api

import (
	"net/http"

	"ozon/internal/service"
)

type Middleware struct {
	s *service.Service
}

func NewMiddleware(s *service.Service) *Middleware {
	return &Middleware{
		s: s,
	}
}

func (a *Middleware) WithAuth(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("ssid")
		if err != nil {
			SendErr(w, http.StatusUnauthorized, service.ErrUnauthorized)
			return
		}

		// todo: прокидывать seller через ctx
		_, err = a.s.SellerBySessionID(r.Context(), cookie.Value)
		if err != nil {
			SendErr(w, http.StatusUnauthorized, service.ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
