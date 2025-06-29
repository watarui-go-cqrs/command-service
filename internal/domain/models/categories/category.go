package categories

import (
	"command-service/internal/errs"

	"github.com/google/uuid"
)

type Category struct {
	id   *CategoryId
	name *CategoryName
}

func (c *Category) Id() *CategoryId {
	return c.id
}

func (c *Category) Name() *CategoryName {
	return c.name
}

func (c *Category) ChangeCategoryName(name *CategoryName) {
	c.name = name
}

func (c *Category) Equals(value *Category) (bool, *errs.DomainError) {
	if value == nil {
		return false, errs.NewDomainError("category is nil")
	}
	result := c.id.Equals(value.Id())
	return result, nil
}

func NewCategory(name *CategoryName) (*Category, *errs.DomainError) {
	if uid, err := uuid.NewRandom(); err != nil {
		return nil, errs.NewDomainError("failed to generate category ID")
	} else {
		if id, err := NewCategoryId(uid.String()); err != nil {
			return nil, errs.NewDomainError("failed to create category ID: " + err.Error())
		} else {
			return &Category{id: id, name: name}, nil
		}
	}
}

func BuildCategory(id *CategoryId, name *CategoryName) *Category {
	return &Category{id: id, name: name}
}
