package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"ozon/internal/entity"
	"ozon/internal/service"

	"github.com/google/uuid"
)

type Service interface {
	AddSeller(ctx context.Context, sl entity.Seller) (entity.Seller, error)
	SellerBySessionID(ctx context.Context, ssid string) (entity.Seller, error)
	SellerByLogin(ctx context.Context, login string) (entity.Seller, error)
	Login(ctx context.Context, sellerID uuid.UUID) (entity.Session, error)
}

type Handler struct {
	s Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		s: s,
	}
}

// Sessions

func (h *Handler) AddSeller(w http.ResponseWriter, r *http.Request) {
	var seller entity.Seller

	err := json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		SendErr(w, http.StatusBadRequest, err)
		return
	}

	seller, err = h.s.AddSeller(r.Context(), seller)
	if err != nil {
		if errors.Is(err, service.ErrAlreadyExists) {
			SendErr(w, http.StatusConflict, err)
			return
		}

		SendErr(w, http.StatusInternalServerError, err)
		return
	}

	seller.Sanitize()

	SendJSON(w, seller)
}

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
