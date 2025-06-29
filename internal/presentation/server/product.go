package server

import (
	"command-service/internal/application/service"
	"command-service/internal/presentation/adapter"
	"context"

	"github.com/watarui-go-cqrs/pb/pb"
)

type productServer struct {
	adapter adapter.ProductAdapter
	service service.ProductService
	pb.UnimplementedProductCommandServer
}

func NewProductServer(service service.ProductService, adapter adapter.ProductAdapter) pb.ProductCommandServer {
	return &productServer{
		service: service,
		adapter: adapter,
	}
}

func (s *productServer) Create(ctx context.Context, param *pb.ProductUpParam) (*pb.ProductUpResult, error) {
	product, err := s.adapter.ToEntity(param)
	if err != nil {
		return s.adapter.ToResult(err), nil
	}
	if err := s.service.Add(ctx, product); err != nil {
		return s.adapter.ToResult(err), nil
	}
	return s.adapter.ToResult(product), nil
}

func (s *productServer) Update(ctx context.Context, param *pb.ProductUpParam) (*pb.ProductUpResult, error) {
	product, err := s.adapter.ToEntity(param)
	if err != nil {
		return s.adapter.ToResult(err), nil
	}
	if err := s.service.Update(ctx, product); err != nil {
		return s.adapter.ToResult(err), nil
	}
	return s.adapter.ToResult(product), nil
}

func (s *productServer) Delete(ctx context.Context, param *pb.ProductUpParam) (*pb.ProductUpResult, error) {
	product, err := s.adapter.ToEntity(param)
	if err != nil {
		return s.adapter.ToResult(err), nil
	}
	if err := s.service.Delete(ctx, product); err != nil {
		return s.adapter.ToResult(err), nil
	}
	return s.adapter.ToResult(product), nil
}
