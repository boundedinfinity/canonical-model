package dental

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/person"
)

type Plan struct {
	Id           ider.Id            `json:"id,omitempty"`
	Member       person.PersonModel `json:"member,omitempty"`
	GroupId      string             `json:"group-id,omitempty"`
	GroupName    string             `json:"group-name,omitempty"`
	PlanType     string             `json:"plan-type,omitempty"`
	ContractType string             `json:"contract-type,omitempty"`
	ParStatus    string             `json:"par-status,omitempty"`
	// BenefitPersiod time.DateRange     `json:"benefit-period,omitempty"`
	Dentist Dentist `json:"dentist,omitempty"`
	Insurer Insurer `json:"insurer,omitempty"`
}
