package validatorer

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"golang.org/x/exp/constraints"
)

func Float[N constraints.Float]() *floatValidator[N] {
	return &floatValidator[N]{}
}

var (
	ErrFloatValidate   = errorer.New("float")
	errFloatValidateFn = errorer.Func(ErrValidate, ErrFloatValidate)
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ Validator = &floatValidator[float32]{}

type floatFn[N constraints.Float] func(string, N) error

type floatValidator[N constraints.Float] struct {
	fns []floatFn[N]
}

func (this floatValidator[N]) Validate(name string, v any) error {
	var errs []error

	if i, ok := v.(N); ok {
		for _, fn := range this.fns {
			if err := fn(name, i); err != nil {
				errs = append(errs, err)
			}
		}
	} else {
		var zero N
		errs = append(errs, errFloatValidateFn("input must be an %V", zero))
	}

	return errors.Join(errs...)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// https://zod.dev/api#numbers

func floatJoin[N constraints.Float](i N) string {
	return fmt.Sprintf("%f", i)
}

func (this *floatValidator[N]) OneOf(oneOf ...N) *floatValidator[N] {
	this.fns = append(this.fns, func(name string, v N) error {
		var err error

		if !slicer.Contains(v, oneOf...) {
			values := stringer.JoinFunc(oneOf, ", ", floatJoin)
			err = errFloatValidateFn("%s is not one of %s", name, values)
		}

		return err
	})

	return this
}

func (this *floatValidator[N]) NoneOf(noneOf ...N) *floatValidator[N] {
	this.fns = append(this.fns, func(name string, v N) error {
		var err error

		if slicer.Contains(v, noneOf...) {
			values := stringer.JoinFunc(noneOf, ", ", floatJoin)
			err = errFloatValidateFn("%s is one of %s", name, values)
		}

		return err
	})

	return this
}

func (this *floatValidator[N]) Max(max N) *floatValidator[N] {
	this.fns = append(this.fns, func(name string, v N) error {
		var err error

		if v > max {
			err = errFloatValidateFn("%s is greater than %f", name, max)
		}

		return err
	})

	return this
}

func (this *floatValidator[N]) Min(min N) *floatValidator[N] {
	this.fns = append(this.fns, func(name string, v N) error {
		var err error

		if v > min {
			err = errFloatValidateFn("%s is less than %f", name, min)
		}

		return err
	})

	return this
}
