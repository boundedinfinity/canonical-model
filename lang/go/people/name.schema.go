package people

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/boundedinfinity/optioner"
)

// ///////////////////////////////////////////////////
// Model
// ///////////////////////////////////////////////////

type Name struct {
	Id          optioner.Option[uuid.UUID]    `json:"id,omitempty"`
	Prefixes    optioner.Option[[]Prefix]     `json:"prefixes,omitempty"`
	GivenNames  []string                      `json:"givenNames,omitempty"`
	FamilyNames []string                      `json:"familyNames,omitempty"`
	Suffixes    optioner.Option[[]Suffix]     `json:"suffixes,omitempty"`
	Format      optioner.Option[NameOrdering] `json:"ordering,omitempty"`
}

func (t Name) Validate(groups ...string) error {
	for i, prefix := range t.Prefixes.Get() {
		if err := prefix.Validate(); err != nil {
			return fmt.Errorf("prefixes[%v] %w", i, err)
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
			return fmt.Errorf("suffixes[%v] %w", i, err)
		}
	}

	return nil
}

func (t Name) String() string {
	var names []string

	if t.Prefixes.Defined() && len(t.Prefixes.Get()) > 0 {
		names = append(names, toStrings(t.Prefixes.Get())...)
	}

	ordering := t.Format.OrElse(GivenNameFamilyName)

	switch ordering {
	case FamilyNameGivenName:
		names = append(names, strings.Join(t.FamilyNames, " "))
		names = append(names, strings.Join(t.GivenNames, " "))
	default:
		names = append(names, strings.Join(t.GivenNames, " "))
		names = append(names, strings.Join(t.FamilyNames, " "))
	}

	if t.Suffixes.Defined() && len(t.Suffixes.Get()) > 0 {
		names = append(names, toStrings(t.Suffixes.Get())...)
	}

	s := strings.Join(names, " ")
	s = strings.Join(strings.Fields(s), " ")

	return s
}

// ///////////////////////////////////////////////////
// Builder
// ///////////////////////////////////////////////////

type nameFn func(*Name) error

type nameBuilder struct {
	fns []nameFn
}

func BuildName() *nameBuilder {
	return &nameBuilder{
		fns: make([]nameFn, 0),
	}
}

func (t *nameBuilder) Build() (Name, error) {
	var v Name

	for _, fn := range t.fns {
		if err := fn(&v); err != nil {
			return v, err
		}
	}

	return v, nil
}

func (t *nameBuilder) Must() Name {
	v, err := t.Build()

	if err != nil {
		panic(err)
	}

	return v
}

func (t *nameBuilder) Id(v string) *nameBuilder {
	t.fns = append(t.fns, func(p *Name) error {
		id, err := uuid.Parse(v)

		if err != nil {
			return err
		}

		p.Id = optioner.Some(id)
		return nil
	})

	return t
}

func (t *nameBuilder) GivenName(v string) *nameBuilder {
	t.fns = append(t.fns, func(p *Name) error {
		p.GivenNames = sAppend(p.GivenNames, v)
		return nil
	})

	return t
}

func (t *nameBuilder) FamilyName(v string) *nameBuilder {
	t.fns = append(t.fns, func(p *Name) error {
		p.FamilyNames = sAppend(p.FamilyNames, v)
		return nil
	})

	return t
}

func (t *nameBuilder) Prefix(v Prefix) error {
	t.fns = append(t.fns, func(p *Name) error {
		p.Prefixes = soAppend(p.Prefixes, v)
		return nil
	})

	return nil
}

func (t *nameBuilder) Suffix(v Suffix) error {
	t.fns = append(t.fns, func(p *Name) error {
		p.Suffixes = soAppend(p.Suffixes, v)
		return nil
	})

	return nil
}

func (t *nameBuilder) Format(v NameOrdering) *nameBuilder {
	t.fns = append(t.fns, func(p *Name) error {
		p.Format = optioner.Some(v)
		return nil
	})

	return t
}
