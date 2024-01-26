package people

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

// ///////////////////////////////////////////////////
// Model
// ///////////////////////////////////////////////////

var (
	ErrPrefixInvalidId           = errors.New("invalid prefix id")
	ErrPrefixInvalidDescription  = errors.New("invalid prefix description")
	ErrPrefixInvalidText         = errors.New("invalid prefix text")
	ErrPrefixInvalidAbbreviation = errors.New("invalid prefix abbreviation")
)

type Prefix struct {
	Id         id.Id            `json:"id,omitempty"`
	Text       string           `json:"text,omitempty"`
	Abbr       string           `json:"abbr,omitempty"`
	Descriptor PrefixDescriptor `json:"descriptor,omitempty"`
}

var _ id.TypeNamer = &Prefix{}

func (t Prefix) TypeName() string {
	return id.TypeNamers.Dotted(Prefix{})
}

func (t Prefix) String() string {
	if len(t.Abbr) > 0 {
		return fmt.Sprintf("%s (%s)", t.Text, t.Abbr)
	}

	return t.Text
}

func (t Prefix) Validate(groups ...string) error {
	if err := t.Id.Validate(groups...); err != nil {
		return fmt.Errorf("%w : %w", ErrPrefixInvalidId, err)
	}

	if len(t.Text) == 0 {
		return ErrPrefixInvalidText
	}

	return nil
}

// ///////////////////////////////////////////////////
// Assigned
// ///////////////////////////////////////////////////

var (
	ErrPrefixAssignedInvalidId     = errors.New("invalid prefix assigned id")
	ErrPrefixAssignedInvalidPrefix = errors.New("invalid prefix assigned prefix")
)

type PrefixAssigned struct {
	Id        id.Id                   `json:"id,omitempty"`
	Prefix    Prefix                  `json:"prefix,omitempty"`
	StartDate rfc3339date.Rfc3339Date `json:"start-date,omitempty"`
	EndDate   rfc3339date.Rfc3339Date `json:"end-date,omitempty"`
	Format    PrefixFormat            `json:"format,omitempty"`
}

func (t PrefixAssigned) String() string {
	if !t.StartDate.IsZero() && !t.EndDate.IsZero() {
		return fmt.Sprintf("%s (%s - %s)", t.Prefix, t.StartDate, t.EndDate)
	} else if !t.StartDate.IsZero() && t.EndDate.IsZero() {
		return fmt.Sprintf("%s (%s - present)", t.Prefix, t.StartDate)
	} else if t.StartDate.IsZero() && !t.EndDate.IsZero() {
		return fmt.Sprintf("%s (past - %s)", t.Prefix, t.EndDate)
	} else {
		return t.Prefix.String()
	}
}

func (t PrefixAssigned) Validate(groups ...string) error {
	if err := t.Id.Validate(groups...); err != nil {
		return errors.Join(ErrPrefixAssignedInvalidId, err)
	}

	if err := t.Prefix.Validate(groups...); err != nil {
		return errors.Join(ErrPrefixAssignedInvalidPrefix, err)
	}

	return nil
}

// ///////////////////////////////////////////////////
// Descriptor
// ///////////////////////////////////////////////////

var (
	ErrPrefixDescriptorParsableInvalidId = errors.New("invalid prefix descriptor parsable")
)

type PrefixDescriptor struct {
	Description   string       `json:"description,omitempty"`
	Parsable      []string     `json:"parsable,omitempty"`
	DefaultFormat PrefixFormat `json:"default-format,omitempty"`
}

func (t PrefixDescriptor) Validate(groups ...string) error {
	for i, parsable := range t.Parsable {
		if parsable == "" {
			return fmt.Errorf("%w [%d]", ErrPrefixDescriptorParsableInvalidId, i)
		}
	}

	return nil
}
