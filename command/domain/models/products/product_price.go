package products

import (
	"command-service/command/errs"
	"fmt"
)

type ProductPrice struct {
	value uint32
}

func (p *ProductPrice) Value() uint32 {
	return p.value
}

func NewProductPrice(value uint32) (*ProductPrice, *errs.DomainError) {
	const MIN_PRICE = 50
	const MAX_PRICE = 10000

	if value < MIN_PRICE || value > MAX_PRICE {
		return nil, errs.NewDomainError(fmt.Sprintf("product price must be between %d and %d", MIN_PRICE, MAX_PRICE))
	}

	return &ProductPrice{value: value}, nil
}
