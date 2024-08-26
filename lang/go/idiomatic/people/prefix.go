package people

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/schema/idiomatic/id"
)

// ///////////////////////////////////////////////////
// Model
// ///////////////////////////////////////////////////

var _ id.TypeNamer = &Prefix{}

type Prefix struct {
	Id          id.Id    `json:"id,omitempty"`
	Text        string   `json:"text,omitempty"`
	Abbrs       []string `json:"abbrs,omitempty"`
	Description string   `json:"description,omitempty"`
}

func (this Prefix) TypeName() string {
	return id.TypeNamers.Dotted(Prefix{})
}

func (this Prefix) String() string {
	if len(this.Abbrs) > 0 {
		return this.Abbrs[0]
	}

	return this.Text
}

func (this Prefix) Validate() error {
	if err := this.Id.Validate(); err != nil {
		return &ErrPrefixDetails{err: ErrPrefixInvalidId, prefix: this}
	}

	if len(this.Text) == 0 {
		return &ErrPrefixDetails{err: ErrPrefixInvalidText, prefix: this}
	}

	return nil
}

// ///////////////////////////////////////////////////
// Errors
// ///////////////////////////////////////////////////

var (
	ErrPrefixInvalidId           = errors.New("invalid prefix id")
	ErrPrefixInvalidDescription  = errors.New("invalid prefix description")
	ErrPrefixInvalidText         = errors.New("invalid prefix text")
	ErrPrefixInvalidAbbreviation = errors.New("invalid prefix abbreviation")
)

type ErrPrefixDetails struct {
	err    error
	prefix Prefix
}

func (e ErrPrefixDetails) Error() string {
	return fmt.Sprintf("%s : %v", e.err.Error(), e.prefix)
}

func (e ErrPrefixDetails) Unwrap() error {
	return e.err
}
