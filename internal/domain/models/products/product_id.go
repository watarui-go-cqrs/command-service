package products

import (
	"command-service/internal/errs"
	"regexp"
	"unicode/utf8"
)

type ProductId struct {
	value string
}

func (p *ProductId) Value() string {
	return p.value
}

func (p *ProductId) Equals(value *ProductId) bool {
	if p == value {
		return true
	}
	return p.value == value.Value()
}

func NewProductId(value string) (*ProductId, *errs.DomainError) {
	const LENGTH = 36
	const UUID_REGEX = "^([0-9a-fA-F]{8})-([0-9a-fA-F]{4})-([0-9a-fA-F]{4})-([0-9a-fA-F]{4})-([0-9a-fA-F]{12})$"

	if utf8.RuneCountInString(value) != LENGTH {
		return nil, errs.NewDomainError("product ID must be 36 characters long")
	}

	if !regexp.MustCompile(UUID_REGEX).MatchString(value) {
		return nil, errs.NewDomainError("product ID must be a valid UUID format")
	}

	return &ProductId{value: value}, nil
}
