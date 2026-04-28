package rfc5322

import (
	"regexp"

	"github.com/boundedinfinity/go-commoner/idiomatic/errorer"
)

// https://datatracker.ietf.org/doc/html/rfc5322#section-3.4.1

type EmailAddress string

var (
	ErrRrc5322      = errorer.New("RFC 5332 email address")
	errErrRrc5322Fn = errorer.Func(ErrRrc5322)

	emailRegex = regexp.MustCompile(`/^(?!\.)(?!.*\.\.)([a-z0-9_'+\-\.]*)[a-z0-9_+-]@([a-z0-9][a-z0-9\-]*\.)+[a-z]{2,}$/i`)
)

func (this EmailAddress) Validate() error {
	if !emailRegex.MatchString(string(this)) {
		return errErrRrc5322Fn("isn't a valid email address")
	}

	return nil
}
