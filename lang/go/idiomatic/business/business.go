package business

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/email_address"
	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/boundedinfinity/schema/idiomatic/location"
	"github.com/boundedinfinity/schema/idiomatic/mailing_address"
	"github.com/boundedinfinity/schema/idiomatic/phone"
)

type Business struct {
	Id               id.Id                            `json:"id,omitempty"`
	LegalName        string                           `json:"legal-name,omitempty"`
	TradeName        string                           `json:"trade-name,omitempty"`
	Purpose          string                           `json:"purpose,omitempty"`
	Executor         Executor                         `json:"executor,omitempty"`
	MailingAddress   []mailing_address.MailingAddress `json:"mailing-addresses,omitempty"`
	StreetAddress    []mailing_address.MailingAddress `json:"street-addresses,omitempty"`
	County           location.County                  `json:"county,omitempty"`
	State            location.State                   `json:"state,omitempty"`
	ResponsibleParty ResponsibleParty                 `json:"responsible-party,omitempty"`
	ForiengEntities  []ForeignEntity                  `json:"forieng-entities,omitempty"`
	EmailAddresses   []email_address.EmailAddress     `json:"email-addresses,omitempty"`
	PhoneNumbers     []phone.NanpNumber               `json:"phone-numbers,omitempty"`
	StartDate        rfc3339date.Rfc3339Date          `json:"start-date,omitempty"`
	EndDate          rfc3339date.Rfc3339Date          `json:"end-date,omitempty"`
	Units            int                              `json:"units,omitempty"`
}

type Executor interface {
	String() string
	Address() string
}

type ResponsibleParty interface {
	String() string
	Address() string
}
