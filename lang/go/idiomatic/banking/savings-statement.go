package banking

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type SavingsStatement struct {
	Id              id.Id                   `json:"id,omitempty"`
	Account         BankAccount             `json:"account,omitempty"`
	BeginningOnDate rfc3339date.Rfc3339Date `json:"begin-on-date,omitempty"`
	EndOnDate       rfc3339date.Rfc3339Date `json:"end-on-date,omitempty"`
	StatementDate   rfc3339date.Rfc3339Date `json:"statement-date,omitempty"`
}
