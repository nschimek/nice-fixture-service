package mocks

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type mockFieldError struct {
	validator.FieldError // by embedding this interface, we do not have to implment the rest of the methods
	tag   string
	field string
	param string
}

func NewMockFieldError(tag, field, param string) mockFieldError {
	return mockFieldError{tag: tag, field: field, param: param}
}

// only implement the methods our bind() function uses
func (e mockFieldError) Field() string { return e.field }
func (e mockFieldError) Tag() string { return e.tag }
func (e mockFieldError) Param() string { return e.param }
func (e mockFieldError) Error() string { return fmt.Sprintf("field=%s, tag=%s, param=%s", e.field, e.tag, e.param) }