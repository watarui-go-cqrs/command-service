package adapter

import (
	"command-service/internal/domain/models/categories"

	"github.com/watarui-go-cqrs/pb/pb"
)

type CategoryAdapter interface {
	ToEntity(param *pb.CategoryUpParam) (*categories.Category, error)
	ToResult(result any) *pb.CategoryUpResult
}
