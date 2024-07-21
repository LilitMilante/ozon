package main

import (
	"context"
	"log/slog"
	"os"

	"sellers-ms/internal/api"
	"sellers-ms/internal/app"
	"sellers-ms/internal/clients"
	"sellers-ms/internal/repository"
	"sellers-ms/internal/service"
)

func main() {
	ctx := context.Background()
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := app.NewConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	conn, err := app.NewPostgresClient(ctx, cfg.Postgres)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	repo := repository.NewRepository(conn)
	client := clients.NewProductsClient(cfg.ApiKey)
	s := service.NewService(repo, cfg.SessionAge)
	p := service.NewProducts(client)
	h := api.NewHandler(s, p)
	authMw := api.NewMiddleware(s, l)
	srv := api.NewServer(cfg.Port, h, authMw)

	l.Info("server started!", "port", cfg.Port)

	err = srv.Start()
	if err != nil {
		panic(err)
	}
}
