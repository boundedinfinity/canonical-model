package bookkeeping

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/audit"
	"github.com/boundedinfinity/schema/idiomatic/authentication"
	"github.com/boundedinfinity/schema/idiomatic/banking"
	"github.com/boundedinfinity/schema/idiomatic/business"
	"github.com/boundedinfinity/schema/idiomatic/contact"
	"github.com/boundedinfinity/schema/idiomatic/digital_document"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type Account struct {
	Id          id.Id                `json:"id,omitempty"`
	Name        string               `json:"name,omitempty"`
	Code        int                  `json:"code,omitempty"`
	Description string               `json:"description,omitempty"`
	Parent      *Account             `json:"parent,omitempty"`
	Children    []*Account           `json:"children,omitempty"`
	Audit       audit.Audit          `json:"audit,omitempty"`
	BankAccount *banking.BankAccount `json:"bank-account,omitempty"`
	Business    *business.Business   `json:"business,omitempty"`
}

type Movement struct {
	Account Account                 `json:"account,omitempty"`
	Date    rfc3339date.Rfc3339Date `json:"date,omitempty"`
	Memo    string                  `json:"memo,omitempty"`
}

type Transaction struct {
	Id          id.Id                       `json:"id,omitempty"`
	Movement    []Movement                  `json:"movement,omitempty"`
	Amount      float32                     `json:"amount,omitempty"`
	Description string                      `json:"description,omitempty"`
	CheckNumber int                         `json:"check-number,omitempty"`
	Attachments []digital_document.Document `json:"attachements,omitempty"`
}

type Vendor struct {
	Id          id.Id           `json:"id,omitempty"`
	Number      string          `json:"number,omitempty"`
	Contact     contact.Contact `json:"contact,omitempty"`
	TaxIncluded bool            `json:"tax-included,omitempty"`
}

type Customer struct {
	Id          id.Id           `json:"id,omitempty"`
	Number      string          `json:"number,omitempty"`
	Contact     contact.Contact `json:"contact,omitempty"`
	TaxIncluded bool            `json:"tax-included,omitempty"`
}

type Employee struct {
	Id      id.Id                  `json:"id,omitempty"`
	Number  string                 `json:"number,omitempty"`
	Contact contact.Contact        `json:"contact,omitempty"`
	Account authentication.Account `json:"account,omitempty"`
}
