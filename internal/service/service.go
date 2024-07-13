package service

import (
	"context"
	"fmt"
	"time"

	"ozon/internal/entity"
	"ozon/internal/repository"

	"github.com/google/uuid"
)

type Repository interface {
	SellerByLogin(ctx context.Context, login string) (entity.Seller, error)
	SellerByID(ctx context.Context, id int64) (entity.Seller, error)

	SessionByID(ctx context.Context, id string) (entity.Session, error)
	Login(ctx context.Context, sess entity.Session) error
}

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) SellerByLogin(ctx context.Context, login string) (entity.Seller, error) {
	return s.repo.SellerByLogin(ctx, login)
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

func (s *Service) Login(ctx context.Context, sellerID int64) (entity.Session, error) {
	sess := entity.Session{
		ID:        uuid.New(),
		SellerID:  sellerID,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(time.Minute * 1),
	}

	err := s.repo.Login(ctx, sess)
	if err != nil {
		return entity.Session{}, err
	}

	return sess, nil
}
