package products

import (
	"command-service/internal/errs"
	"fmt"
	"unicode/utf8"
)

type ProductName struct {
	value string
}

func (p *ProductName) Value() string {
	return p.value
}

func NewProductName(value string) (*ProductName, *errs.DomainError) {
	const MIN_LENGTH = 5
	const MAX_LENGTH = 30

	count := utf8.RuneCountInString(value)
	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("product name must be between %d and %d characters long", MIN_LENGTH, MAX_LENGTH))
	}

	return &ProductName{value: value}, nil
}
