package service

import (
	"command-service/command/domain/models/categories"
	"context"
)

type CategoryService interface {
	Add(ctx context.Context, category *categories.Category) error
	Update(ctx context.Context, category *categories.Category) error
	Delete(ctx context.Context, categoryId *categories.CategoryId) error
}
