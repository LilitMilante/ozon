package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"ozon/internal/entity"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	SellerByLogin(ctx context.Context, login string) (entity.Seller, error)
	SellerByID(ctx context.Context, id uuid.UUID) (entity.Seller, error)
	CreateSeller(ctx context.Context, s entity.Seller) (entity.Seller, error)

	SessionByID(ctx context.Context, id string) (entity.Session, error)
	CreateSession(ctx context.Context, sess entity.Session) error
}

type Service struct {
	repo       Repository
	sessionAge time.Duration
}

func NewService(repo Repository, ssAge time.Duration) *Service {
	return &Service{
		repo:       repo,
		sessionAge: ssAge,
	}
}

func (s *Service) AddSeller(ctx context.Context, sl entity.Seller) (entity.Seller, error) {
	_, err := s.repo.SellerByLogin(ctx, sl.Login)
	if err == nil {
		return sl, fmt.Errorf("seller with login %q: %w", sl.Login, ErrAlreadyExists)
	}

	if !errors.Is(err, ErrNotFound) {
		return sl, fmt.Errorf("seller with login %q: %w", sl.Login, err)
	}

	sl.ID = uuid.New()
	sl.Password, err = s.hashPassword(sl.Password)
	if err != nil {
		return sl, err
	}
	sl.CreatedAt = time.Now()
	sl.UpdatedAt = sl.CreatedAt

	sl, err = s.repo.CreateSeller(ctx, sl)
	if err != nil {
		return sl, fmt.Errorf("create seller: %w", err)
	}

	sl.Sanitize()

	return sl, nil
}

func (s *Service) hashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

// Session

func (s *Service) SellerBySessionID(ctx context.Context, ssid string) (entity.Seller, error) {
	session, err := s.repo.SessionByID(ctx, ssid)
	if err != nil {
		return entity.Seller{}, err
	}

	if time.Now().After(session.ExpiredAt) {
		return entity.Seller{}, fmt.Errorf("%w: session expired", ErrUnauthorized)
	}

	seller, err := s.repo.SellerByID(ctx, session.SellerID)
	if err != nil {
		return seller, err
	}

	return seller, nil
}

func (s *Service) Login(ctx context.Context, sellerLogin string) (entity.Session, error) {
	seller, err := s.repo.SellerByLogin(ctx, sellerLogin)
	if err != nil {
		if errors.Is(err, ErrNotFound) || !seller.ComparePassword(seller.Password) {
			return entity.Session{}, fmt.Errorf("%w: incorrect login or password", ErrUnauthorized)
		}

		return entity.Session{}, err
	}

	sess := entity.Session{
		ID:        uuid.New(),
		SellerID:  seller.ID,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(s.sessionAge),
	}

	err = s.repo.CreateSession(ctx, sess)
	if err != nil {
		return entity.Session{}, err
	}

	return sess, nil
}
