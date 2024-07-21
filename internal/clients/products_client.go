package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"sellers-ms/internal/entity"
)

type ProductsClient struct {
	client *http.Client
	apiKey string
}

func NewProductsClient(apiKey string) *ProductsClient {
	return &ProductsClient{
		client: &http.Client{
			Timeout: time.Second * 5,
		},
		apiKey: apiKey,
	}
}

func (pc *ProductsClient) ProductByID(ctx context.Context, id int64) (entity.Product, error) {
	url := fmt.Sprintf("http://localhost:5060/products/%d", id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return entity.Product{}, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", pc.apiKey)

	resp, err := pc.client.Do(req)
	if err != nil {
		return entity.Product{}, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return entity.Product{}, entity.ErrNotFound
	}
	if resp.StatusCode >= http.StatusInternalServerError {
		return entity.Product{}, entity.ErrInternal
	}

	var product entity.Product

	err = json.NewDecoder(resp.Body).Decode(&product)
	if err != nil {
		return entity.Product{}, fmt.Errorf("decode response: %w", err)
	}

	return product, nil
}
