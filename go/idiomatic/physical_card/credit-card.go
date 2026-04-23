package physical_card

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/banking"
	"github.com/boundedinfinity/canonical_model/idiomatic/finanical"
	"github.com/boundedinfinity/canonical_model/idiomatic/id"
	"github.com/boundedinfinity/canonical_model/idiomatic/people"
	"github.com/boundedinfinity/canonical_model/idiomatic/phone"
	"github.com/boundedinfinity/canonical_model/idiomatic/website"
	"github.com/boundedinfinity/rfc3339date"
)

// https://en.wikipedia.org/wiki/Payment_card_number
type CreditCardNumber interface {
}

// https://en.wikipedia.org/wiki/Card_security_code
type SecurityCode interface {
}

type CreditCard struct {
	Id             id.Id                   `json:"id,omitempty"`
	Number         CreditCardNumber        `json:"number,omitempty"`
	Code           SecurityCode            `json:"code,omitempty"`
	Name           people.Name             `json:"name,omitempty"`
	IssueDate      rfc3339date.Rfc3339Date `json:"issue-date,omitempty"`
	ExpirationDate rfc3339date.Rfc3339Date `json:"expiration-date,omitempty"`
	OriginDate     rfc3339date.Rfc3339Date `json:"origin-date,omitempty"`
	InterestRate   finanical.InterestRate  `json:"interest-rate,omitempty"`
	Phone          phone.NanpNumber        `json:"phone,omitempty"`
	Issuer         banking.Bank            `json:"issuer,omitempty"`
	Site           website.PortalWebSite   `json:"site,omitempty"`
}

var _ id.TypeNamer = &CreditCard{}

func (t CreditCard) TypeName() string {
	return id.TypeNamers.Dotted(CreditCard{})
}
