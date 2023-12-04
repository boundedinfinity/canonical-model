package people

import (
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type Suffix struct {
	Id           id.Id        `json:"id,omitempty"`
	Description  string       `json:"description,omitempty"`
	Text         string       `json:"text,omitempty"`
	Abbreviation []string     `json:"abbreviation,omitempty"`
	Format       SuffixFormat `json:"format,omitempty"`
}

var _ id.TypeNamer = &Suffix{}

func (t Suffix) TypeName() string {
	return id.TypeNamers.Dotted(Suffix{})
}

func (t Suffix) Validate(groups ...string) error {
	return nil
}

func (t Suffix) String() string {
	return NewSuffixFormatter(SuffixFormats.Abbreviation).Format(t)
}
