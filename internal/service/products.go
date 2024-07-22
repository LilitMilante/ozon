package service

import (
	"context"

	"sellers-ms/internal/clients"
	"sellers-ms/internal/entity"

	"github.com/google/uuid"
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

func (p *Products) ProductsBySellerID(ctx context.Context, sellerID uuid.UUID) ([]entity.Product, error) {
	return p.client.ProductsBySellerID(ctx, sellerID)
}

func (p *Products) AddProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	return p.client.AddProduct(ctx, product)
}

func (p *Products) UpdateProduct(ctx context.Context, product entity.UpdateProduct) (entity.UpdateProduct, error) {
	return p.client.UpdateProduct(ctx, product)
}
