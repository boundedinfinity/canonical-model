package validatorer

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"golang.org/x/exp/constraints"
)

func Integer[N constraints.Integer]() *integerValidator[N] {
	return &integerValidator[N]{}
}

var (
	ErrIntegerValidate   = errorer.New("integer")
	errIntegerValidateFn = errorer.Func(ErrValidate, ErrIntegerValidate)
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ Validator = &integerValidator[int]{}

type integerFn[N constraints.Integer] func(string, N) error

type integerValidator[N constraints.Integer] struct {
	fns []integerFn[N]
}

func (this integerValidator[N]) Validate(name string, v any) error {
	var errs []error

	if i, ok := v.(N); ok {
		for _, fn := range this.fns {
			if err := fn(name, i); err != nil {
				errs = append(errs, err)
			}
		}
	} else {
		var zero N
		errs = append(errs, errIntegerValidateFn("input must be an %V", zero))
	}

	return errors.Join(errs...)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// https://zod.dev/api#numbers

func integerJoin[N constraints.Integer](i N) string {
	return fmt.Sprintf("%d", i)
}

func (this *integerValidator[N]) OneOf(oneOf ...N) *integerValidator[N] {
	this.fns = append(this.fns, func(name string, v N) error {
		var err error

		if !slicer.Contains(v, oneOf...) {
			values := stringer.JoinFunc(oneOf, ", ", integerJoin)
			err = errIntegerValidateFn("%s is not one of %s", name, values)
		}

		return err
	})

	return this
}

func (this *integerValidator[N]) NoneOf(noneOf ...N) *integerValidator[N] {
	this.fns = append(this.fns, func(name string, v N) error {
		var err error

		if slicer.Contains(v, noneOf...) {
			values := stringer.JoinFunc(noneOf, ", ", integerJoin)
			err = errIntegerValidateFn("%s is one of %s", name, values)
		}

		return err
	})

	return this
}

func (this *integerValidator[N]) Max(max N) *integerValidator[N] {
	this.fns = append(this.fns, func(name string, v N) error {
		var err error

		if v > max {
			err = errIntegerValidateFn("%s is greater than %d", name, max)
		}

		return err
	})

	return this
}

func (this *integerValidator[N]) Min(min N) *integerValidator[N] {
	this.fns = append(this.fns, func(name string, v N) error {
		var err error

		if v > min {
			err = errIntegerValidateFn("%s is less than %d", name, min)
		}

		return err
	})

	return this
}

func (this *integerValidator[N]) MultipleOf(multipleOf N) *integerValidator[N] {
	this.fns = append(this.fns, func(name string, v N) error {
		var err error

		if r := v % multipleOf; r != 0 {
			err = errIntegerValidateFn("%s is not a multiple of %d", name, multipleOf)
		}

		return err
	})

	return this
}
