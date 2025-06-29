package service

import (
	"command-service/internal/domain/models/products"
	"context"
)

type ProductService interface {
	Add(ctx context.Context, product *products.Product) error
	Update(ctx context.Context, product *products.Product) error
	Delete(ctx context.Context, product *products.Product) error
}
