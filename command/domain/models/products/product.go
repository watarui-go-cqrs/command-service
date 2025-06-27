package products

import (
	"command-service/command/domain/models/categories"
	"command-service/errs"

	"github.com/google/uuid"
)

type Product struct {
	id       *ProductId
	price    *ProductPrice
	name     *ProductName
	category *categories.Category
}

func (p *Product) Id() *ProductId {
	return p.id
}

func (p *Product) Price() *ProductPrice {
	return p.price
}

func (p *Product) Name() *ProductName {
	return p.name
}

func (p *Product) Category() *categories.Category {
	return p.category
}

func (p *Product) ChangeProductName(name *ProductName) {
	p.name = name
}

func (p *Product) ChangeProductPrice(price *ProductPrice) {
	p.price = price
}

func (p *Product) ChangeProductCategory(category *categories.Category) {
	p.category = category
}

func (p *Product) Equals(value *Product) (bool, *errs.DomainError) {
	if value == nil {
		return false, errs.NewDomainError("product is nil")
	}
	result := p.id.Equals(value.Id())
	return result, nil
}

func NewProduct(name *ProductName, price *ProductPrice, category *categories.Category) (*Product, *errs.DomainError) {
	if uid, err := uuid.NewRandom(); err != nil {
		return nil, errs.NewDomainError("failed to generate product ID")
	} else {
		if id, err := NewProductId(uid.String()); err != nil {
			return nil, errs.NewDomainError("failed to create product ID: " + err.Error())
		} else {
			return &Product{id: id, name: name, price: price, category: category}, nil
		}
	}
}

func BuildProduct(id *ProductId, name *ProductName, price *ProductPrice, category *categories.Category) *Product {
	return &Product{id: id, name: name, price: price, category: category}
}
