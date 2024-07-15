package main

import (
	"context"
	"log"

	"ozon/internal/api"
	"ozon/internal/app"
	"ozon/internal/repository"
	"ozon/internal/service"
)

var ctx = context.Background()

func main() {
	c, err := app.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := app.NewPostgresClient(ctx, c.Database)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	repo := repository.NewRepository(conn)
	s := service.NewService(repo)
	h := api.NewHandler(s)
	authMw := api.NewAuthMiddleware(s)
	srv := api.NewServer(c.Port, h, authMw)

	log.Println("server started at:", c.Port)
	err = srv.Start()
	if err != nil {
		panic(err)
	}
}
