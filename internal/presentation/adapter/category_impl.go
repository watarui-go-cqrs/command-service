package adapter

import (
	"command-service/internal/domain/models/categories"
	"command-service/internal/errs"

	"github.com/watarui-go-cqrs/pb/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type categoryAdapterImpl struct{}

func NewCategoryAdapterImpl() CategoryAdapter {
	return &categoryAdapterImpl{}
}

func (c *categoryAdapterImpl) ToEntity(param *pb.CategoryUpParam) (*categories.Category, error) {
	switch param.GetCrud() {
	case pb.CRUD_INSERT:
		name, err := categories.NewCategoryName(param.GetName())
		if err != nil {
			return nil, err
		}
		category, err := categories.NewCategory(name)
		if err != nil {
			return nil, err
		}
		return category, nil
	case pb.CRUD_UPDATE:
		id, err := categories.NewCategoryId(param.GetId())
		if err != nil {
			return nil, err
		}
		name, err := categories.NewCategoryName(param.GetName())
		if err != nil {
			return nil, err
		}
		return categories.BuildCategory(id, name), nil
	case pb.CRUD_DELETE:
		id, err := categories.NewCategoryId(param.GetId())
		if err != nil {
			return nil, err
		}
		return categories.BuildCategory(id, nil), nil
	default:
		return nil, errs.NewCRUDError("invalid CRUD operation")
	}
}

func (c *categoryAdapterImpl) ToResult(result any) *pb.CategoryUpResult {
	var (
		category *pb.Category
		err      *pb.Error
	)
	switch v := result.(type) {
	case *categories.Category:
		if v.Name() == nil {
			category = &pb.Category{
				Id:   v.Id().Value(),
				Name: "",
			}
		} else {
			category = &pb.Category{
				Id:   v.Id().Value(),
				Name: v.Name().Value(),
			}
		}
	case *errs.DomainError:
		err = &pb.Error{Type: "Domain Error", Message: v.Error()}
	case *errs.CRUDError:
		err = &pb.Error{Type: "CRUD Error", Message: v.Error()}
	case *errs.InternalError:
		err = &pb.Error{Type: "Internal Error", Message: v.Error()}
	}
	return &pb.CategoryUpResult{
		Category:  category,
		Error:     err,
		Timestamp: timestamppb.Now(),
	}
}
