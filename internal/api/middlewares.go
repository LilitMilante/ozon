package api

import (
	"context"
	"log/slog"
	"net/http"

	"sellers-ms/internal/entity"
	"sellers-ms/internal/service"

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
		ctx := r.Context()

		cookie, err := r.Cookie("ssid")
		if err != nil {
			SendErr(ctx, w, http.StatusUnauthorized, entity.ErrUnauthorized)
			return
		}

		// todo: прокидывать seller через ctx
		_, err = m.s.SellerBySessionID(r.Context(), cookie.Value)
		if err != nil {
			SendErr(ctx, w, http.StatusUnauthorized, entity.ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := m.l.With("request_id", uuid.New())

		l.Info("incoming request", "method", r.Method, "url", r.URL.String(), "from", r.RemoteAddr)

		ctx := context.WithValue(r.Context(), "logger", l)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
