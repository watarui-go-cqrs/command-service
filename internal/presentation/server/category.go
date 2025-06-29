package server

import (
	"command-service/internal/application/service"
	"command-service/internal/presentation/adapter"
	"context"

	"github.com/watarui-go-cqrs/pb/pb"
)

type categoryServer struct {
	adapter adapter.CategoryAdapter
	service service.CategoryService
	pb.UnimplementedCategoryCommandServer
}

func NewCategoryServer(service service.CategoryService, adapter adapter.CategoryAdapter) pb.CategoryCommandServer {
	return &categoryServer{
		service: service,
		adapter: adapter,
	}
}

func (s *categoryServer) Create(ctx context.Context, param *pb.CategoryUpParam) (*pb.CategoryUpResult, error) {
	category, err := s.adapter.ToEntity(param)
	if err != nil {
		return s.adapter.ToResult(err), nil
	}
	if err := s.service.Add(ctx, category); err != nil {
		return s.adapter.ToResult(err), nil
	}
	return s.adapter.ToResult(category), nil
}

func (s *categoryServer) Update(ctx context.Context, param *pb.CategoryUpParam) (*pb.CategoryUpResult, error) {
	category, err := s.adapter.ToEntity(param)
	if err != nil {
		return s.adapter.ToResult(err), nil
	}
	if err := s.service.Update(ctx, category); err != nil {
		return s.adapter.ToResult(err), nil
	}
	return s.adapter.ToResult(category), nil
}

func (s *categoryServer) Delete(ctx context.Context, param *pb.CategoryUpParam) (*pb.CategoryUpResult, error) {
	category, err := s.adapter.ToEntity(param)
	if err != nil {
		return s.adapter.ToResult(err), nil
	}
	if err := s.service.Delete(ctx, category.Id()); err != nil {
		return s.adapter.ToResult(err), nil
	}
	return s.adapter.ToResult(category), nil
}
