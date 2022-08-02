package main

import (
	"fmt"
	"strings"

	"github.com/boundedinfinity/commons/slices"
	"github.com/boundedinfinity/optioner"
)

type Name struct {
	Prefixes    optioner.Option[[]Prefix]
	GivenNames  []string
	FamilyNames []string
	Suffixes    optioner.Option[[]Suffix]
	Ordering    NameOrdering
}

func (t Name) Validate() error {
	for i, prefix := range t.Prefixes.Get() {
		if err := prefix.Validate(); err != nil {
			return fmt.Errorf("prefixes[%v] %w", err, i)
		}
	}

	if len(t.GivenNames) < 1 {
		return fmt.Errorf("names must contain at least 1 item")
	}

	for i, name := range t.GivenNames {
		if len(name) < 1 || len(name) > 10 {
			return fmt.Errorf("givenNames[%v] must be between 1 and 10 characters in lenght", i)
		}
	}

	if len(t.FamilyNames) < 1 {
		return fmt.Errorf("names must contain at least 1 item")
	}

	for i, name := range t.FamilyNames {
		if len(name) < 1 || len(name) > 10 {
			return fmt.Errorf("familyNames[%v] must be between 1 and 10 characters in lenght", i)
		}
	}

	for i, suffix := range t.Suffixes.Get() {
		if err := suffix.Validate(); err != nil {
			return fmt.Errorf("suffixes[%v] %w", err, i)
		}
	}

	return nil
}

func (t Name) String() string {
	var name []string

	prefixesFn := func(p Prefix) string {
		return p.String()
	}

	suffixesFn := func(p Suffix) string {
		return p.String()
	}

	if t.Prefixes.Defined() {
		ss := slices.Map(t.Prefixes.Get(), prefixesFn)
		name = append(name, strings.Join(ss, " "))
	}

	switch t.Ordering {
	case FamilyNameGivenName:
		name = append(name, strings.Join(t.FamilyNames, " "))
		name = append(name, strings.Join(t.GivenNames, " "))
	default:
		name = append(name, strings.Join(t.GivenNames, " "))
		name = append(name, strings.Join(t.FamilyNames, " "))
	}

	if t.Suffixes.Defined() {
		ss := slices.Map(t.Suffixes.Get(), suffixesFn)
		name = append(name, strings.Join(ss, " "))
	}

	return strings.Join(name, " ")
}
