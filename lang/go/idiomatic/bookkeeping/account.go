package bookkeeping

import (
	"github.com/boundedinfinity/schema/idiomatic/authentication"
	"github.com/boundedinfinity/schema/idiomatic/banking"
	"github.com/boundedinfinity/schema/idiomatic/business"
	"github.com/boundedinfinity/schema/idiomatic/contact"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type Account struct {
	Id          id.Id                `json:"id,omitempty"`
	Name        string               `json:"name,omitempty"`
	Code        int                  `json:"code,omitempty"`
	Description string               `json:"description,omitempty"`
	Parent      *Account             `json:"parent,omitempty"`
	Children    []*Account           `json:"children,omitempty"`
	BankAccount *banking.BankAccount `json:"bank-account,omitempty"`
	Business    *business.Business   `json:"business,omitempty"`
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
