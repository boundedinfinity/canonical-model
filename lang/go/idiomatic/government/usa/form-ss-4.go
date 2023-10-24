package usa

import (
	"errors"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/schema/idiomatic/audit"
	"github.com/boundedinfinity/schema/idiomatic/business"
	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/boundedinfinity/schema/idiomatic/location"
	"github.com/boundedinfinity/schema/idiomatic/mailing_address"
	"github.com/boundedinfinity/schema/idiomatic/people"
)

// https://www.irs.gov/pub/irs-pdf/fss4.pdf
// https://www.irs.gov/forms-pubs/about-form-ss-4

var (
	ErrLegalNameInvalid = errors.New("invalid legal name")
)

func NewFormSs4RevDecember2019(business business.Business) (FormSs4RevDecember2019, error) {
	var form FormSs4RevDecember2019

	if business.LegalName == "" {
		return form, ErrLegalNameInvalid
	} else {
		form.LegalName = business.LegalName
	}

	if business.TradeName != "" {
		form.TradeName = business.TradeName
	}

	return form, nil
}

type FormSs4RevDecember2019 struct {
	Id                  id.Id                          `json:"id,omitempty"`
	LegalName           string                         `json:"legal-name,omitempty"`
	TradeName           string                         `json:"trade-name,omitempty"`
	Executor            people.Person                  `json:"executor,omitempty"`
	MailingAddress      mailing_address.MailingAddress `json:"mailing-address,omitempty"`
	StreetAddress       mailing_address.MailingAddress `json:"street-address,omitempty"`
	County              location.County                `json:"county,omitempty"`
	ResponsibleParty    people.Person                  `json:"responsible-party,omitempty"`
	ResponsiblePartyTin string                         `json:"responsible-party-tin,omitempty"`
	IsLlc               bool                           `json:"is-llc,omitempty"`
	LlcMemberCount      int                            `json:"llc-member-count,omitempty"`
	LlcOrganizedInUs    bool                           `json:"llc-organized-in-us,omitempty"`
	Audit               audit.Audit                    `json:"audit,omitempty"`
}

func (t FormSs4RevDecember2019) F1() string {
	return t.LegalName
}

func (t FormSs4RevDecember2019) F2() string {
	return t.TradeName
}

func (t FormSs4RevDecember2019) F3() string {
	return t.Executor.Name.String()
}

func (t FormSs4RevDecember2019) F4a() string {
	return stringer.Join(", ", t.MailingAddress.Lines...)
}

func (t FormSs4RevDecember2019) F4b() string {
	return stringer.Join(
		", ",
		t.MailingAddress.City.String(),
		t.MailingAddress.State.String(),
		t.MailingAddress.Zip.String(),
	)
}

func (t FormSs4RevDecember2019) F7() string {
	return t.County.String()
}
