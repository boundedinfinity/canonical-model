package name

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/person/affix"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

type NameKind string

const (
	Unknown    NameKind = "unknown"
	Pseudonym  NameKind = "pseudonym"
	LegalName  NameKind = "legal-name"
	NickName   NameKind = "nick-name"
	StageName  NameKind = "stage-name"
	StreetName NameKind = "street-name"
	PenName    NameKind = "pen-name"
)

type Name struct {
	Id     ider.Id       `json:"id"`
	Given  []string      `json:"given"`
	Middle []string      `json:"middle"`
	Family []string      `json:"family"`
	Prefix []affix.Affix `json:"prefix"`
	Suffix []affix.Affix `json:"suffix"`
	Kind   NameKind      `json:"kind"`
}

func (this Name) First() []string {
	return this.Given
}

func (this Name) Last() []string {
	return this.Family
}

func (this Name) String() string {
	var names []string

	names = append(names, this.Given...)
	names = append(names, this.Middle...)
	names = append(names, this.Family...)

	return stringer.Join(" ", names)
}

// var zeroName = Name{}

// func (this Name) Zero() bool {
// 	return this == zeroName
// }
