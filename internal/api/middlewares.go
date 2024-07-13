package api

import (
	"net/http"

	"ozon/internal/service"
)

type AuthMiddleware struct {
	s *service.Service
}

func NewAuthMiddleware(s *service.Service) *AuthMiddleware {
	return &AuthMiddleware{
		s: s,
	}
}

func (a *AuthMiddleware) Require(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("ssid")
		if err != nil {
			SendErr(w, http.StatusUnauthorized, service.ErrUnauthorized)
			return
		}

		_, err = a.s.SellerBySessionID(r.Context(), cookie.Value)
		if err != nil {
			SendErr(w, http.StatusUnauthorized, service.ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
