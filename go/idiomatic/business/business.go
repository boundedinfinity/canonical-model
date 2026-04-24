package business

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/contact"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/digital/address/email"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/location/county"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/location/mailing_address"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/location/state"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/phone"
	"github.com/boundedinfinity/go-commoner/idiomatic/reflecter"
	"github.com/boundedinfinity/rfc3339date"
)

type Business struct {
	Id                 ider.Id                   `json:"id,omitempty"`
	LegalName          string                    `json:"legal-name,omitempty"`
	TradeName          string                    `json:"trade-name,omitempty"`
	Members            []Member                  `json:"members,omitempty"`
	Managers           []Manager                 `json:"managers,omitempty"`
	Purpose            string                    `json:"purpose,omitempty"`
	Executor           Executor                  `json:"executor,omitempty"`
	MailingAddress     []mailing_address.Address `json:"mailing-addresses,omitempty"`
	StreetAddress      []mailing_address.Address `json:"street-addresses,omitempty"`
	County             county.County             `json:"county,omitempty"`
	State              state.State               `json:"state,omitempty"`
	ResponsibleParty   ResponsibleParty          `json:"responsible-party,omitempty"`
	ForiengEntities    []ForeignEntity           `json:"forieng-entities,omitempty"`
	EmailAddresses     []email.Email             `json:"email-addresses,omitempty"`
	PhoneNumbers       []phone.NanpNumber        `json:"phone-numbers,omitempty"`
	Units              int                       `json:"units,omitempty"`
	FormationDate      rfc3339date.Rfc3339Date   `json:"formation-date,omitempty"`
	DissolutionDate    rfc3339date.Rfc3339Date   `json:"dissolution-date,omitempty"`
	OperatingAgreement OperatingAgreement        `json:"operating-agreement,omitempty"`
	EntityType         EntityType                `json:"entity-type,omitempty"`
}

type Member struct {
	Individual contact.Contact `json:"individuals,omitempty"`
	Business   Business        `json:"business,omitempty"`
	Shares     int             `json:"share,omitempty"`
}

func (t Member) Name() string {
	if !reflecter.IsZero[Member](t.Business) {
		return t.Business.LegalName
	}

	if !reflecter.IsZero[Member](t.Individual) {
		return t.Individual.Name()
	}

	return "X"
}

type Manager struct {
	Individual contact.Contact `json:"individual,omitempty"`
	Business   Business        `json:"business,omitempty"`
	Position   string          `json:"position,omitempty"`
}

func (t Manager) Name() string {
	if !reflecter.IsZero[Member](t.Individual) {
		return t.Individual.Name()
	}

	if !reflecter.IsZero[Member](t.Business) {
		return t.Business.LegalName
	}

	return "<no name>"
}

type Executor interface {
	Name() string
	Address() string
}

type ResponsibleParty interface {
	Name() string
	Address() string
}
