package email_address

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type EmailAddress struct {
	Id        id.Id                   `json:"id,omitempty"`
	Address   string                  `json:"address,omitempty"`
	StartDate rfc3339date.Rfc3339Date `json:"start-date,omitempty"`
	EndDate   rfc3339date.Rfc3339Date `json:"end-date,omitempty"`
}

var _ id.TypeNamer = &EmailAddress{}

func (t EmailAddress) TypeName() string {
	return id.TypeNamers.Dotted(EmailAddress{})
}
