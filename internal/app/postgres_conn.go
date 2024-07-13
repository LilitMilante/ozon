package app

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func NewPostgresClient(ctx context.Context, c DBConfig) (*pgx.Conn, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name)

	conn, err := pgx.Connect(ctx, psqlInfo)
	if err != nil {
		log.Fatal("connect to DB:", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
