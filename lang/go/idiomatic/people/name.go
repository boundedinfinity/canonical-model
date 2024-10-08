package people

// https://en.wikipedia.org/wiki/Personal_name
// https://www.family-historian.co.uk/help/fh7/nameformats.html
// https://en.wikipedia.org/wiki/Honorific
// https://en.wikipedia.org/wiki/English_honorifics
// https://www.genealogyexplained.com/what-is-suffix-in-a-name/
// https://emilypost.com/advice/mens-names-and-titles
// https://en.wikipedia.org/wiki/Suffix_(name)
// https://en.wikipedia.org/wiki/List_of_professional_designations_in_the_United_States
// https://www.nari.org/Certification-Accreditation/Individual-Certification
// https://en.wikipedia.org/wiki/Post-nominal_letters
// https://en.wikipedia.org/wiki/Surname
// https://en.wikipedia.org/wiki/Service_number_(United_States_Armed_Forces)#Prefix_and_suffix_codes

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

// ///////////////////////////////////////////////////
// Model
// ///////////////////////////////////////////////////

var (
	ErrNameInvalidId     = errors.New("invalid name id")
	ErrNameInvalidPrefix = errors.New("invalid name prefix")
	ErrNameInvalidFirst  = errors.New("invalid name first")
	ErrNameInvalidMiddle = errors.New("invalid name middle")
	ErrNameInvalidLast   = errors.New("invalid name last")
	ErrNameInvalidSuffix = errors.New("invalid name suffix")
)

type Name struct {
	Id       id.Id      `json:"id,omitempty"`
	Prefixes []Prefix   `json:"prefixes,omitempty"`
	Firsts   []string   `json:"firsts,omitempty"`
	Middles  []string   `json:"middles,omitempty"`
	Lasts    []string   `json:"lasts,omitempty"`
	Suffixes []Suffix   `json:"suffixes,omitempty"`
	Format   NameFormat `json:"format,omitempty"`
}

var _ id.TypeNamer = &Name{}

func (t Name) TypeName() string {
	return id.TypeNamers.Dotted(Name{})
}

func (t Name) String() string {
	return t.Full()
}

func (t Name) Validate() error {
	if err := t.Id.Validate(); err != nil {
		return fmt.Errorf("%w : %w", ErrNameInvalidId, err)
	}

	for i, prefix := range t.Prefixes {
		if err := prefix.Validate(); err != nil {
			return errors.Join(
				fmt.Errorf("%w[%d]", ErrNameInvalidPrefix, i),
				err,
			)
		}
	}

	for i, first := range t.Firsts {
		if first == "" {
			return fmt.Errorf("%w[%d]", ErrNameInvalidFirst, i)
		}
	}

	for i, middle := range t.Middles {
		if middle == "" {
			return fmt.Errorf("%w[%d]", ErrNameInvalidMiddle, i)
		}
	}

	for i, last := range t.Lasts {
		if last == "" {
			return fmt.Errorf("%w[%d]", ErrNameInvalidLast, i)
		}
	}

	for i, suffix := range t.Suffixes {
		if err := suffix.Validate(); err != nil {
			return errors.Join(
				fmt.Errorf("%w[%d]", ErrNameInvalidSuffix, i),
				err,
			)
		}
	}

	return nil
}

func (t Name) Prefix() string {
	strings := slicer.Map(
		func(_ int, item Prefix) string { return item.String() },
		t.Prefixes...,
	)

	return stringer.Join(" ", strings...)
}

func (t Name) First() string {
	return stringer.Join(" ", t.Firsts...)
}

func (t Name) Middle() string {
	return stringer.Join(" ", t.Middles...)
}

func (t Name) Last() string {
	return stringer.Join(" ", t.Lasts...)
}

func (t Name) Suffix() string {
	strings := slicer.Map(
		func(_ int, item Suffix) string { return item.String() },
		t.Suffixes...,
	)

	return stringer.Join(" ", strings...)
}

func (t Name) Full() string {
	return stringer.Join(
		" ",
		t.Prefix(),
		t.First(),
		t.Middle(),
		t.Last(),
		t.Suffix(),
	)
}
