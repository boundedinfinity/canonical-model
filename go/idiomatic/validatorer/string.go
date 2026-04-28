package validatorer

import (
	"errors"

	"github.com/boundedinfinity/go-commoner/idiomatic/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func String() *stringValidator {
	return &stringValidator{}
}

var (
	ErrStringValidate   = errorer.New("string")
	errStringValidateFn = errorer.Func(ErrValidate, ErrStringValidate)
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ Validator = &stringValidator{}

type stringFn func(string, string) error

type stringValidator struct {
	fns []stringFn
}

func (this stringValidator) Validate(name string, v any) error {
	var errs []error

	if s, ok := v.(string); ok {
		for _, fn := range this.fns {
			if err := fn(name, s); err != nil {
				errs = append(errs, err)
			}
		}
	} else {
		errs = append(errs, errStringValidateFn("input must be a string"))
	}

	return errors.Join(errs...)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// https://zod.dev/api#strings

func joinStringValues(vs []string) string {
	vs = slicer.Map(func(v string) string { return "'" + v + "'" }, vs...)
	return stringer.Join(", ", vs...)
}

func (this *stringValidator) StartsWith(startsWith string) *stringValidator {
	this.fns = append(this.fns, func(name, s string) error {
		var err error

		if !stringer.HasPrefix(s, startsWith) {
			err = errStringValidateFn("%s does not start with %s", name, startsWith)
		}

		return err
	})

	return this
}

func (this *stringValidator) NotStartsWith(notStartsWith string) *stringValidator {
	this.fns = append(this.fns, func(name, s string) error {
		var err error

		if stringer.HasPrefix(s, notStartsWith) {
			err = errStringValidateFn("%s starts with %s", name, notStartsWith)
		}

		return err
	})

	return this
}

func (this *stringValidator) EndsWith(endsWith string) *stringValidator {
	this.fns = append(this.fns, func(name, s string) error {
		var err error

		if !stringer.HasSuffix(s, endsWith) {
			err = errStringValidateFn("%s does not end with %s", name, endsWith)
		}

		return err
	})

	return this
}

func (this *stringValidator) NotEndsWith(notEndsWith string) *stringValidator {
	this.fns = append(this.fns, func(name, s string) error {
		var err error

		if stringer.HasSuffix(s, notEndsWith) {
			err = errStringValidateFn("%s ends with %s", name, notEndsWith)
		}

		return err
	})

	return this
}

func (this *stringValidator) Contains(contains string) *stringValidator {
	this.fns = append(this.fns, func(name, s string) error {
		var err error

		if !stringer.Contains(s, contains) {
			err = errStringValidateFn("%s does not contain %s", name, contains)
		}

		return err
	})

	return this
}

func (this *stringValidator) NotContains(notContains string) *stringValidator {
	this.fns = append(this.fns, func(name, s string) error {
		var err error

		if stringer.Contains(s, notContains) {
			err = errStringValidateFn("%s contains %s", name, notContains)
		}

		return err
	})

	return this
}

func (this *stringValidator) Max(max int) *stringValidator {
	this.fns = append(this.fns, func(name, s string) error {
		var err error

		if len(s) > max {
			err = errStringValidateFn("%s length is greater than %d", name, max)
		}

		return err
	})

	return this
}

func (this *stringValidator) Min(min int) *stringValidator {
	this.fns = append(this.fns, func(name, s string) error {
		var err error

		if len(s) < min {
			err = errStringValidateFn("%s length is less than %d", name, min)
		}

		return err
	})

	return this
}

func (this *stringValidator) TitleCased() *stringValidator {
	this.fns = append(this.fns, func(name, s string) error {
		var err error

		if !stringer.IsTitle(s) {
			err = errStringValidateFn("%s is not title cased %d", name)
		}

		return err
	})

	return this
}

func (this *stringValidator) UpperCased() *stringValidator {
	this.fns = append(this.fns, func(name, s string) error {
		var err error

		if !stringer.IsUpper(s) {
			err = errStringValidateFn("%s is not upper cased %d", name)
		}

		return err
	})

	return this
}

func (this *stringValidator) LowerCased() *stringValidator {
	this.fns = append(this.fns, func(name, s string) error {
		var err error

		if !stringer.IsLower(s) {
			err = errStringValidateFn("%s is not lower cased %d", name)
		}

		return err
	})

	return this
}

func (this *stringValidator) OneOf(oneOf ...string) *stringValidator {
	this.fns = append(this.fns, func(name, s string) error {
		var err error

		if !slicer.Contains(s, oneOf...) {
			values := joinStringValues(oneOf)
			err = errStringValidateFn("%s does not contain one of %s", name, values)
		}

		return err
	})

	return this
}

func (this *stringValidator) NoneOf(noneOf ...string) *stringValidator {
	this.fns = append(this.fns, func(name, s string) error {
		var err error

		if slicer.Contains(s, noneOf...) {
			values := joinStringValues(noneOf)
			err = errStringValidateFn("%s contains one of %s", name, values)
		}

		return err
	})

	return this
}
