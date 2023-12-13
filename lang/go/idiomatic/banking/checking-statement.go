package banking

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/digital_document"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type CheckingStatement struct {
	Id            id.Id                       `json:"id,omitempty"`
	Account       BankAccount                 `json:"account,omitempty"`
	BeginOnDate   rfc3339date.Rfc3339Date     `json:"begin-on-date,omitempty"`
	EndOnDate     rfc3339date.Rfc3339Date     `json:"end-on-date,omitempty"`
	StatementDate rfc3339date.Rfc3339Date     `json:"statement-date,omitempty"`
	Capture       digital_document.DocumentV1 `json:"capture,omitempty"`
	Summary       CheckingStatementSummary    `json:"summary,omitempty"`
	Incoming      []StatementTransaction      `json:"incoming,omitempty"`
	OutGoing      []StatementTransaction      `json:"outgoing,omitempty"`
}

type CheckingStatementSummary struct {
}
