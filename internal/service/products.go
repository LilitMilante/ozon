package service

import (
	"context"

	"sellers-ms/internal/clients"
	"sellers-ms/internal/entity"
)

type Products struct {
	client *clients.ProductsClient
}

func NewProducts(c *clients.ProductsClient) *Products {
	return &Products{
		client: c,
	}
}

func (p *Products) ProductByID(ctx context.Context, id int64) (entity.Product, error) {
	return p.client.ProductByID(ctx, id)
}
