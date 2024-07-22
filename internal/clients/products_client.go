package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"sellers-ms/internal/entity"

	"github.com/google/uuid"
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
		return entity.Product{}, fmt.Errorf("unexpected code %d", resp.StatusCode)
	}

	var product entity.Product

	err = json.NewDecoder(resp.Body).Decode(&product)
	if err != nil {
		return entity.Product{}, fmt.Errorf("decode response: %w", err)
	}

	return product, nil
}

func (pc *ProductsClient) ProductsBySellerID(ctx context.Context, sellerID uuid.UUID) ([]entity.Product, error) {
	url := fmt.Sprintf("http://localhost:5060/sellers/%s/products", sellerID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", pc.apiKey)

	resp, err := pc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return nil, fmt.Errorf("unexpected code %d", resp.StatusCode)
	}
	if resp.StatusCode >= http.StatusInternalServerError {
		return nil, fmt.Errorf("unexpected code %d", resp.StatusCode)
	}

	var products []entity.Product

	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return products, nil
}

func (pc *ProductsClient) AddProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	url := fmt.Sprintf("http://localhost:5060/products")

	jsonData, err := json.Marshal(product)
	if err != nil {
		return entity.Product{}, fmt.Errorf("marshal product in JSON: %w", err)
	}
	var body io.Reader
	body = bytes.NewReader(jsonData)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return entity.Product{}, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", pc.apiKey)

	resp, err := pc.client.Do(req)
	if err != nil {
		return entity.Product{}, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return entity.Product{}, fmt.Errorf("unexpected code %d", resp.StatusCode)
	}
	if resp.StatusCode >= http.StatusInternalServerError {
		return entity.Product{}, fmt.Errorf("unexpected code %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&product)
	if err != nil {
		return entity.Product{}, fmt.Errorf("decode response: %w", err)
	}

	return product, nil
}

func (pc *ProductsClient) UpdateProduct(ctx context.Context, product entity.UpdateProduct) (entity.UpdateProduct, error) {
	url := fmt.Sprintf("http://localhost:5060/products/update")

	jsonData, err := json.Marshal(product)
	if err != nil {
		return entity.UpdateProduct{}, fmt.Errorf("marshal product in JSON: %w", err)
	}
	var body io.Reader
	body = bytes.NewReader(jsonData)

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, body)
	if err != nil {
		return entity.UpdateProduct{}, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", pc.apiKey)

	resp, err := pc.client.Do(req)
	if err != nil {
		return entity.UpdateProduct{}, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return entity.UpdateProduct{}, fmt.Errorf("unexpected code %d", resp.StatusCode)
	}
	if resp.StatusCode >= http.StatusInternalServerError {
		return entity.UpdateProduct{}, fmt.Errorf("unexpected code %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&product)
	if err != nil {
		return entity.UpdateProduct{}, fmt.Errorf("decode response: %w", err)
	}

	return product, nil
}
