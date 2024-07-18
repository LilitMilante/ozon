package main

import (
	"context"
	"log/slog"
	"os"

	"sellers-ms/internal/api"
	"sellers-ms/internal/app"
	"sellers-ms/internal/repository"
	"sellers-ms/internal/service"
)

func main() {
	ctx := context.Background()
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	c, err := app.NewConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	conn, err := app.NewPostgresClient(ctx, c.Postgres)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	repo := repository.NewRepository(conn)
	s := service.NewService(repo, c.SessionAge)
	h := api.NewHandler(s)
	authMw := api.NewMiddleware(s, l)
	srv := api.NewServer(c.Port, h, authMw)

	l.Info("server started!", "port", c.Port)

	err = srv.Start()
	if err != nil {
		panic(err)
	}
}
