package adapter

import (
	"command-service/internal/domain/models/products"

	"github.com/watarui-go-cqrs/pb/pb"
)

type ProductAdapter interface {
	ToEntity(param *pb.ProductUpParam) (*products.Product, error)
	ToResult(result any) *pb.ProductUpResult
}
