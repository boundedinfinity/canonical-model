package bookkeeping

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/audit"
	"github.com/boundedinfinity/schema/idiomatic/banking"
	"github.com/boundedinfinity/schema/idiomatic/business"
	"github.com/boundedinfinity/schema/idiomatic/digital_document"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type Account struct {
	Id          id.Id                `json:"id,omitempty"`
	Name        string               `json:"name,omitempty"`
	Description string               `json:"description,omitempty"`
	Parent      *Account             `json:"parent,omitempty"`
	Children    []*Account           `json:"children,omitempty"`
	Audit       audit.Audit          `json:"audit,omitempty"`
	BankAccount *banking.BankAccount `json:"bank-account,omitempty"`
	Business    *business.Business   `json:"business,omitempty"`
}

type Transfer struct {
	Id          id.Id                      `json:"id,omitempty"`
	From        *Account                   `json:"from,omitempty"`
	FromDate    rfc3339date.Rfc3339Date    `json:"from-date,omitempty"`
	To          *Account                   `json:"to,omitempty"`
	ToDate      rfc3339date.Rfc3339Date    `json:"to-date,omitempty"`
	Amount      float32                    `json:"amount,omitempty"`
	Description string                     `json:"description,omitempty"`
	Memo        string                     `json:"memo,omitempty"`
	CheckNumber int                        `json:"check-number,omitempty"`
	Attachment  *digital_document.Document `json:"attachement,omitempty"`
}

type Reconile struct {
	Id              id.Id                      `json:"id,omitempty"`
	StatementDate   rfc3339date.Rfc3339Date    `json:"statement-date,omitempty"`
	StartingBalance float32                    `json:"starting-balance,omitempty"`
	EndingBalance   float32                    `json:"ending-balance,omitempty"`
	Description     string                     `json:"description,omitempty"`
	Memo            string                     `json:"memo,omitempty"`
	Attachment      *digital_document.Document `json:"attachement,omitempty"`
}
