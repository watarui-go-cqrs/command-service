package categories

import (
	"command-service/command/errs"
	"unicode/utf8"
)

type CategoryName struct {
	value string
}

func (c *CategoryName) Value() string {
	return c.value
}

func NewCategoryName(value string) (*CategoryName, *errs.DomainError) {
	const MIN_LENGTH = 2
	const MAX_LENGTH = 20

	count := utf8.RuneCountInString(value)
	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errs.NewDomainError("category name must be between 3 and 50 characters long")
	}

	return &CategoryName{value: value}, nil
}
