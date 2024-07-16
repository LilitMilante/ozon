package repository

import (
	"context"
	"errors"
	"fmt"

	"ozon/internal/entity"
	"ozon/internal/service"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{db: db}
}

// Seller methods

func (r *Repository) SellerByLogin(ctx context.Context, login string) (entity.Seller, error) {
	var s entity.Seller

	q := "SELECT id, full_name, login, created_at, updated_at FROM sellers WHERE login = $1"

	err := r.db.QueryRow(ctx, q, login).
		Scan(
			&s.ID,
			&s.FullName,
			&s.Login,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return s, fmt.Errorf("get seller by login %v: %w", login, service.ErrNotFound)
		}

		return s, fmt.Errorf("get seller by login %v: %w", login, err)
	}

	return s, nil
}

func (r *Repository) SellerByID(ctx context.Context, id uuid.UUID) (entity.Seller, error) {
	var s entity.Seller

	q := "SELECT id, full_name, login, created_at, updated_at FROM sellers WHERE id = $1"

	err := r.db.QueryRow(ctx, q, id).
		Scan(
			&s.ID,
			&s.FullName,
			&s.Login,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return s, fmt.Errorf("get seller by login %v: %w", id, service.ErrNotFound)
		}

		return s, fmt.Errorf("get seller by login %v: %w", id, err)
	}

	return s, nil
}

// Session

func (r *Repository) CreateSeller(ctx context.Context, s entity.Seller) (entity.Seller, error) {
	q := `
INSERT INTO sellers (id, full_name, login, password, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
`
	_, err := r.db.Exec(
		ctx,
		q,
		s.ID,
		s.FullName,
		s.Login,
		s.Password,
		s.CreatedAt,
		s.UpdatedAt)

	return s, err
}

func (r *Repository) CreateSession(ctx context.Context, s entity.Session) error {
	q := `
INSERT INTO sessions (id, seller_id, created_at, expired_at) VALUES ($1, $2, $3, $4)
`
	_, err := r.db.Exec(ctx, q, s.ID, s.SellerID, s.CreatedAt, s.ExpiredAt)
	return err
}

func (r *Repository) SessionByID(ctx context.Context, id string) (s entity.Session, err error) {
	q := "SELECT id, seller_id, created_at, expired_at FROM sessions WHERE id = $1"

	err = r.db.QueryRow(ctx, q, id).
		Scan(&s.ID, &s.SellerID, &s.CreatedAt, &s.ExpiredAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return s, service.ErrNotFound
		}

		return s, err
	}

	return s, nil
}
