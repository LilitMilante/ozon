package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"ozon/internal/entity"
	"ozon/internal/service"

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

	q := "SELECT id, full_name, login, created_at, updated_at FROM seller WHERE login = $1"

	err := r.db.QueryRow(ctx, q, login).
		Scan(
			&s.ID,
			&s.FullName,
			&s.Login,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return s, fmt.Errorf("get seller by login %v: %w", login, service.ErrNotFound)
		}

		return s, fmt.Errorf("get seller by login %v: %w", login, err)
	}

	return s, nil
}

func (r *Repository) SellerByID(ctx context.Context, id int64) (entity.Seller, error) {
	var s entity.Seller

	q := "SELECT id, full_name, login, created_at, updated_at FROM seller WHERE id = $1"

	err := r.db.QueryRow(ctx, q, id).
		Scan(
			&s.ID,
			&s.FullName,
			&s.Login,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return s, fmt.Errorf("get seller by login %v: %w", id, service.ErrNotFound)
		}

		return s, fmt.Errorf("get seller by login %v: %w", id, err)
	}

	return s, nil
}

// Session

func (r *Repository) Login(ctx context.Context, sess entity.Session) error {
	q := `
INSERT INTO sessions (id, patient_id, created_at, expired_at) VALUES ($1, $2, $3, $4)
`
	_, err := r.db.Exec(ctx, q, sess.ID, sess.SellerID, sess.CreatedAt, sess.ExpiredAt)
	return err
}

func (r *Repository) SessionByID(ctx context.Context, id string) (sess entity.Session, err error) {
	q := "SELECT id, seller_id, created_at, expired_at FROM sessions WHERE id = $1"

	err = r.db.QueryRow(ctx, q, id).
		Scan(&sess.ID, &sess.SellerID, &sess.CreatedAt, &sess.ExpiredAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sess, service.ErrNotFound
		}

		return sess, err
	}

	return sess, nil
}
