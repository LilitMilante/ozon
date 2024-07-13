package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"ozon/internal/entity"
	"ozon/internal/service"
)

type Service interface {
	SellerBySessionID(ctx context.Context, ssid string) (entity.Seller, error)
	SellerByLogin(ctx context.Context, login string) (entity.Seller, error)
	Login(ctx context.Context, sellerID int64) (entity.Session, error)
}

type Handler struct {
	s *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		s: s,
	}
}

// Sessions

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var seller entity.Seller

	err := json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		SendErr(w, http.StatusBadRequest, err)
		return
	}

	seller, err = h.s.SellerByLogin(r.Context(), seller.Login)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) || !seller.ComparePassword(seller.Password) {
			SendErr(w, http.StatusUnauthorized, fmt.Errorf("%w: incorrect login or password", service.ErrUnauthorized))
			return
		}

		SendErr(w, http.StatusInternalServerError, err)
		return
	}

	sess, err := h.s.Login(r.Context(), seller.ID)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err)
		return
	}

	cookie := &http.Cookie{
		Name:    "ssid",
		Value:   sess.ID.String(),
		Expires: sess.ExpiredAt,
	}

	http.SetCookie(w, cookie)
}
