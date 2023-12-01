package banking

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type BankAccountNumber struct {
	Id     id.Id               `json:"id,omitempty"`
	Number string              `json:"number,omitempty"`
	Format AccountNumberFormat `json:"ordering,omitempty"`
}

type BankAccount struct {
	Id          id.Id                   `json:"id,omitempty"`
	Number      BankAccountNumber       `json:"number,omitempty"`
	OpeningDate rfc3339date.Rfc3339Date `json:"opening-date,omitempty"`
	ClosingDate rfc3339date.Rfc3339Date `json:"closing-date,omitempty"`
	Bank        Bank                    `json:"bank,omitempty"`
}
