package business

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/reflecter"
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/contact"
	"github.com/boundedinfinity/schema/idiomatic/email_address"
	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/boundedinfinity/schema/idiomatic/location"
	"github.com/boundedinfinity/schema/idiomatic/mailing_address"
	"github.com/boundedinfinity/schema/idiomatic/phone"
)

type Business struct {
	Id                 id.Id                            `json:"id,omitempty"`
	LegalName          string                           `json:"legal-name,omitempty"`
	TradeName          string                           `json:"trade-name,omitempty"`
	Members            []Member                         `json:"members,omitempty"`
	Managers           []Manager                        `json:"managers,omitempty"`
	Purpose            string                           `json:"purpose,omitempty"`
	Executor           Executor                         `json:"executor,omitempty"`
	MailingAddress     []mailing_address.MailingAddress `json:"mailing-addresses,omitempty"`
	StreetAddress      []mailing_address.MailingAddress `json:"street-addresses,omitempty"`
	County             location.County                  `json:"county,omitempty"`
	State              location.State                   `json:"state,omitempty"`
	ResponsibleParty   ResponsibleParty                 `json:"responsible-party,omitempty"`
	ForiengEntities    []ForeignEntity                  `json:"forieng-entities,omitempty"`
	EmailAddresses     []email_address.EmailAddress     `json:"email-addresses,omitempty"`
	PhoneNumbers       []phone.NanpNumber               `json:"phone-numbers,omitempty"`
	Units              int                              `json:"units,omitempty"`
	FormationDate      rfc3339date.Rfc3339Date          `json:"formation-date,omitempty"`
	DissolutionDate    rfc3339date.Rfc3339Date          `json:"dissolution-date,omitempty"`
	OperatingAgreement OperatingAgreement               `json:"operating-agreement,omitempty"`
}

type Member struct {
	Individual contact.Contact `json:"individuals,omitempty"`
	Business   Business        `json:"business,omitempty"`
	Shares     int             `json:"share,omitempty"`
}

func (t Member) Name() string {
	if !reflecter.Instances.IsZero(t.Individual) {
		return t.Individual.Name()
	}

	if !reflecter.Instances.IsZero(t.Business) {
		return t.Business.LegalName
	}

	return "X"
}

type Manager struct {
	Individual contact.Contact `json:"individual,omitempty"`
	Business   Business        `json:"business,omitempty"`
	Position   string          `json:"position,omitempty"`
}

func (t Manager) Name() string {
	if !reflecter.Instances.IsZero(t.Individual) {
		return t.Individual.Name()
	}

	if !reflecter.Instances.IsZero(t.Business) {
		return t.Business.LegalName
	}

	return "X"
}

type Executor interface {
	Name() string
	Address() string
}

type ResponsibleParty interface {
	Name() string
	Address() string
}
