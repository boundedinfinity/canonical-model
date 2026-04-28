package validatorer

import (
	"errors"

	"github.com/boundedinfinity/go-commoner/idiomatic/errorer"
)

func Boolean() *booleanValidator {
	return &booleanValidator{}
}

var (
	ErrBooleanValidate   = errorer.New("boolean")
	errBooleanValidateFn = errorer.Func(ErrValidate, ErrBooleanValidate)
)

var _ Validator = &booleanValidator{}

type booleanValidator struct {
	isTrue  bool
	isFalse bool
}

// https://zod.dev/api#booleans

func (this booleanValidator) Validate(name string, v any) error {
	var errs []error
	b := v.(bool)

	if this.isTrue && !b {
		errs = append(errs, errBooleanValidateFn("%s is false", name))
	}

	if this.isFalse && b {
		errs = append(errs, errBooleanValidateFn("%s is true", name))
	}

	return errors.Join(errs...)
}

func (this *booleanValidator) True() *booleanValidator {
	this.isTrue = !this.isTrue
	return this
}

func (this *booleanValidator) False() *booleanValidator {
	this.isFalse = !this.isFalse
	return this
}
