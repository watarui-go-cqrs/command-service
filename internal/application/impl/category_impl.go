package impl

import (
	"command-service/internal/application/service"
	"command-service/internal/domain/models/categories"
	"context"
)

type categoryServiceImpl struct {
	rep categories.CategoryRepository
	transaction
}

func NewCategoryServiceImpl(rep categories.CategoryRepository) service.CategoryService {
	return &categoryServiceImpl{
		rep: rep,
	}
}

func (c *categoryServiceImpl) Add(ctx context.Context, category *categories.Category) error {
	tran, err := c.begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = c.complete(tran, err)
	}()

	if err := c.rep.Exists(ctx, tran, category); err != nil {
		return err
	}
	if err := c.rep.Create(ctx, tran, category); err != nil {
		return err
	}

	return nil
}

func (c *categoryServiceImpl) Update(ctx context.Context, category *categories.Category) error {
	tran, err := c.begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = c.complete(tran, err)
	}()

	if err := c.rep.UpdateById(ctx, tran, category); err != nil {
		return err
	}

	return nil
}

func (c *categoryServiceImpl) Delete(ctx context.Context, categoryId *categories.CategoryId) error {
	tran, err := c.begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = c.complete(tran, err)
	}()

	if err := c.rep.DeleteById(ctx, tran, categoryId); err != nil {
		return err
	}

	return nil
}
