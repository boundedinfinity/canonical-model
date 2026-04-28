package validatorer

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/errorer"
)

type Validator interface {
	Validate(string, any) error
}

var (
	ErrValidate = errorer.New("validation error")
)
