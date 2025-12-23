package people

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/schema/idiomatic/id"
)

// ///////////////////////////////////////////////////
// Model
// ///////////////////////////////////////////////////

type AffixCategory struct {
	Id          id.Id  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Affix struct {
	Id          id.Id         `json:"id,omitempty"`
	Category    AffixCategory `json:"category,omitempty"`
	Text        string        `json:"text,omitempty"`
	Abbrs       []string      `json:"abbrs,omitempty"`
	Description string        `json:"description,omitempty"`
}

func (this Affix) String() string {
	if len(this.Abbrs) > 0 {
		return this.Abbrs[0]
	}

	return this.Text
}

func (this Affix) Validate() error {
	if err := this.Id.Validate(); err != nil {
		return &ErrAffixDetails{err: ErrAffixInvalidId, affix: this}
	}

	if this.Text == "" {
		return &ErrAffixDetails{err: ErrAffixInvalidText, affix: this}
	}

	return nil
}

// ///////////////////////////////////////////////////
// Errors
// ///////////////////////////////////////////////////

var (
	ErrAffixInvalidId           = errors.New("invalid id")
	ErrAffixInvalidDescription  = errors.New("invalid description")
	ErrAffixInvalidText         = errors.New("invalid text")
	ErrAffixInvalidAbbreviation = errors.New("invalid abbreviation")
)

type ErrAffixDetails struct {
	err   error
	affix Affix
}

func (this ErrAffixDetails) Error() string {
	return fmt.Sprintf("%s : %v", this.err.Error(), this.affix)
}

func (this ErrAffixDetails) Unwrap() error {
	return this.err
}

// ///////////////////////////////////////////////////
// Prefix
// ///////////////////////////////////////////////////

var _ id.TypeNamer = &Prefix{}

type Prefix struct {
	Affix
}

func (this Prefix) TypeName() string {
	return id.TypeNamers.Dotted(Prefix{})
}

// ///////////////////////////////////////////////////
// Suffix
// ///////////////////////////////////////////////////

type Suffix struct {
	Affix
}

var _ id.TypeNamer = &Suffix{}

func (t Suffix) TypeName() string {
	return id.TypeNamers.Dotted(Suffix{})
}
