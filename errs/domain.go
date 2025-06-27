package errs

type DomainError struct {
	message string
}

func (e *DomainError) Error() string {
	return e.message
}

func NewDomainError(message string) *DomainError {
	return &DomainError{message: message}
}
