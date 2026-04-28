package insurance

// https://www.bankrate.com/insurance/car/carriers/
// https://www.allstate.com/auto-insurance/car-policy-declarations
// https://help.everlance.com/hc/en-us/articles/7021058743949-How-to-find-and-read-your-insurance-declaration-page
// https://www.bankrate.com/insurance/homeowners-insurance/insurance-declaration-page/
// https://www.bankrate.com/insurance/car/certificate-of-insurance/#what-is-a-certificate-of-insurance
// https://www.thebalancemoney.com/insurance-services-office-iso-462706

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/business"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/contact"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/vehicle"
	"github.com/boundedinfinity/rfc3339date"
)

type InsuranceCoverage struct {
	Symbol     string            `json:"symbol,omitempty"`
	Desc       string            `json:"desc,omitempty"`
	Limits     []InsureanceLimit `json:"limits,omitempty"`
	Premiums   float32           `json:"premiums,omitempty"`
	Deductible float32           `json:"deductible,omitempty"`
}

type InsuranceCoveragePeriod struct {
	StartDate rfc3339date.Rfc3339Date `json:"start-date,omitempty"`
	EndDate   rfc3339date.Rfc3339Date `json:"end-date,omitempty"`
}

type InsureanceLimit struct {
	Desc   string                 `json:"desc,omitempty"`
	Values []InsureanceLimitValue `json:"values,omitempty"`
}

type InsureanceLimitValue struct {
	Desc  string  `json:"desc,omitempty"`
	Limit float32 `json:"limit,omitempty"`
}

type AutoInsurancePolicy struct {
	Id             ider.Id                 `json:"id,omitempty"`
	PolicyNumber   string                  `json:"policy-number,omitempty"`
	Vechicle       vehicle.Vehicle         `json:"vehicle,omitempty"`
	Class          string                  `json:"class,omitempty"`
	MonthlyPremium float32                 `json:"monthly-premium,omitempty"`
	Provider       business.Business       `json:"provider,omitempty"`
	Agent          contact.Contact         `json:"agent,omitempty"`
	Deductible     float32                 `json:"deductible,omitempty"`
	Coverages      []InsuranceCoverage     `json:"coverages,omitempty"`
	Period         InsuranceCoveragePeriod `json:"period,omitempty"`
}

func (t AutoInsurancePolicy) Premium() (float32, error) {
	var total float32

	for _, coverage := range t.Coverages {
		total += coverage.Deductible
	}

	return total, nil
}
