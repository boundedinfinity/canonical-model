package payment_card

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/idiomatic/person/name"
	"github.com/boundedinfinity/canonical_model/idiomatic/util/modeller"
	"github.com/boundedinfinity/rfc3339date"
)

var _ modeller.Modeller = &CardModel{}

type CardModel struct {
	Id               ider.Id                 `json:"id"`
	Name             string                  `json:"name"`
	NameOnCard       name.Name               `json:"name-on-card"`
	NameFormatter    name.Formatter          `json:"name-formatter"`
	Kind             Kind                    `json:"kind"`
	Number           Number                  `json:"number"`
	IssueDate        rfc3339date.Rfc3339Date `json:"issue-date"`
	ActivationDate   rfc3339date.Rfc3339Date `json:"activation-date"`
	ExpirationDate   rfc3339date.Rfc3339Date `json:"expiration-date"`
	CardSecurityCode CardVerificationValue   `json:"card-security-code"`
	PersonalIdNumber string                  `json:"personal-id-number"`
}

func (this CardModel) GetId() ider.Id {
	return this.Id
}

func (this CardModel) GetName() string {
	return this.Name
}

func (this CardModel) SecurityCode() CardVerificationValue {
	return this.CardSecurityCode
}

func (this CardModel) Ccv() CardVerificationValue {
	return this.CardSecurityCode
}

func (this CardModel) Csc() CardVerificationValue {
	return this.CardSecurityCode
}

func (this CardModel) Cvc() CardVerificationValue {
	return this.CardSecurityCode
}

func (this CardModel) Cid() CardVerificationValue {
	return this.CardSecurityCode
}

func (this CardModel) Pin() string {
	return this.PersonalIdNumber
}
