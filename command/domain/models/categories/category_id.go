package categories

import (
	"command-service/errs"
	"regexp"
	"unicode/utf8"
)

type CategoryId struct {
	value string
}

func (c *CategoryId) Value() string {
	return c.value
}

func (c *CategoryId) Equals(value *CategoryId) bool {
	if c == value {
		return true
	}
	return c.value == value.Value()
}

func NewCategoryId(value string) (*CategoryId, *errs.DomainError) {
	const LENGTH = 36
	const UUID_REGEX = "^([0-9a-fA-F]{8})-([0-9a-fA-F]{4})-([0-9a-fA-F]{4})-([0-9a-fA-F]{4})-([0-9a-fA-F]{12})$"

	if utf8.RuneCountInString(value) != LENGTH {
		return nil, errs.NewDomainError("category ID must be 36 characters long")
	}

	if !regexp.MustCompile(UUID_REGEX).MatchString(value) {
		return nil, errs.NewDomainError("category ID must be a valid UUID format")
	}

	return &CategoryId{value: value}, nil
}
