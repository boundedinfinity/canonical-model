package payment_card

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/location/postal_code"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/modeller"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/person"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/person/name"
	"github.com/boundedinfinity/rfc3339date"
)

var _ modeller.Modeller = &Card{}

type Card struct {
	Id               ider.Id                 `json:"id"`
	Name             string                  `json:"name"`
	Person           person.Person           `json:"name-on-card"`
	NameFormatter    name.Formatter          `json:"name-formatter"`
	Kind             Kind                    `json:"kind"`
	Number           Number                  `json:"number"`
	IssueDate        rfc3339date.Rfc3339Date `json:"issue-date"`
	ActivationDate   rfc3339date.Rfc3339Date `json:"activation-date"`
	ExpirationDate   rfc3339date.Rfc3339Date `json:"expiration-date"`
	CardSecurityCode CardVerificationValue   `json:"card-security-code"`
	PersonalIdNumber string                  `json:"personal-id-number"`
	PostalCode       postal_code.PostalCode  `json:"postal-code"`
}

func (this Card) GetId() ider.Id {
	return this.Id
}

func (this Card) GetName() string {
	return this.Name
}

func (this Card) SecurityCode() CardVerificationValue {
	return this.CardSecurityCode
}

func (this Card) Ccv() CardVerificationValue {
	return this.CardSecurityCode
}

func (this Card) Csc() CardVerificationValue {
	return this.CardSecurityCode
}

func (this Card) Cvc() CardVerificationValue {
	return this.CardSecurityCode
}

func (this Card) Cid() CardVerificationValue {
	return this.CardSecurityCode
}

func (this Card) Pin() string {
	return this.PersonalIdNumber
}
