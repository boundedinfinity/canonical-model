package physical_card

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/id"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/person/name"
	"github.com/boundedinfinity/rfc3339date"
)

// https://en.wikipedia.org/wiki/Payment_card_number
type CreditCardNumber interface {
}

// https://en.wikipedia.org/wiki/Card_security_code
type SecurityCode interface {
}

type CreditCard struct {
	Id             ider.Id                 `json:"id,omitempty"`
	Number         CreditCardNumber        `json:"number,omitempty"`
	Code           SecurityCode            `json:"code,omitempty"`
	Name           name.Name               `json:"name,omitempty"`
	IssueDate      rfc3339date.Rfc3339Date `json:"issue-date,omitempty"`
	ExpirationDate rfc3339date.Rfc3339Date `json:"expiration-date,omitempty"`
	OriginDate     rfc3339date.Rfc3339Date `json:"origin-date,omitempty"`
	// InterestRate   finanical.InterestRate  `json:"interest-rate,omitempty"`
	// Phone phone.NanpNumber `json:"phone,omitempty"`
	// Issuer         banking.Bank            `json:"issuer,omitempty"`
	// Site bookmark.Bookmark `json:"site,omitempty"`
}

var _ id.TypeNamer = &CreditCard{}

func (t CreditCard) TypeName() string {
	return id.TypeNamers.Dotted(CreditCard{})
}
