package banking

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/audit"
	"github.com/boundedinfinity/schema/idiomatic/contact"
	"github.com/boundedinfinity/schema/idiomatic/currency"
	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/boundedinfinity/schema/idiomatic/mailing_address"
)

type BankingCheck struct {
	Id               id.Id                          `json:"id,omitempty"`
	PayTo            contact.Contact                `json:"pay-to,omitempty"`
	EndorsedBy       contact.Contact                `json:"endorsed-by,omitempty"`
	SingedBy         contact.Contact                `json:"signed-by,omitempty"`
	Amount           currency.CurrencyAmount        `json:"amount,omitempty"`
	Account          BankAccount                    `json:"account,omitempty"`
	Number           string                         `json:"number,omitempty"`
	Memo             string                         `json:"memo,omitempty"`
	Address          mailing_address.MailingAddress `json:"mailing-address,omitempty"`
	IssuedOnDate     rfc3339date.Rfc3339Date        `json:"issued-on-date,omitempty"`
	CashedOnDate     rfc3339date.Rfc3339Date        `json:"cashed-on-date,omitempty"`
	ValidThroughDate rfc3339date.Rfc3339Date        `json:"valid-through-date,omitempty"`
	Audit            audit.Audit                    `json:"audit,omitempty"`
}
