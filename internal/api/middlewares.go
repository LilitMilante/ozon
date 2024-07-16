package api

import (
	"log/slog"
	"net/http"

	"ozon/internal/service"

	"github.com/google/uuid"
)

type Middleware struct {
	s *service.Service
	l *slog.Logger
}

func NewMiddleware(s *service.Service, l *slog.Logger) *Middleware {
	return &Middleware{
		s: s,
		l: l,
	}
}

func (m *Middleware) WithAuth(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("ssid")
		if err != nil {
			SendErr(w, http.StatusUnauthorized, service.ErrUnauthorized)
			return
		}

		// todo: прокидывать seller через ctx
		_, err = m.s.SellerBySessionID(r.Context(), cookie.Value)
		if err != nil {
			SendErr(w, http.StatusUnauthorized, service.ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.l.Info("incoming request", "request_id", uuid.New(), "method", r.Method, "url", r.URL.String(), "from", r.RemoteAddr)

		next.ServeHTTP(w, r)
	})
}
