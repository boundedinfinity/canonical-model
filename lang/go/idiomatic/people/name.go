package people

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type NameMessage struct {
	Type string `json:"type,omitempty"`
	Name
}

func NewNameMessage(name Name) NameMessage {
	return NameMessage{Type: "people.Name", Name: name}
}

// ///////////////////////////////////////////////////
// Model
// ///////////////////////////////////////////////////

type Name struct {
	Id          id.Id      `json:"id,omitempty"`
	Prefixes    []Prefix   `json:"prefixes,omitempty"`
	GivenNames  []string   `json:"given-names,omitempty"`
	FamilyNames []string   `json:"family-names,omitempty"`
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
	return NewNameFormatter(
		PrefixFormats.Abbreviation,
		NameFormats.GivenNameFamilyName,
		SuffixFormats.Abbreviation,
		false,
	).Format(t)
}
