package name

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/person/affix"
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

var (
	defaultFormatter = WesternFormatter()
)

func (this Name) String() string {
	return defaultFormatter.Format(this)
}
