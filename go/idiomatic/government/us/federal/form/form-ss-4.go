package form

import (
	"errors"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/business"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/location/county"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/location/mailing_address"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/person"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// https://www.irs.gov/pub/irs-pdf/fss4.pdf
// https://www.irs.gov/forms-pubs/about-form-ss-4

var (
	ErrLegalNameInvalid = errors.New("invalid legal name")
)

func NewFormSs4RevDecember2019(business business.Business) (FormSs4RevDecember_2019, error) {
	var form FormSs4RevDecember_2019

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

type FormSs4RevDecember_2019 struct {
	Id                  ider.Id                 `json:"id,omitempty"`
	LegalName           string                  `json:"legal-name,omitempty"`
	TradeName           string                  `json:"trade-name,omitempty"`
	Executor            person.PersonModel      `json:"executor,omitempty"`
	MailingAddress      mailing_address.Address `json:"mailing-address,omitempty"`
	StreetAddress       mailing_address.Address `json:"street-address,omitempty"`
	County              county.County           `json:"county,omitempty"`
	ResponsibleParty    person.PersonModel      `json:"responsible-party,omitempty"`
	ResponsiblePartyTin string                  `json:"responsible-party-tin,omitempty"`
	IsLlc               bool                    `json:"is-llc,omitempty"`
	LlcMemberCount      int                     `json:"llc-member-count,omitempty"`
	LlcOrganizedInUs    bool                    `json:"llc-organized-in-us,omitempty"`
}

func (t FormSs4RevDecember_2019) F1() string {
	return t.LegalName
}

func (t FormSs4RevDecember_2019) F2() string {
	return t.TradeName
}

func (t FormSs4RevDecember_2019) F3() string {
	return t.Executor.String()
}

func (t FormSs4RevDecember_2019) F4a() string {
	return stringer.Join(", ", t.MailingAddress.Lines...)
}

func (t FormSs4RevDecember_2019) F4b() string {
	return stringer.Join(
		", ",
		t.MailingAddress.City,
		t.MailingAddress.State,
		t.MailingAddress.Zip,
	)
}

func (t FormSs4RevDecember_2019) F7() string {
	return t.County.String()
}
