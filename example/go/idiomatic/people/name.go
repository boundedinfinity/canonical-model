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

	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
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
	Id       id.Id    `json:"id,omitempty,omitzero"`
	Prefixes []Prefix `json:"prefix,omitempty,omitzero"`
	Given    []string `json:"given,omitempty,omitzero"`
	Middle   []string `json:"middle,omitempty,omitzero"`
	Family   []string `json:"family,omitempty,omitzero"`
	Suffix   []Suffix `json:"suffix,omitempty,omitzero"`
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

	for i, first := range t.Given {
		if first == "" {
			return fmt.Errorf("%w[%d]", ErrNameInvalidFirst, i)
		}
	}

	for i, middle := range t.Middle {
		if middle == "" {
			return fmt.Errorf("%w[%d]", ErrNameInvalidMiddle, i)
		}
	}

	for i, last := range t.Family {
		if last == "" {
			return fmt.Errorf("%w[%d]", ErrNameInvalidLast, i)
		}
	}

	for i, suffix := range t.Suffix {
		if err := suffix.Validate(); err != nil {
			return errors.Join(
				fmt.Errorf("%w[%d]", ErrNameInvalidSuffix, i),
				err,
			)
		}
	}

	return nil
}

func (t Name) Full() string {
	prefix := slicer.Map(
		func(_ int, item Prefix) string { return item.String() },
		t.Prefixes...,
	)

	suffix := slicer.Map(
		func(_ int, item Suffix) string { return item.String() },
		t.Suffix...,
	)

	return stringer.Join(
		" ",
		stringer.Join(" ", prefix...),
		stringer.Join(" ", t.Given...),
		stringer.Join(" ", t.Middle...),
		stringer.Join(" ", t.Family...),
		stringer.Join(" ", suffix...),
	)
}

type StringQuery struct {
	Includes string
	Excludes string
}

type NameFind struct {
	Term   StringQuery
	Given  StringQuery
	Middle StringQuery
	Family StringQuery
}

type NameRepository interface {
	Find(query NameFind) []Name
}

var _ NameRepository = &NameMemoryRepository{}

type NameMemoryRepository struct {
	names []Name
}

func (this *NameMemoryRepository) Find(query NameFind) []Name {
	m := map[id.Id]Name{}

	match := func(name Name, list []string, fn func(int, string) bool) {
		items := slicer.Filter(fn, list...)
		if len(items) > 0 {
			if _, ok := m[name.Id]; !ok {
				m[name.Id] = name
			}
		}
	}

	includes := func(terms ...string) func(int, string) bool {
		return func(_ int, s string) bool {
			for _, term := range terms {
				if term != "" && stringer.Contains(s, term) {
					return true
				}
			}
			return false
		}
	}

	excludes := func(terms ...string) func(int, string) bool {
		return func(_ int, s string) bool {
			for _, term := range terms {
				if term != "" && !stringer.Contains(s, term) {
					return true
				}
			}
			return false
		}
	}

	for _, name := range this.names {
		match(name, name.Given, includes(query.Given.Includes, query.Term.Includes))
		match(name, name.Middle, includes(query.Middle.Includes, query.Term.Includes))
		match(name, name.Family, includes(query.Family.Includes, query.Term.Includes))

		match(name, name.Given, excludes(query.Given.Excludes, query.Term.Excludes))
		match(name, name.Middle, excludes(query.Middle.Excludes, query.Term.Excludes))
		match(name, name.Family, excludes(query.Family.Excludes, query.Term.Excludes))
	}

	return mapper.Values(m)
}
