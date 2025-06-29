package adapter

import (
	"command-service/internal/domain/models/categories"
	"command-service/internal/domain/models/products"
	"command-service/internal/errs"

	"github.com/watarui-go-cqrs/pb/pb"
)

type productAdapterImpl struct{}

func NewProductAdapterImpl() ProductAdapter {
	return &productAdapterImpl{}
}

func (p *productAdapterImpl) ToEntity(param *pb.ProductUpParam) (*products.Product, error) {
	switch param.GetCrud() {
	case pb.CRUD_INSERT:
		name, err := products.NewProductName(param.GetName())
		if err != nil {
			return nil, err
		}
		price, err := products.NewProductPrice(uint32(param.GetPrice()))
		if err != nil {
			return nil, err
		}
		categoryId, err := categories.NewCategoryId(param.GetCategoryId())
		if err != nil {
			return nil, err
		}
		product, err := products.NewProduct(name, price, categories.BuildCategory(categoryId, nil))
		if err != nil {
			return nil, err
		}
		return product, nil
	case pb.CRUD_UPDATE:
		id, err := products.NewProductId(param.GetId())
		if err != nil {
			return nil, err
		}
		name, err := products.NewProductName(param.GetName())
		if err != nil {
			return nil, err
		}
		price, err := products.NewProductPrice(uint32(param.GetPrice()))
		if err != nil {
			return nil, err
		}
		categoryId, err := categories.NewCategoryId(param.GetCategoryId())
		if err != nil {
			return nil, err
		}
		return products.BuildProduct(id, name, price, categories.BuildCategory(categoryId, nil)), nil
	case pb.CRUD_DELETE:
		id, err := products.NewProductId(param.GetId())
		if err != nil {
			return nil, err
		}
		return products.BuildProduct(id, nil, nil, nil), nil
	default:
		return nil, errs.NewCRUDError("invalid CRUD operation")
	}
}

func (p *productAdapterImpl) ToResult(result any) *pb.ProductUpResult {
	var (
		product *pb.Product
		err     *pb.Error
	)
	switch v := result.(type) {
	case *products.Product:
		var category *pb.Category
		if v.Category() == nil {
			category = &pb.Category{Id: "", Name: ""}
		} else {
			category = &pb.Category{Id: v.Category().Id().Value(), Name: v.Category().Name().Value()}
		}
		var name string = ""
		if v.Name() != nil {
			name = v.Name().Value()
		}
		var price int32 = 0
		if v.Price() != nil {
			price = int32(v.Price().Value())
		}
		product = &pb.Product{
			Id:       v.Id().Value(),
			Name:     name,
			Price:    float64(price),
			Category: category,
		}
	case errs.DomainError:
		err = &pb.Error{Type: "Domain Error", Message: v.Error()}
	case errs.CRUDError:
		err = &pb.Error{Type: "CRUD Error", Message: v.Error()}
	case errs.InternalError:
		err = &pb.Error{Type: "Internal Error", Message: v.Error()}
	}
	return &pb.ProductUpResult{Product: product, Error: err}
}
