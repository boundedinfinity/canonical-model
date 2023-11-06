package physical_card

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/audit"
	"github.com/boundedinfinity/schema/idiomatic/banking"
	"github.com/boundedinfinity/schema/idiomatic/finanical"
	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/boundedinfinity/schema/idiomatic/people"
	"github.com/boundedinfinity/schema/idiomatic/phone"
	"github.com/boundedinfinity/schema/idiomatic/website"
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
	Audit          audit.Audit             `json:"audit,omitempty"`
}
