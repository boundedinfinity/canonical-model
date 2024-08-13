package mailing_address

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/reflecter"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/boundedinfinity/schema/idiomatic/location"
)

type MailingAddress struct {
	Id    id.Id          `json:"id,omitempty"`
	Lines []string       `json:"lines,omitempty"`
	City  location.City  `json:"city,omitempty"`
	State location.State `json:"state,omitempty"`
	Zip   Zip            `json:"zip,omitempty"`
}

var _ id.TypeNamer = &MailingAddress{}

func (t MailingAddress) TypeName() string {
	return id.TypeNamers.Dotted(MailingAddress{})
}

func (t MailingAddress) IsZero() bool {
	return reflecter.IsZero[MailingAddress](t)
}

func (t MailingAddress) String() string {
	var lines []string

	for _, line := range t.Lines {
		if line != "" {
			lines = append(lines, line)
		}
	}

	lines = append(lines, t.City.String())
	lines = append(lines, t.State.String())
	lines = append(lines, t.Zip.String())

	return stringer.Join(", ", lines...)
}
