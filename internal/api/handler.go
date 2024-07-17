package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"ozon/internal/entity"
	"ozon/internal/service"
)

type Service interface {
	AddSeller(ctx context.Context, sl entity.Seller) (entity.Seller, error)
	SellerBySessionID(ctx context.Context, ssid string) (entity.Seller, error)
	Login(ctx context.Context, sellerLogin string) (entity.Session, error)
}

type Handler struct {
	s Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		s: s,
	}
}

func (h *Handler) AddSeller(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var seller entity.Seller

	err := json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		SendErr(ctx, w, http.StatusBadRequest, err)
		return
	}

	seller.Normalize()

	err = seller.Validate()
	if err != nil {
		SendErr(ctx, w, http.StatusBadRequest, err)
		return
	}

	seller, err = h.s.AddSeller(r.Context(), seller)
	if err != nil {
		if errors.Is(err, service.ErrAlreadyExists) {
			SendErr(ctx, w, http.StatusConflict, err)
			return
		}

		SendErr(ctx, w, http.StatusInternalServerError, err)
		return
	}

	SendJSON(ctx, w, seller)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var seller entity.Seller

	err := json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		SendErr(ctx, w, http.StatusBadRequest, err)
		return
	}

	sess, err := h.s.Login(r.Context(), seller.Login)
	if err != nil {
		if errors.Is(err, service.ErrUnauthorized) {
			SendErr(ctx, w, http.StatusUnauthorized, err)
			return
		}

		SendErr(ctx, w, http.StatusInternalServerError, err)
		return
	}

	cookie := &http.Cookie{
		Name:    "ssid",
		Value:   sess.ID.String(),
		Expires: sess.ExpiredAt,
		MaxAge:  int(time.Until(sess.ExpiredAt).Seconds()),
	}

	http.SetCookie(w, cookie)
}

func (h *Handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ok!")
}
