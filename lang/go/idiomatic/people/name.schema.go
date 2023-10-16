package people

import (
	"strings"

	"github.com/google/uuid"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// ///////////////////////////////////////////////////
// Model
// ///////////////////////////////////////////////////

type Name struct {
	Id          uuid.UUID  `json:"id,omitempty"`
	Prefixes    []Prefix   `json:"prefixes,omitempty"`
	GivenNames  []string   `json:"givenNames,omitempty"`
	FamilyNames []string   `json:"familyNames,omitempty"`
	Suffixes    []Suffix   `json:"suffixes,omitempty"`
	Format      NameFormat `json:"ordering,omitempty"`
}

func (t Name) Validate(groups ...string) error {

	// for i, prefix := range t.Prefixes.Get() {
	// 	if err := prefix.Validate(); err != nil {
	// 		return fmt.Errorf("prefixes[%v] %w", i, err)
	// 	}
	// }

	// if len(t.GivenNames) < 1 {
	// 	return fmt.Errorf("names must contain at least 1 item")
	// }

	// for i, name := range t.GivenNames {
	// 	if len(name) < 1 || len(name) > 10 {
	// 		return fmt.Errorf("givenNames[%v] must be between 1 and 10 characters in lenght", i)
	// 	}
	// }

	// if len(t.FamilyNames) < 1 {
	// 	return fmt.Errorf("names must contain at least 1 item")
	// }

	// for i, name := range t.FamilyNames {
	// 	if len(name) < 1 || len(name) > 10 {
	// 		return fmt.Errorf("familyNames[%v] must be between 1 and 10 characters in lenght", i)
	// 	}
	// }

	// for i, suffix := range t.Suffixes.Get() {
	// 	if err := suffix.Validate(); err != nil {
	// 		return fmt.Errorf("suffixes[%v] %w", i, err)
	// 	}
	// }

	return nil
}

func (t Name) Prefix() string {
	strings := slicer.Map(func(item Prefix) string { return item.String() }, t.Prefixes...)
	return stringer.Join(" ", strings...)
}

func (t Name) GivenName() string {
	return stringer.Join(" ", t.GivenNames...)
}

func (t Name) FamilyName() string {
	return stringer.Join(" ", t.FamilyNames...)
}

func (t Name) Suffix() string {
	strings := slicer.Map(func(item Suffix) string { return item.String() }, t.Suffixes...)
	return stringer.Join(" ", strings...)
}

func (t Name) String() string {
	var names []string
	var ordering NameFormat

	names = append(names, t.Prefix())

	if t.Format != "" {
		ordering = t.Format
	} else {
		ordering = NameFormats.GivenNameFamilyName
	}

	switch ordering {
	case NameFormats.FamilyNameGivenName:
		names = append(names, t.FamilyName())
		names = append(names, t.GivenName())
	default:
		names = append(names, t.GivenName())
		names = append(names, t.FamilyName())
	}

	names = append(names, t.Suffix())

	s := strings.Join(names, " ")
	s = strings.Join(strings.Fields(s), " ")

	return s
}
