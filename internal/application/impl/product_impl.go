package impl

import (
	"command-service/internal/application/service"
	"command-service/internal/domain/models/products"
	"command-service/internal/infrastructure/sqlboiler/handler"
	"context"
)

type productServiceImpl struct {
	rep products.ProductRepository
	transaction
}

func NewProductServiceImpl(rep products.ProductRepository) service.ProductService {
	return &productServiceImpl{rep: rep}
}

func (ins *productServiceImpl) Add(ctx context.Context, product *products.Product) error {
	tran, err := ins.begin(ctx)
	if err != nil {
		return handler.DBErrHandler(err)
	}
	defer func() {
		err = ins.complete(tran, err)
	}()
	if err = ins.rep.Exists(ctx, tran, product); err != nil {
		return err
	}
	if err = ins.rep.Create(ctx, tran, product); err != nil {
		return err
	}
	return err
}

func (ins *productServiceImpl) Update(ctx context.Context, product *products.Product) error {
	tran, err := ins.begin(ctx)
	if err != nil {
		return handler.DBErrHandler(err)
	}
	defer func() {
		err = ins.complete(tran, err)
	}()
	if err = ins.rep.UpdateById(ctx, tran, product); err != nil {
		return err
	}
	return err
}

func (ins *productServiceImpl) Delete(ctx context.Context, product *products.Product) error {
	tran, err := ins.begin(ctx)
	if err != nil {
		return handler.DBErrHandler(err)
	}
	defer func() {
		err = ins.complete(tran, err)
	}()
	if err = ins.rep.DeleteById(ctx, tran, product.Id()); err != nil {
		return err
	}
	return err
}
