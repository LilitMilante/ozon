package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"sellers-ms/internal/entity"
	"sellers-ms/internal/service"
)

type SellersService interface {
	AddSeller(ctx context.Context, sl entity.Seller) (entity.Seller, error)
	SellerBySessionID(ctx context.Context, ssid string) (entity.Seller, error)
	Login(ctx context.Context, sellerLogin string) (entity.Session, error)
}

type Handler struct {
	sellers  SellersService
	products *service.Products
}

func NewHandler(s SellersService, products *service.Products) *Handler {
	return &Handler{
		sellers:  s,
		products: products,
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

	seller, err = h.sellers.AddSeller(r.Context(), seller)
	if err != nil {
		if errors.Is(err, entity.ErrAlreadyExists) {
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

	sess, err := h.sellers.Login(r.Context(), seller.Login)
	if err != nil {
		if errors.Is(err, entity.ErrUnauthorized) {
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

func (h *Handler) ProductByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(r.PathValue("product_id"))
	if err != nil {
		SendErr(ctx, w, http.StatusBadRequest, fmt.Errorf("id must be a number"))
		return
	}

	p, err := h.products.ProductByID(ctx, int64(id))
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			SendErr(ctx, w, http.StatusNotFound, err)
			return
		}
		SendErr(ctx, w, http.StatusInternalServerError, err)
		return
	}

	SendJSON(ctx, w, p)
}
