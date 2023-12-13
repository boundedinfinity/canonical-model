package banking

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/contact"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type BankAccountNumber struct {
	Id     id.Id               `json:"id,omitempty"`
	Number string              `json:"number,omitempty"`
	Format AccountNumberFormat `json:"ordering,omitempty"`
}

func (t BankAccountNumber) String() string {
	var out string

	switch t.Format {
	case AccountNumberFormats.SpaceAfter3:
		out = stringer.Join(" ", stringer.Chunk(3, t.Number)...)
	case AccountNumberFormats.SpaceAfter4:
		out = stringer.Join(" ", stringer.Chunk(4, t.Number)...)
	default:
		out = t.Number
	}

	return out
}

type BankAccount struct {
	Id          id.Id                   `json:"id,omitempty"`
	Number      BankAccountNumber       `json:"number,omitempty"`
	OpeningDate rfc3339date.Rfc3339Date `json:"opening-date,omitempty"`
	ClosingDate rfc3339date.Rfc3339Date `json:"closing-date,omitempty"`
	Bank        Bank                    `json:"bank,omitempty"`
	Owner       contact.Contact         `json:"owner,omitempty"`
	Authorized  []contact.Contact       `json:"authorized,omitempty"`
}
