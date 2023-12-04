package people

import (
	"github.com/boundedinfinity/schema/idiomatic/id"
)

// ///////////////////////////////////////////////////
// Model
// ///////////////////////////////////////////////////

type Prefix struct {
	Id           id.Id        `json:"id,omitempty"`
	Description  string       `json:"description,omitempty"`
	Text         string       `json:"text,omitempty"`
	Abbreviation []string     `json:"abbreviation,omitempty"`
	Format       PrefixFormat `json:"format,omitempty"`
}

var _ id.TypeNamer = &Prefix{}

func (t Prefix) TypeName() string {
	return id.TypeNamers.Dotted(Prefix{})
}

func (t Prefix) String() string {
	return NewPrefixFormatter(PrefixFormats.Abbreviation).Format(t)
}

func (t Prefix) Validate(groups ...string) error {
	return nil
}
